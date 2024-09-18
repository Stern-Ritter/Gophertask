package server

import (
	"github.com/bufbuild/protovalidate-go"

	config "github.com/Stern-Ritter/gophertask/internal/config/server"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/server"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type Server struct {
	UserService    UserService
	AuthService    AuthService
	TaskService    TaskService
	SubtaskService SubtaskService
	EpicService    EpicService
	Validator      *protovalidate.Validator
	Config         *config.ServerConfig
	Logger         *logger.ServerLogger
	pb.UnimplementedAuthServiceV1Server
	pb.UnimplementedTaskServiceV1Server
	pb.UnimplementedSubtaskServiceV1Server
	pb.UnimplementedEpicServiceV1Server
}

func NewServer(userService UserService, authService AuthService, taskService TaskService, subtaskService SubtaskService,
	epicService EpicService, validator *protovalidate.Validator, config *config.ServerConfig, logger *logger.ServerLogger) *Server {
	return &Server{
		UserService:    userService,
		AuthService:    authService,
		TaskService:    taskService,
		SubtaskService: subtaskService,
		EpicService:    epicService,
		Validator:      validator,
		Config:         config,
		Logger:         logger,
	}
}
