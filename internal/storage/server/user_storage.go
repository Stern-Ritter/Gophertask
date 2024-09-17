package server

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	"github.com/Stern-Ritter/gophertask/internal/model"
)

type UserStorage interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetOneByLogin(ctx context.Context, login string) (model.User, error)
}

type UserStorageImpl struct {
	db     PgxIface
	Logger *logger.ServerLogger
}

func NewUserStorage(db PgxIface, logger *logger.ServerLogger) UserStorage {
	return &UserStorageImpl{
		db:     db,
		Logger: logger,
	}
}

func (s *UserStorageImpl) Create(ctx context.Context, user model.User) (model.User, error) {
	var userID string
	row := s.db.QueryRow(ctx, `
		INSERT INTO gophertask.users (login, password)
		VALUES (@login, @password)
		RETURNING id
	`, pgx.NamedArgs{
		"login":    user.Login,
		"password": user.Password,
	})

	err := row.Scan(&userID)
	if err != nil {
		s.Logger.Debug("Failed to insert user", zap.String("event", "create user"),
			zap.String("user", fmt.Sprintf("%v", user)), zap.Error(err))
		return model.User{}, err
	}

	user.ID = userID
	return user, nil
}

func (s *UserStorageImpl) GetOneByLogin(ctx context.Context, login string) (model.User, error) {
	row := s.db.QueryRow(ctx, `
		SELECT id, login, password
		FROM gophertask.users
		WHERE login = @login
	`, pgx.NamedArgs{
		"login": login,
	})

	user := model.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		s.Logger.Debug("Failed select user by login", zap.String("event", "get user by login"),
			zap.String("login", login), zap.Error(err))
		return model.User{}, err
	}

	return user, nil
}
