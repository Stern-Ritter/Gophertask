package model

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type Task struct {
	id          string
	userID      string
	name        string
	description string
	status      TaskStatus
	duration    time.Duration
	startedAt   time.Time
}

func NewTask(id string, userID string, name string, description string, status TaskStatus, duration time.Duration,
	startedAt time.Time) Task {
	return Task{
		id:          id,
		userID:      userID,
		name:        name,
		description: description,
		status:      status,
		duration:    duration,
		startedAt:   startedAt,
	}
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) UserID() string {
	return t.userID
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() TaskStatus {
	return t.status
}

func (t *Task) Duration() time.Duration {
	return t.duration
}

func (t *Task) StartedAt() time.Time {
	return t.startedAt
}

func (t *Task) EndedAt() time.Time {
	return t.startedAt.Add(t.duration)
}

func (t *Task) SetUserID(userID string) {
	t.userID = userID
}

func AddTaskRequestToTask(req *pb.AddTaskRequestV1) Task {
	return Task{
		name:        req.GetName(),
		description: req.GetDescription(),
		status:      New,
		duration:    req.GetDuration().AsDuration(),
		startedAt:   req.GetStartedAt().AsTime(),
	}
}

func UpdateTaskRequestToTask(req *pb.UpdateTaskRequestV1, task Task) Task {
	return Task{
		id:          task.ID(),
		userID:      task.UserID(),
		name:        req.GetName(),
		description: req.GetDescription(),
		status:      MapMessageTaskStatusToTaskStatus(req.GetStatus()),
		duration:    req.GetDuration().AsDuration(),
		startedAt:   req.GetStartedAt().AsTime(),
	}
}

func TaskToTaskMessage(t Task) *pb.TaskV1 {
	return &pb.TaskV1{
		Id:          t.ID(),
		Name:        t.Name(),
		Description: t.Description(),
		Status:      MapTaskStatusToMessageTaskStatus(t.Status()),
		Duration:    durationpb.New(t.Duration()),
		StartedAt:   timestamppb.New(t.StartedAt()),
	}
}

func TasksToRepeatedTaskMessage(t []Task) []*pb.TaskV1 {
	tasks := make([]*pb.TaskV1, len(t))
	for i, task := range t {
		tasks[i] = TaskToTaskMessage(task)
	}

	return tasks
}
