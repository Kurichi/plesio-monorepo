// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: schema/mission.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string    `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Target      string    `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Amount      int32     `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit        string    `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Term        string    `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty"`
	Rewards     []*Reward `protobuf:"bytes,6,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *CreateMissionRequest) Reset() {
	*x = CreateMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMissionRequest) ProtoMessage() {}

func (x *CreateMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMissionRequest.ProtoReflect.Descriptor instead.
func (*CreateMissionRequest) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMissionRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMissionRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *CreateMissionRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateMissionRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreateMissionRequest) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *CreateMissionRequest) GetRewards() []*Reward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type CreateMissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mission *SimpleMission `protobuf:"bytes,1,opt,name=mission,proto3" json:"mission,omitempty"`
}

func (x *CreateMissionResponse) Reset() {
	*x = CreateMissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMissionResponse) ProtoMessage() {}

func (x *CreateMissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMissionResponse.ProtoReflect.Descriptor instead.
func (*CreateMissionResponse) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMissionResponse) GetMission() *SimpleMission {
	if x != nil {
		return x.Mission
	}
	return nil
}

type GetMissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Term   *string `protobuf:"bytes,2,opt,name=term,proto3,oneof" json:"term,omitempty"`
}

func (x *GetMissionsRequest) Reset() {
	*x = GetMissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionsRequest) ProtoMessage() {}

func (x *GetMissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionsRequest.ProtoReflect.Descriptor instead.
func (*GetMissionsRequest) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{2}
}

func (x *GetMissionsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetMissionsRequest) GetTerm() string {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return ""
}

type GetMissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Missions []*Mission `protobuf:"bytes,1,rep,name=missions,proto3" json:"missions,omitempty"`
}

func (x *GetMissionsResponse) Reset() {
	*x = GetMissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionsResponse) ProtoMessage() {}

func (x *GetMissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionsResponse.ProtoReflect.Descriptor instead.
func (*GetMissionsResponse) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{3}
}

func (x *GetMissionsResponse) GetMissions() []*Mission {
	if x != nil {
		return x.Missions
	}
	return nil
}

type ProgressMissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MissionId string `protobuf:"bytes,2,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	Progress  int32  `protobuf:"varint,3,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *ProgressMissionRequest) Reset() {
	*x = ProgressMissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressMissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressMissionRequest) ProtoMessage() {}

func (x *ProgressMissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressMissionRequest.ProtoReflect.Descriptor instead.
func (*ProgressMissionRequest) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{4}
}

func (x *ProgressMissionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProgressMissionRequest) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

func (x *ProgressMissionRequest) GetProgress() int32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type ProgressMissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mission     *Mission `protobuf:"bytes,1,opt,name=mission,proto3" json:"mission,omitempty"`
	IsCompleted bool     `protobuf:"varint,2,opt,name=is_completed,json=isCompleted,proto3" json:"is_completed,omitempty"`
}

func (x *ProgressMissionResponse) Reset() {
	*x = ProgressMissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressMissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressMissionResponse) ProtoMessage() {}

func (x *ProgressMissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressMissionResponse.ProtoReflect.Descriptor instead.
func (*ProgressMissionResponse) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{5}
}

func (x *ProgressMissionResponse) GetMission() *Mission {
	if x != nil {
		return x.Mission
	}
	return nil
}

func (x *ProgressMissionResponse) GetIsCompleted() bool {
	if x != nil {
		return x.IsCompleted
	}
	return false
}

type SimpleMission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Target      string    `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Amount      int32     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit        string    `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	Term        string    `protobuf:"bytes,6,opt,name=term,proto3" json:"term,omitempty"`
	Rewards     []*Reward `protobuf:"bytes,7,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *SimpleMission) Reset() {
	*x = SimpleMission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMission) ProtoMessage() {}

func (x *SimpleMission) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMission.ProtoReflect.Descriptor instead.
func (*SimpleMission) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{6}
}

func (x *SimpleMission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimpleMission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SimpleMission) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SimpleMission) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *SimpleMission) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *SimpleMission) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *SimpleMission) GetRewards() []*Reward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type Mission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Target      string    `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Amount      int32     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Unit        string    `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	Term        string    `protobuf:"bytes,6,opt,name=term,proto3" json:"term,omitempty"`
	Rewards     []*Reward `protobuf:"bytes,7,rep,name=rewards,proto3" json:"rewards,omitempty"`
	Progress    uint32    `protobuf:"varint,8,opt,name=progress,proto3" json:"progress,omitempty"`
	Deadline    int64     `protobuf:"varint,9,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *Mission) Reset() {
	*x = Mission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mission) ProtoMessage() {}

