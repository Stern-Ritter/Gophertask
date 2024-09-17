package model

import (
	"sort"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type Epic struct {
	Task
	subtasks []Subtask
}

func NewEpic(id string, userID string, name string, description string, startedAt time.Time, subtasks []Subtask) Epic {
	return Epic{
		Task:     Task{id: id, userID: userID, name: name, description: description, startedAt: startedAt},
		subtasks: subtasks,
	}
}

func (e *Epic) UpdateStatus() {
	if len(e.subtasks) == 0 {
		e.status = New
	}

	hasOnlyNewSubtasks := true
	for _, subtask := range e.subtasks {
		if subtask.status != New {
			hasOnlyNewSubtasks = false
			break
		}
	}

	hasOnlyDoneSubtasks := true
	for _, subtask := range e.subtasks {
		if subtask.status != Done {
			hasOnlyDoneSubtasks = false
			break
		}
	}

	switch {
	case hasOnlyNewSubtasks:
		e.status = New
	case hasOnlyDoneSubtasks:
		e.status = Done
	default:
		e.status = InProgress
	}
}

func (e *Epic) Status() TaskStatus {
	e.UpdateStatus()
	return e.status
}

func (e *Epic) Duration() time.Duration {
	duration := time.Duration(0)
	for _, subtask := range e.subtasks {
		duration += subtask.Duration()
	}

	return duration
}

func (e *Epic) StartedAt() time.Time {
	if len(e.subtasks) == 0 {
		return e.startedAt
	}

	subtasksStartedAt := make([]time.Time, len(e.subtasks))
	for i, subtask := range e.subtasks {
		subtasksStartedAt[i] = subtask.startedAt
	}
	sort.Slice(subtasksStartedAt, func(i, j int) bool {
		return subtasksStartedAt[i].Before(subtasksStartedAt[j])
	})

	return subtasksStartedAt[0]
}

func (e *Epic) EndedAt() time.Time {
	if len(e.subtasks) == 0 {
		return e.Task.EndedAt()
	}

	subtasksEndedAt := make([]time.Time, len(e.subtasks))
	for i, subtask := range e.subtasks {
		subtasksEndedAt[i] = subtask.EndedAt()
	}
	sort.Slice(subtasksEndedAt, func(i, j int) bool {
		return subtasksEndedAt[i].After(subtasksEndedAt[j])
	})

	return subtasksEndedAt[0]
}

func (e *Epic) Subtasks() []Subtask {
	return e.subtasks
}

func (e *Epic) AddSubtask(subtask Subtask) {
	e.subtasks = append(e.subtasks, subtask)
}

func AddEpicRequestToEpic(req *pb.AddEpicRequestV1) Epic {
	return Epic{
		Task: Task{
			name:        req.GetName(),
			description: req.GetDescription(),
			status:      New,
			startedAt:   req.GetStartedAt().AsTime(),
		},
	}
}

func UpdateEpicRequestToEpic(req *pb.UpdateEpicRequestV1, epic Epic) Epic {
	return Epic{
		Task: Task{
			id:          epic.ID(),
			userID:      epic.UserID(),
			name:        req.GetName(),
			description: req.GetDescription(),
			startedAt:   req.GetStartedAt().AsTime(),
		},
	}
}

func EpicToEpicMessage(e Epic) *pb.EpicV1 {
	return &pb.EpicV1{
		Id:          e.ID(),
		Name:        e.Name(),
		Description: e.Description(),
		Status:      MapTaskStatusToMessageTaskStatus(e.Status()),
		Duration:    durationpb.New(e.Duration()),
		StartedAt:   timestamppb.New(e.StartedAt()),
		Subtasks:    SubtaskToRepeatedSubtaskMessage(e.Subtasks()),
	}
}

func EpicsToRepeatedEpicMessage(e []Epic) []*pb.EpicV1 {
	epics := make([]*pb.EpicV1, len(e))
	for i, epic := range e {
		epics[i] = EpicToEpicMessage(epic)
	}

	return epics
}
