// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gophertask/gophertaskapi/v1/task.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus             `protobuf:"varint,5,opt,name=status,proto3,enum=gophertask.gophertaskapi.v1.TaskStatus" json:"status,omitempty"`
	Duration    *durationpb.Duration   `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
}

func (x *TaskV1) Reset() {
	*x = TaskV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskV1) ProtoMessage() {}

func (x *TaskV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskV1.ProtoReflect.Descriptor instead.
func (*TaskV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskV1) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TaskV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskV1) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN
}

func (x *TaskV1) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TaskV1) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

type AddTaskRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Duration    *durationpb.Duration   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
}

func (x *AddTaskRequestV1) Reset() {
	*x = AddTaskRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequestV1) ProtoMessage() {}

func (x *AddTaskRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequestV1.ProtoReflect.Descriptor instead.
func (*AddTaskRequestV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *AddTaskRequestV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddTaskRequestV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddTaskRequestV1) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *AddTaskRequestV1) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

type AddTaskResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddTaskResponseV1) Reset() {
	*x = AddTaskResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponseV1) ProtoMessage() {}

func (x *AddTaskResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponseV1.ProtoReflect.Descriptor instead.
func (*AddTaskResponseV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{2}
}

type UpdateTaskRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      TaskStatus             `protobuf:"varint,4,opt,name=status,proto3,enum=gophertask.gophertaskapi.v1.TaskStatus" json:"status,omitempty"`
	Duration    *durationpb.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
}

func (x *UpdateTaskRequestV1) Reset() {
	*x = UpdateTaskRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequestV1) ProtoMessage() {}

func (x *UpdateTaskRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequestV1.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequestV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskRequestV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTaskRequestV1) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTaskRequestV1) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTaskRequestV1) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN
}

func (x *UpdateTaskRequestV1) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *UpdateTaskRequestV1) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

type UpdateTaskResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTaskResponseV1) Reset() {
	*x = UpdateTaskResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskResponseV1) ProtoMessage() {}

func (x *UpdateTaskResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskResponseV1.ProtoReflect.Descriptor instead.
func (*UpdateTaskResponseV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{4}
}

type DeleteTaskRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequestV1) Reset() {
	*x = DeleteTaskRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequestV1) ProtoMessage() {}

func (x *DeleteTaskRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequestV1.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequestV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTaskRequestV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTaskResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTaskResponseV1) Reset() {
	*x = DeleteTaskResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponseV1) ProtoMessage() {}

func (x *DeleteTaskResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponseV1.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponseV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{6}
}

type GetTaskByIDRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskByIDRequestV1) Reset() {
	*x = GetTaskByIDRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByIDRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByIDRequestV1) ProtoMessage() {}

func (x *GetTaskByIDRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByIDRequestV1.ProtoReflect.Descriptor instead.
func (*GetTaskByIDRequestV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{7}
}

func (x *GetTaskByIDRequestV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTaskByIDResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *TaskV1 `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *GetTaskByIDResponseV1) Reset() {
	*x = GetTaskByIDResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskByIDResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskByIDResponseV1) ProtoMessage() {}