func (x *Mission) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mission.ProtoReflect.Descriptor instead.
func (*Mission) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{7}
}

func (x *Mission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Mission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Mission) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Mission) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Mission) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Mission) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *Mission) GetRewards() []*Reward {
	if x != nil {
		return x.Rewards
	}
	return nil
}

func (x *Mission) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Mission) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_mission_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

func (x *Reward) ProtoReflect() protoreflect.Message {
	mi := &file_schema_mission_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_schema_mission_proto_rawDescGZIP(), []int{8}
}

func (x *Reward) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Reward) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_schema_mission_proto protoreflect.FileDescriptor

var file_schema_mission_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0xbb, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0x49, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6c,
	0x0a, 0x16, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x68, 0x0a, 0x17,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x72, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0xf6, 0x01,
	0x0a, 0x07, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x86, 0x02, 0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_mission_proto_rawDescOnce sync.Once
	file_schema_mission_proto_rawDescData = file_schema_mission_proto_rawDesc
)

func file_schema_mission_proto_rawDescGZIP() []byte {
	file_schema_mission_proto_rawDescOnce.Do(func() {
		file_schema_mission_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_mission_proto_rawDescData)
	})
	return file_schema_mission_proto_rawDescData
}

var file_schema_mission_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_schema_mission_proto_goTypes = []interface{}{
	(*CreateMissionRequest)(nil),    // 0: mission.CreateMissionRequest
	(*CreateMissionResponse)(nil),   // 1: mission.CreateMissionResponse
	(*GetMissionsRequest)(nil),      // 2: mission.GetMissionsRequest
	(*GetMissionsResponse)(nil),     // 3: mission.GetMissionsResponse
	(*ProgressMissionRequest)(nil),  // 4: mission.ProgressMissionRequest
	(*ProgressMissionResponse)(nil), // 5: mission.ProgressMissionResponse
	(*SimpleMission)(nil),           // 6: mission.SimpleMission
	(*Mission)(nil),                 // 7: mission.Mission
	(*Reward)(nil),                  // 8: mission.Reward
}
var file_schema_mission_proto_depIdxs = []int32{
	8, // 0: mission.CreateMissionRequest.rewards:type_name -> mission.Reward
	6, // 1: mission.CreateMissionResponse.mission:type_name -> mission.SimpleMission
	7, // 2: mission.GetMissionsResponse.missions:type_name -> mission.Mission
	7, // 3: mission.ProgressMissionResponse.mission:type_name -> mission.Mission
	8, // 4: mission.SimpleMission.rewards:type_name -> mission.Reward
	8, // 5: mission.Mission.rewards:type_name -> mission.Reward
	0, // 6: mission.MissionService.CreateMission:input_type -> mission.CreateMissionRequest
	2, // 7: mission.MissionService.GetMissions:input_type -> mission.GetMissionsRequest
	4, // 8: mission.MissionService.ProgressMission:input_type -> mission.ProgressMissionRequest
	1, // 9: mission.MissionService.CreateMission:output_type -> mission.CreateMissionResponse
	3, // 10: mission.MissionService.GetMissions:output_type -> mission.GetMissionsResponse
	5, // 11: mission.MissionService.ProgressMission:output_type -> mission.ProgressMissionResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_schema_mission_proto_init() }
func file_schema_mission_proto_init() {
	if File_schema_mission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_mission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMissionRequest); i {
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
		file_schema_mission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMissionResponse); i {
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
		file_schema_mission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionsRequest); i {
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
		file_schema_mission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionsResponse); i {
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
		file_schema_mission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressMissionRequest); i {
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
		file_schema_mission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressMissionResponse); i {
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
		file_schema_mission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMission); i {
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
		file_schema_mission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mission); i {
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
		file_schema_mission_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
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
	file_schema_mission_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_mission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_mission_proto_goTypes,
		DependencyIndexes: file_schema_mission_proto_depIdxs,
		MessageInfos:      file_schema_mission_proto_msgTypes,
	}.Build()
	File_schema_mission_proto = out.File
	file_schema_mission_proto_rawDesc = nil
	file_schema_mission_proto_goTypes = nil
	file_schema_mission_proto_depIdxs = nil
}
