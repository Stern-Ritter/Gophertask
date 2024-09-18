package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bufbuild/protovalidate-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "google.golang.org/grpc/encoding/gzip"

	config "github.com/Stern-Ritter/gophertask/internal/config/server"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	service "github.com/Stern-Ritter/gophertask/internal/service/server"
	storage "github.com/Stern-Ritter/gophertask/internal/storage/server"
	"github.com/Stern-Ritter/gophertask/migrations"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

func Run(cfg *config.ServerConfig, logger *logger.ServerLogger) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	idleConnsClosed := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DatabaseDSN)
	if err != nil {
		logger.Fatal("Failed to create database connection", zap.String("event", "create database connection"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}
	err = db.Ping(ctx)
	if err != nil {
		logger.Fatal("Failed to connect database", zap.String("event", "connect database"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}
	err = migrations.MigrateDatabase(cfg.DatabaseDSN)
	if err != nil {
		logger.Fatal("Failed to migrate database", zap.String("event", "migrate database"),
			zap.String("database url", cfg.DatabaseDSN), zap.Error(err))
	}

	userStorage := storage.NewUserStorage(db, logger)
	taskStorage := storage.NewTaskStorage(db, logger)
	subtaskStorage := storage.NewSubtaskStorage(db, logger)
	epicStorage := storage.NewEpicStorage(db, logger)

	userService := service.NewUserService(userStorage, logger)
	authService := service.NewAuthService(userService, cfg.AuthenticationKey, logger)
	taskService := service.NewTaskService(taskStorage, logger)
	subtaskService := service.NewSubtaskService(subtaskStorage, logger)
	epicService := service.NewEpicService(epicStorage, logger)

	validator, err := protovalidate.New()
	if err != nil {
		logger.Fatal("Failed to initialize validator", zap.String("event", "initialize validator"),
			zap.Error(err))
	}

	server := service.NewServer(userService, authService, taskService, subtaskService, epicService, validator, cfg, logger)

	err = runGrpcServer(server, signals, idleConnsClosed)
	return err
}

func runGrpcServer(server *service.Server, signals chan os.Signal, idleConnsClosed chan struct{}) error {
	listen, err := net.Listen("tcp", server.Config.URL)
	if err != nil {
		return err
	}

	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.ChainUnaryInterceptor(logger.LoggerInterceptor(server.Logger)))
	opts = append(opts, grpc.ChainUnaryInterceptor(server.AuthInterceptor))
	creds, err := credentials.NewServerTLSFromFile(server.Config.TLSCertPath, server.Config.TLSKeyPath)
	if err != nil {
		server.Logger.Fatal(err.Error(), zap.String("event", "load credentials"))
	}
	opts = append(opts, grpc.Creds(creds))

	srv := grpc.NewServer(opts...)
	pb.RegisterAuthServiceV1Server(srv, server)
	pb.RegisterTaskServiceV1Server(srv, server)
	pb.RegisterSubtaskServiceV1Server(srv, server)
	pb.RegisterEpicServiceV1Server(srv, server)

	go func() {
		<-signals

		server.Logger.Info("Shutting down server", zap.String("event", "shutdown server"))
		srv.GracefulStop()

		close(idleConnsClosed)
	}()

	server.Logger.Info("Starting server", zap.String("event", "start server"),
		zap.String("url", server.Config.URL))
	if err := srv.Serve(listen); err != nil {
		return err
	}

	<-idleConnsClosed
	server.Logger.Info("Server shutdown complete", zap.String("event", "shutdown server"))

	return nil
}
