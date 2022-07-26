// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: orderservice.proto

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

type GetOrderListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetOrderListReq) Reset() {
	*x = GetOrderListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderListReq) ProtoMessage() {}

func (x *GetOrderListReq) ProtoReflect() protoreflect.Message {
	mi := &file_orderservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderListReq.ProtoReflect.Descriptor instead.
func (*GetOrderListReq) Descriptor() ([]byte, []int) {
	return file_orderservice_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderListReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetOrderListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIds []string `protobuf:"bytes,1,rep,name=orderIds,proto3" json:"orderIds,omitempty"`
}

func (x *GetOrderListRes) Reset() {
	*x = GetOrderListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderListRes) ProtoMessage() {}

func (x *GetOrderListRes) ProtoReflect() protoreflect.Message {
	mi := &file_orderservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderListRes.ProtoReflect.Descriptor instead.
func (*GetOrderListRes) Descriptor() ([]byte, []int) {
	return file_orderservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderListRes) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

type SearchOrderListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderName string `protobuf:"bytes,1,opt,name=orderName,proto3" json:"orderName,omitempty"`
}

func (x *SearchOrderListReq) Reset() {
	*x = SearchOrderListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrderListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrderListReq) ProtoMessage() {}

func (x *SearchOrderListReq) ProtoReflect() protoreflect.Message {
	mi := &file_orderservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrderListReq.ProtoReflect.Descriptor instead.
func (*SearchOrderListReq) Descriptor() ([]byte, []int) {
	return file_orderservice_proto_rawDescGZIP(), []int{2}
}

func (x *SearchOrderListReq) GetOrderName() string {
	if x != nil {
		return x.OrderName
	}
	return ""
}

type SearchOrderListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderIds []string `protobuf:"bytes,1,rep,name=orderIds,proto3" json:"orderIds,omitempty"`
}

func (x *SearchOrderListRes) Reset() {
	*x = SearchOrderListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrderListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrderListRes) ProtoMessage() {}

func (x *SearchOrderListRes) ProtoReflect() protoreflect.Message {
	mi := &file_orderservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrderListRes.ProtoReflect.Descriptor instead.
func (*SearchOrderListRes) Descriptor() ([]byte, []int) {
	return file_orderservice_proto_rawDescGZIP(), []int{3}
}

func (x *SearchOrderListRes) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

var File_orderservice_proto protoreflect.FileDescriptor

var file_orderservice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x32, 0x0a, 0x12,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x30, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x32, 0x5c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x42, 0x32, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0d, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x62, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x80, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderservice_proto_rawDescOnce sync.Once
	file_orderservice_proto_rawDescData = file_orderservice_proto_rawDesc
)

func file_orderservice_proto_rawDescGZIP() []byte {
	file_orderservice_proto_rawDescOnce.Do(func() {
		file_orderservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderservice_proto_rawDescData)
	})
	return file_orderservice_proto_rawDescData
}

var file_orderservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orderservice_proto_goTypes = []interface{}{
	(*GetOrderListReq)(nil),    // 0: orderservice.GetOrderListReq
	(*GetOrderListRes)(nil),    // 1: orderservice.GetOrderListRes
	(*SearchOrderListReq)(nil), // 2: orderservice.SearchOrderListReq
	(*SearchOrderListRes)(nil), // 3: orderservice.SearchOrderListRes
}
var file_orderservice_proto_depIdxs = []int32{
	0, // 0: orderservice.OrderService.getOrderList:input_type -> orderservice.GetOrderListReq
	1, // 1: orderservice.OrderService.getOrderList:output_type -> orderservice.GetOrderListRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderservice_proto_init() }
func file_orderservice_proto_init() {
	if File_orderservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderListReq); i {
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
		file_orderservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderListRes); i {
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
		file_orderservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrderListReq); i {
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
		file_orderservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrderListRes); i {
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
			RawDescriptor: file_orderservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderservice_proto_goTypes,
		DependencyIndexes: file_orderservice_proto_depIdxs,
		MessageInfos:      file_orderservice_proto_msgTypes,
	}.Build()
	File_orderservice_proto = out.File
	file_orderservice_proto_rawDesc = nil
	file_orderservice_proto_goTypes = nil
	file_orderservice_proto_depIdxs = nil
}
