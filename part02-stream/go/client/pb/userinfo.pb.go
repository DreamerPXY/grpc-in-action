// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: proto/userinfo.proto

package pb

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

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *GetUserInfoRes) Reset() {
	*x = GetUserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRes) ProtoMessage() {}

func (x *GetUserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRes.ProtoReflect.Descriptor instead.
func (*GetUserInfoRes) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserInfoRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoRes) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type ReportUserBehaviorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Behavior string `protobuf:"bytes,2,opt,name=behavior,proto3" json:"behavior,omitempty"`
}

func (x *ReportUserBehaviorReq) Reset() {
	*x = ReportUserBehaviorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportUserBehaviorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUserBehaviorReq) ProtoMessage() {}

func (x *ReportUserBehaviorReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUserBehaviorReq.ProtoReflect.Descriptor instead.
func (*ReportUserBehaviorReq) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{2}
}

func (x *ReportUserBehaviorReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReportUserBehaviorReq) GetBehavior() string {
	if x != nil {
		return x.Behavior
	}
	return ""
}

type ReportUserBehaviorRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode       int32  `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg        string `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	ReceivedCount int64  `protobuf:"varint,3,opt,name=receivedCount,proto3" json:"receivedCount,omitempty"`
}

func (x *ReportUserBehaviorRes) Reset() {
	*x = ReportUserBehaviorRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportUserBehaviorRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUserBehaviorRes) ProtoMessage() {}

func (x *ReportUserBehaviorRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUserBehaviorRes.ProtoReflect.Descriptor instead.
func (*ReportUserBehaviorRes) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{3}
}

func (x *ReportUserBehaviorRes) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *ReportUserBehaviorRes) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *ReportUserBehaviorRes) GetReceivedCount() int64 {
	if x != nil {
		return x.ReceivedCount
	}
	return 0
}

type SupplyUserChangeInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SupplyUserChangeInfoReq) Reset() {
	*x = SupplyUserChangeInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyUserChangeInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyUserChangeInfoReq) ProtoMessage() {}

func (x *SupplyUserChangeInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyUserChangeInfoReq.ProtoReflect.Descriptor instead.
func (*SupplyUserChangeInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{4}
}

func (x *SupplyUserChangeInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupplyUserChangeInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Loc         string `protobuf:"bytes,2,opt,name=loc,proto3" json:"loc,omitempty"`
	Temperature string `protobuf:"bytes,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *SupplyUserChangeInfoRes) Reset() {
	*x = SupplyUserChangeInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyUserChangeInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyUserChangeInfoRes) ProtoMessage() {}

func (x *SupplyUserChangeInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyUserChangeInfoRes.ProtoReflect.Descriptor instead.
func (*SupplyUserChangeInfoRes) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{5}
}

func (x *SupplyUserChangeInfoRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupplyUserChangeInfoRes) GetLoc() string {
	if x != nil {
		return x.Loc
	}
	return ""
}

func (x *SupplyUserChangeInfoRes) GetTemperature() string {
	if x != nil {
		return x.Temperature
	}
	return ""
}

type ExchangeUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat float64 `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *ExchangeUserInfoReq) Reset() {
	*x = ExchangeUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeUserInfoReq) ProtoMessage() {}

func (x *ExchangeUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeUserInfoReq.ProtoReflect.Descriptor instead.
func (*ExchangeUserInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{6}
}

func (x *ExchangeUserInfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExchangeUserInfoReq) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *ExchangeUserInfoReq) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

type ExchangeUserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Loc string `protobuf:"bytes,2,opt,name=loc,proto3" json:"loc,omitempty"`
}

func (x *ExchangeUserInfoRes) Reset() {
	*x = ExchangeUserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_userinfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeUserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeUserInfoRes) ProtoMessage() {}

