package model

import (
	pb "github.com/Stern-Ritter/gophertask/proto/gen/gophertask/gophertaskapi/v1"
)

type TaskStatus string

const (
	New        TaskStatus = "NEW"
	InProgress TaskStatus = "IN PROGRESS"
	Done       TaskStatus = "DONE"
)

func MapMessageTaskStatusToTaskStatus(status pb.TaskStatus) TaskStatus {
	switch status {
	case pb.TaskStatus_TASK_STATUS_NEW:
		return New
	case pb.TaskStatus_TASK_STATUS_IN_PROGRESS:
		return InProgress
	case pb.TaskStatus_TASK_STATUS_DONE:
		return Done
	default:
		return New
	}
}

func MapTaskStatusToMessageTaskStatus(status TaskStatus) pb.TaskStatus {
	switch status {
	case New:
		return pb.TaskStatus_TASK_STATUS_NEW
	case InProgress:
		return pb.TaskStatus_TASK_STATUS_IN_PROGRESS
	case Done:
		return pb.TaskStatus_TASK_STATUS_DONE
	default:
		return pb.TaskStatus_TASK_STATUS_NEW
	}
}
