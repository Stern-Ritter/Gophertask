package model

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Stern-Ritter/gophertask/internal/utils"
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type Subtask struct {
	Task
	epicID string
}

func NewSubtask(id string, userID string, epicID string, name string, description string, status TaskStatus, duration time.Duration,
	startedAt time.Time) Subtask {
	return Subtask{
		Task:   NewTask(id, userID, name, description, status, duration, startedAt),
		epicID: epicID,
	}
}

func (s *Subtask) EpicID() string {
	return s.epicID
}

func AddSubtaskRequestToSubtask(req *pb.AddSubtaskRequestV1) Subtask {
	return Subtask{
		Task: Task{
			name:        req.GetName(),
			description: req.GetDescription(),
			status:      New,
			duration:    req.GetDuration().AsDuration(),
			startedAt:   req.GetStartedAt().AsTime(),
		},
		epicID: req.GetEpicId(),
	}
}

func UpdateSubtaskRequestToSubtask(req *pb.UpdateSubtaskRequestV1, subtask Subtask) Subtask {
	return Subtask{
		Task: Task{
			id:          subtask.ID(),
			userID:      subtask.UserID(),
			name:        req.GetName(),
			description: req.GetDescription(),
			status:      MapMessageTaskStatusToTaskStatus(req.GetStatus()),
			duration:    req.GetDuration().AsDuration(),
			startedAt:   req.GetStartedAt().AsTime(),
		},
		epicID: utils.Coalesce(req.GetEpicId(), subtask.EpicID()),
	}
}

func SubtaskToSubtaskMessage(s Subtask) *pb.SubtaskV1 {
	return &pb.SubtaskV1{
		Id:          s.ID(),
		EpicId:      s.EpicID(),
		Name:        s.Name(),
		Description: s.Description(),
		Status:      MapTaskStatusToMessageTaskStatus(s.Status()),
		Duration:    durationpb.New(s.Duration()),
		StartedAt:   timestamppb.New(s.StartedAt()),
	}
}

func SubtaskToRepeatedSubtaskMessage(s []Subtask) []*pb.SubtaskV1 {
	subtasks := make([]*pb.SubtaskV1, len(s))
	for i, subtask := range s {
		subtasks[i] = SubtaskToSubtaskMessage(subtask)
	}

	return subtasks
}