func (x *ExchangeUserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_userinfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeUserInfoRes.ProtoReflect.Descriptor instead.
func (*ExchangeUserInfoRes) Descriptor() ([]byte, []int) {
	return file_proto_userinfo_proto_rawDescGZIP(), []int{7}
}

func (x *ExchangeUserInfoRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExchangeUserInfoRes) GetLoc() string {
	if x != nil {
		return x.Loc
	}
	return ""
}

var File_proto_userinfo_proto protoreflect.FileDescriptor

var file_proto_userinfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x43, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x22, 0x6f, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5d, 0x0a, 0x17, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12,
	0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x49, 0x0a, 0x13, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22, 0x37, 0x0a, 0x13,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6c, 0x6f, 0x63, 0x32, 0x9c, 0x03, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x66, 0x0a, 0x12, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x12, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x28, 0x01, 0x12, 0x6c, 0x0a, 0x14, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x30, 0x01,
	0x12, 0x62, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x36, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x42, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x80, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_userinfo_proto_rawDescOnce sync.Once
	file_proto_userinfo_proto_rawDescData = file_proto_userinfo_proto_rawDesc
)

func file_proto_userinfo_proto_rawDescGZIP() []byte {
	file_proto_userinfo_proto_rawDescOnce.Do(func() {
		file_proto_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_userinfo_proto_rawDescData)
	})
	return file_proto_userinfo_proto_rawDescData
}

var file_proto_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_userinfo_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil),          // 0: userinfoservice.GetUserInfoReq
	(*GetUserInfoRes)(nil),          // 1: userinfoservice.GetUserInfoRes
	(*ReportUserBehaviorReq)(nil),   // 2: userinfoservice.ReportUserBehaviorReq
	(*ReportUserBehaviorRes)(nil),   // 3: userinfoservice.ReportUserBehaviorRes
	(*SupplyUserChangeInfoReq)(nil), // 4: userinfoservice.SupplyUserChangeInfoReq
	(*SupplyUserChangeInfoRes)(nil), // 5: userinfoservice.SupplyUserChangeInfoRes
	(*ExchangeUserInfoReq)(nil),     // 6: userinfoservice.ExchangeUserInfoReq
	(*ExchangeUserInfoRes)(nil),     // 7: userinfoservice.ExchangeUserInfoRes
}
var file_proto_userinfo_proto_depIdxs = []int32{
	0, // 0: userinfoservice.UserInfoService.getUserInfo:input_type -> userinfoservice.GetUserInfoReq
	2, // 1: userinfoservice.UserInfoService.reportUserBehavior:input_type -> userinfoservice.ReportUserBehaviorReq
	4, // 2: userinfoservice.UserInfoService.supplyUserChangeInfo:input_type -> userinfoservice.SupplyUserChangeInfoReq
	6, // 3: userinfoservice.UserInfoService.exchangeUserInfo:input_type -> userinfoservice.ExchangeUserInfoReq
	1, // 4: userinfoservice.UserInfoService.getUserInfo:output_type -> userinfoservice.GetUserInfoRes
	3, // 5: userinfoservice.UserInfoService.reportUserBehavior:output_type -> userinfoservice.ReportUserBehaviorRes
	5, // 6: userinfoservice.UserInfoService.supplyUserChangeInfo:output_type -> userinfoservice.SupplyUserChangeInfoRes
	7, // 7: userinfoservice.UserInfoService.exchangeUserInfo:output_type -> userinfoservice.ExchangeUserInfoRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_userinfo_proto_init() }
func file_proto_userinfo_proto_init() {
	if File_proto_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_proto_userinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRes); i {
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
		file_proto_userinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportUserBehaviorReq); i {
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
		file_proto_userinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportUserBehaviorRes); i {
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
		file_proto_userinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyUserChangeInfoReq); i {
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
		file_proto_userinfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyUserChangeInfoRes); i {
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
		file_proto_userinfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeUserInfoReq); i {
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
		file_proto_userinfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeUserInfoRes); i {
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
			RawDescriptor: file_proto_userinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_userinfo_proto_goTypes,
		DependencyIndexes: file_proto_userinfo_proto_depIdxs,
		MessageInfos:      file_proto_userinfo_proto_msgTypes,
	}.Build()
	File_proto_userinfo_proto = out.File
	file_proto_userinfo_proto_rawDesc = nil
	file_proto_userinfo_proto_goTypes = nil
	file_proto_userinfo_proto_depIdxs = nil
}
