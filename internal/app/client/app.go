package client

import (
	"fmt"

	"github.com/rivo/tview"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"

	config "github.com/Stern-Ritter/gophertask/internal/config/client"
	crypto "github.com/Stern-Ritter/gophertask/internal/crypto/client"
	logger "github.com/Stern-Ritter/gophertask/internal/logger/client"
	service "github.com/Stern-Ritter/gophertask/internal/service/client"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

func Run(client service.Client, app *tview.Application, cfg *config.ClientConfig, logger *logger.ClientLogger) error {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	opts = append(opts, grpc.WithUnaryInterceptor(client.AuthInterceptor))
	creds, err := crypto.GetTransportCredentials()
	if err != nil {
		logger.Fatal(err.Error(), zap.String("event", "load credentials"))
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.NewClient(cfg.ServerURL, opts...)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to connect to server: %s", cfg.ServerURL),
			zap.String("event", "connect server"), zap.Error(err))
	}
	defer conn.Close()

	authService := service.NewAuthService(pb.NewAuthServiceV1Client(conn))
	taskService := service.NewTaskService(pb.NewTaskServiceV1Client(conn))
	subtaskService := service.NewSubtaskService(pb.NewSubtaskServiceV1Client(conn))
	epicService := service.NewEpicService(pb.NewEpicServiceV1Client(conn))

	client.SetAuthService(authService)
	client.SetTaskService(taskService)
	client.SetSubtaskService(subtaskService)
	client.SetEpicService(epicService)
	client.SetApp(app)

	client.AuthView()
	if err := app.Run(); err != nil {
		logger.Error(fmt.Sprintf("Failed to run application: %s", cfg.ServerURL),
			zap.String("event", "run application"), zap.Error(err))
		return err
	}

	return nil
}
