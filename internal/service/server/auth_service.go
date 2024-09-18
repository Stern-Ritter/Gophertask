package server

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/Stern-Ritter/gophertask/internal/auth"
	er "github.com/Stern-Ritter/gophertask/internal/errors"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

const (
	AuthTokenExpiration          = time.Hour * 24 // AuthTokenExpiration defines the duration for which an authentication token is valid.
	UserLoginUniqueConstrainName = "users_login_unique"
)

type AuthService interface {
	SignUp(ctx context.Context, req model.SignUpRequest) (string, error)
	SignIn(ctx context.Context, req model.SignInRequest) (string, error)
}

type AuthServiceImpl struct {
	userService   UserService
	authSecretKey string
	logger        *logger.ServerLogger
}

func NewAuthService(userService UserService, authSecretKey string, logger *logger.ServerLogger) AuthService {
	return &AuthServiceImpl{
		userService:   userService,
		authSecretKey: authSecretKey,
		logger:        logger,
	}
}

func (s *AuthServiceImpl) SignUp(ctx context.Context, req model.SignUpRequest) (string, error) {
	user := model.SignUpRequestToUser(req)

	passwordHash, err := auth.GetPasswordHash(user.Password)
	if err != nil {
		s.logger.Error("Failed to generate password hash", zap.String("event", "user registration"),
			zap.Error(err))
		return "", fmt.Errorf("failed to generate password hash: %w", err)
	}

	user.Password = passwordHash

	savedUser, err := s.userService.CreateUser(ctx, user)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == UserLoginUniqueConstrainName {
			s.logger.Info("User already exists", zap.String("event", "user registration"),
				zap.String("login", user.Login), zap.Error(err))
			return "", er.NewConflictError(fmt.Sprintf("user with login: %s already exists", user.Login), err)
		}

		return "", err
	}

	token, err := auth.NewToken(savedUser, s.authSecretKey, AuthTokenExpiration)
	if err != nil {
		s.logger.Error("Failed to generate jwt token", zap.String("event", "user registration"),
			zap.String("login", user.Login), zap.Error(err))
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}

func (s *AuthServiceImpl) SignIn(ctx context.Context, req model.SignInRequest) (string, error) {
	user, err := s.userService.GetUserByLogin(ctx, req.Login)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			s.logger.Info("User not exists", zap.String("event", "user authentication"),
				zap.String("login", req.Login), zap.Error(err))
			return "", er.NewUnauthorizedError("invalid login or password", err)
		default:
			s.logger.Error("Failed to get user by login", zap.String("event", "user authentication"),
				zap.String("login", req.Login), zap.Error(err))
			return "", fmt.Errorf("failed to get user by login: %w", err)
		}
	}

	if !auth.CheckPasswordHash(req.Password, user.Password) {
		s.logger.Info("Invalid password", zap.String("event", "user authentication"),
			zap.String("login", req.Login), zap.Error(err))
		return "", er.NewUnauthorizedError("invalid login or password", err)
	}

	token, err := auth.NewToken(user, s.authSecretKey, AuthTokenExpiration)
	if err != nil {
		s.logger.Error("Failed to generate jwt token", zap.String("event", "user authentication"),
			zap.String("login", req.Login), zap.Error(err))
		return "", fmt.Errorf("failed to generate jwt token: %w", err)
	}

	return token, nil
}
