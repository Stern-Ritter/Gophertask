package client

import (
	"context"

	"github.com/rivo/tview"
	"google.golang.org/grpc"

	config "github.com/Stern-Ritter/gophertask/internal/config/client"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type Client interface {
	SetAuthService(authService AuthService)
	SetTaskService(taskService TaskService)
	SetSubtaskService(subtaskService SubtaskService)
	SetEpicService(epicService EpicService)
	SetApp(*tview.Application)
	AuthInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error
	SelectView(view tview.Primitive)
	AuthView() tview.Primitive
	MainView() tview.Primitive
	AddView() tview.Primitive
	AddTaskView() tview.Primitive
	AddEpicView() tview.Primitive
	AddSubtaskView(epicID string) tview.Primitive
	UpdateTaskView(task *pb.TaskV1) tview.Primitive
	UpdateEpicView(epic *pb.EpicV1) tview.Primitive
	UpdateSubtaskView(subtask *pb.SubtaskV1) tview.Primitive
	TasksView(previousView tview.Primitive) tview.Primitive
	EpicsView(previousView tview.Primitive) tview.Primitive
	SubtasksView(epicID string, previousView tview.Primitive) tview.Primitive
	ShowInfoModal(text string, currentView tview.Primitive) tview.Primitive
	ShowConfirmModal(text string, currentView tview.Primitive, handler func()) tview.Primitive
	ShowRetryModal(text string, currentView tview.Primitive, previousView tview.Primitive) tview.Primitive
}

type ClientImpl struct {
	authService    AuthService
	taskService    TaskService
	subtaskService SubtaskService
	epicService    EpicService
	app            *tview.Application
	authToken      string
	config         *config.ClientConfig
}

func NewClient(config *config.ClientConfig) Client {
	return &ClientImpl{
		config: config,
	}
}

func (c *ClientImpl) SetAuthService(authService AuthService) {
	c.authService = authService
}

func (c *ClientImpl) SetTaskService(taskService TaskService) {
	c.taskService = taskService
}

func (c *ClientImpl) SetSubtaskService(subtaskService SubtaskService) {
	c.subtaskService = subtaskService
}

func (c *ClientImpl) SetEpicService(epicService EpicService) {
	c.epicService = epicService
}

func (c *ClientImpl) SetApp(app *tview.Application) {
	c.app = app
}