func (x *GetTaskByIDResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskByIDResponseV1.ProtoReflect.Descriptor instead.
func (*GetTaskByIDResponseV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{8}
}

func (x *GetTaskByIDResponseV1) GetTask() *TaskV1 {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTasksRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTasksRequestV1) Reset() {
	*x = GetTasksRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequestV1) ProtoMessage() {}

func (x *GetTasksRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequestV1.ProtoReflect.Descriptor instead.
func (*GetTasksRequestV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{9}
}

type GetTasksResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*TaskV1 `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTasksResponseV1) Reset() {
	*x = GetTasksResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksResponseV1) ProtoMessage() {}

func (x *GetTasksResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_gophertask_gophertaskapi_v1_task_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksResponseV1.ProtoReflect.Descriptor instead.
func (*GetTasksResponseV1) Descriptor() ([]byte, []int) {
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP(), []int{10}
}

func (x *GetTasksResponseV1) GetTasks() []*TaskV1 {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_gophertask_gophertaskapi_v1_task_proto protoreflect.FileDescriptor

var file_gophertask_gophertaskapi_v1_task_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02,
	0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x02, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12,
	0x5a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x46, 0xba,
	0x48, 0x43, 0xba, 0x01, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x4e, 0x61, 0x6d,
	0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x63,
	0x65, 0x65, 0x64, 0x20, 0x32, 0x35, 0x36, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x11, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3c,
	0x3d, 0x20, 0x32, 0x35, 0x36, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7a, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x58, 0xba, 0x48, 0x55, 0xba, 0x01, 0x52, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78,
	0x63, 0x65, 0x65, 0x64, 0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x13, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x29, 0x20, 0x3c, 0x3d, 0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x64, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0xe5,
	0x03, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x43, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x33, 0xba, 0x48, 0x30, 0xba, 0x01, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x49, 0x44, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x29, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x52, 0x02, 0x69, 0x64, 0x12, 0x5a, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x46, 0xba, 0x48, 0x43, 0xba, 0x01,
	0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x20,
	0x32, 0x35, 0x36, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x11,
	0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x32, 0x35,
	0x36, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x58, 0xba, 0x48,
	0x55, 0xba, 0x01, 0x52, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64,
	0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x1a, 0x13, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3c, 0x3d,
	0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x67,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x50, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x40, 0xba, 0x48, 0x3d, 0xba, 0x01, 0x3a, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x23, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20, 0x49, 0x44, 0x20, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3e,
	0x3d, 0x20, 0x31, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22,
	0x6a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x52, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x42, 0xba, 0x48, 0x3f, 0xba, 0x01, 0x3c, 0x0a, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x20, 0x74, 0x61, 0x73, 0x6b, 0x20,
	0x49, 0x44, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x29, 0x20, 0x3e, 0x3d, 0x20, 0x31, 0x52, 0x02, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x31, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x31, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x39, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x31, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x74, 0x61, 0x73, 0x6b, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gophertask_gophertaskapi_v1_task_proto_rawDescOnce sync.Once
	file_gophertask_gophertaskapi_v1_task_proto_rawDescData = file_gophertask_gophertaskapi_v1_task_proto_rawDesc
)

func file_gophertask_gophertaskapi_v1_task_proto_rawDescGZIP() []byte {
	file_gophertask_gophertaskapi_v1_task_proto_rawDescOnce.Do(func() {
		file_gophertask_gophertaskapi_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_gophertask_gophertaskapi_v1_task_proto_rawDescData)
	})
	return file_gophertask_gophertaskapi_v1_task_proto_rawDescData
}

var file_gophertask_gophertaskapi_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_gophertask_gophertaskapi_v1_task_proto_goTypes = []any{
	(*TaskV1)(nil),                // 0: gophertask.gophertaskapi.v1.TaskV1
	(*AddTaskRequestV1)(nil),      // 1: gophertask.gophertaskapi.v1.AddTaskRequestV1
	(*AddTaskResponseV1)(nil),     // 2: gophertask.gophertaskapi.v1.AddTaskResponseV1
	(*UpdateTaskRequestV1)(nil),   // 3: gophertask.gophertaskapi.v1.UpdateTaskRequestV1
	(*UpdateTaskResponseV1)(nil),  // 4: gophertask.gophertaskapi.v1.UpdateTaskResponseV1
	(*DeleteTaskRequestV1)(nil),   // 5: gophertask.gophertaskapi.v1.DeleteTaskRequestV1
	(*DeleteTaskResponseV1)(nil),  // 6: gophertask.gophertaskapi.v1.DeleteTaskResponseV1
	(*GetTaskByIDRequestV1)(nil),  // 7: gophertask.gophertaskapi.v1.GetTaskByIDRequestV1
	(*GetTaskByIDResponseV1)(nil), // 8: gophertask.gophertaskapi.v1.GetTaskByIDResponseV1
	(*GetTasksRequestV1)(nil),     // 9: gophertask.gophertaskapi.v1.GetTasksRequestV1
	(*GetTasksResponseV1)(nil),    // 10: gophertask.gophertaskapi.v1.GetTasksResponseV1
	(TaskStatus)(0),               // 11: gophertask.gophertaskapi.v1.TaskStatus
	(*durationpb.Duration)(nil),   // 12: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_gophertask_gophertaskapi_v1_task_proto_depIdxs = []int32{
	11, // 0: gophertask.gophertaskapi.v1.TaskV1.status:type_name -> gophertask.gophertaskapi.v1.TaskStatus
	12, // 1: gophertask.gophertaskapi.v1.TaskV1.duration:type_name -> google.protobuf.Duration
	13, // 2: gophertask.gophertaskapi.v1.TaskV1.started_at:type_name -> google.protobuf.Timestamp
	12, // 3: gophertask.gophertaskapi.v1.AddTaskRequestV1.duration:type_name -> google.protobuf.Duration
	13, // 4: gophertask.gophertaskapi.v1.AddTaskRequestV1.started_at:type_name -> google.protobuf.Timestamp
	11, // 5: gophertask.gophertaskapi.v1.UpdateTaskRequestV1.status:type_name -> gophertask.gophertaskapi.v1.TaskStatus
	12, // 6: gophertask.gophertaskapi.v1.UpdateTaskRequestV1.duration:type_name -> google.protobuf.Duration
	13, // 7: gophertask.gophertaskapi.v1.UpdateTaskRequestV1.started_at:type_name -> google.protobuf.Timestamp
	0,  // 8: gophertask.gophertaskapi.v1.GetTaskByIDResponseV1.task:type_name -> gophertask.gophertaskapi.v1.TaskV1
	0,  // 9: gophertask.gophertaskapi.v1.GetTasksResponseV1.tasks:type_name -> gophertask.gophertaskapi.v1.TaskV1
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_gophertask_gophertaskapi_v1_task_proto_init() }
func file_gophertask_gophertaskapi_v1_task_proto_init() {
	if File_gophertask_gophertaskapi_v1_task_proto != nil {
		return
	}
	file_gophertask_gophertaskapi_v1_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TaskV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddTaskRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddTaskResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTaskRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTaskResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTaskResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskByIDRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetTaskByIDResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetTasksRequestV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gophertask_gophertaskapi_v1_task_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetTasksResponseV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gophertask_gophertaskapi_v1_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gophertask_gophertaskapi_v1_task_proto_goTypes,
		DependencyIndexes: file_gophertask_gophertaskapi_v1_task_proto_depIdxs,
		MessageInfos:      file_gophertask_gophertaskapi_v1_task_proto_msgTypes,
	}.Build()
	File_gophertask_gophertaskapi_v1_task_proto = out.File
	file_gophertask_gophertaskapi_v1_task_proto_rawDesc = nil
	file_gophertask_gophertaskapi_v1_task_proto_goTypes = nil
	file_gophertask_gophertaskapi_v1_task_proto_depIdxs = nil
}
