package server

import (
	"context"

	"github.com/golang-jwt/jwt/v4"

	er "github.com/Stern-Ritter/gophertask/internal/errors"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
	storage "github.com/Stern-Ritter/gophertask/internal/storage/server"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByLogin(ctx context.Context, login string) (model.User, error)
	GetCurrentUser(ctx context.Context) (model.User, error)
}

type UserServiceImpl struct {
	userStorage storage.UserStorage
	logger      *logger.ServerLogger
}

func NewUserService(userStorage storage.UserStorage, logger *logger.ServerLogger) UserService {
	return &UserServiceImpl{
		userStorage: userStorage,
		logger:      logger,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	return s.userStorage.Create(ctx, user)
}

func (s *UserServiceImpl) GetUserByLogin(ctx context.Context, login string) (model.User, error) {
	return s.userStorage.GetOneByLogin(ctx, login)
}

func (s *UserServiceImpl) GetCurrentUser(ctx context.Context) (model.User, error) {
	claims, ok := ctx.Value(AuthorizationTokenContextKey).(jwt.MapClaims)
	if !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	if _, ok := claims["login"]; !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	login, ok := claims["login"].(string)
	if !ok {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}

	currentUser, err := s.userStorage.GetOneByLogin(ctx, login)
	if err != nil {
		return model.User{}, er.NewUnauthorizedError("user is not authorized to access this resource", nil)
	}
	return currentUser, err
}
