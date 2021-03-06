// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: pkg/timeservice/timeservice.proto

package timeservice

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Response int32

const (
	Response_SUCCESS Response = 0
	Response_FAILURE Response = 1
	Response_ERROR   Response = 2
)

// Enum value maps for Response.
var (
	Response_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
		2: "ERROR",
	}
	Response_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
		"ERROR":   2,
	}
)

func (x Response) Enum() *Response {
	p := new(Response)
	*p = x
	return p
}

func (x Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_timeservice_timeservice_proto_enumTypes[0].Descriptor()
}

func (Response) Type() protoreflect.EnumType {
	return &file_pkg_timeservice_timeservice_proto_enumTypes[0]
}

func (x Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response.Descriptor instead.
func (Response) EnumDescriptor() ([]byte, []int) {
	return file_pkg_timeservice_timeservice_proto_rawDescGZIP(), []int{0}
}

// Request for sending current clock time.
type ClockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=CurrTime,proto3" json:"CurrTime,omitempty"`
	MsgId    uint64               `protobuf:"varint,2,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	Latency  uint64               `protobuf:"varint,3,opt,name=latency,proto3" json:"latency,omitempty"`
	Label    string               `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *ClockRequest) Reset() {
	*x = ClockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_timeservice_timeservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockRequest) ProtoMessage() {}

func (x *ClockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_timeservice_timeservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockRequest.ProtoReflect.Descriptor instead.
func (*ClockRequest) Descriptor() ([]byte, []int) {
	return file_pkg_timeservice_timeservice_proto_rawDescGZIP(), []int{0}
}

func (x *ClockRequest) GetCurrTime() *timestamp.Timestamp {
	if x != nil {
		return x.CurrTime
	}
	return nil
}

func (x *ClockRequest) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *ClockRequest) GetLatency() uint64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *ClockRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Response to a clock request.
type ClockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgId  uint64   `protobuf:"varint,1,opt,name=MsgId,proto3" json:"MsgId,omitempty"`
	Resp   Response `protobuf:"varint,2,opt,name=resp,proto3,enum=timeservice.Response" json:"resp,omitempty"`
	ErrMsg string   `protobuf:"bytes,3,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
}

func (x *ClockReply) Reset() {
	*x = ClockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_timeservice_timeservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockReply) ProtoMessage() {}

func (x *ClockReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_timeservice_timeservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockReply.ProtoReflect.Descriptor instead.
func (*ClockReply) Descriptor() ([]byte, []int) {
	return file_pkg_timeservice_timeservice_proto_rawDescGZIP(), []int{1}
}

func (x *ClockReply) GetMsgId() uint64 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *ClockReply) GetResp() Response {
	if x != nil {
		return x.Resp
	}
	return Response_SUCCESS
}

func (x *ClockReply) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

var File_pkg_timeservice_timeservice_proto protoreflect.FileDescriptor

var file_pkg_timeservice_timeservice_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x43, 0x75, 0x72, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73,
	0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x22, 0x65, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x2a, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0x50, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x6b,
	0x67, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_timeservice_timeservice_proto_rawDescOnce sync.Once
	file_pkg_timeservice_timeservice_proto_rawDescData = file_pkg_timeservice_timeservice_proto_rawDesc
)

func file_pkg_timeservice_timeservice_proto_rawDescGZIP() []byte {
	file_pkg_timeservice_timeservice_proto_rawDescOnce.Do(func() {
		file_pkg_timeservice_timeservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_timeservice_timeservice_proto_rawDescData)
	})
	return file_pkg_timeservice_timeservice_proto_rawDescData
}

var file_pkg_timeservice_timeservice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_timeservice_timeservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_timeservice_timeservice_proto_goTypes = []interface{}{
	(Response)(0),               // 0: timeservice.Response
	(*ClockRequest)(nil),        // 1: timeservice.ClockRequest
	(*ClockReply)(nil),          // 2: timeservice.ClockReply
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pkg_timeservice_timeservice_proto_depIdxs = []int32{
	3, // 0: timeservice.ClockRequest.CurrTime:type_name -> google.protobuf.Timestamp
	0, // 1: timeservice.ClockReply.resp:type_name -> timeservice.Response
	1, // 2: timeservice.TimeService.SendClock:input_type -> timeservice.ClockRequest
	2, // 3: timeservice.TimeService.SendClock:output_type -> timeservice.ClockReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_timeservice_timeservice_proto_init() }
func file_pkg_timeservice_timeservice_proto_init() {
	if File_pkg_timeservice_timeservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_timeservice_timeservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockRequest); i {
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
		file_pkg_timeservice_timeservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockReply); i {
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
			RawDescriptor: file_pkg_timeservice_timeservice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_timeservice_timeservice_proto_goTypes,
		DependencyIndexes: file_pkg_timeservice_timeservice_proto_depIdxs,
		EnumInfos:         file_pkg_timeservice_timeservice_proto_enumTypes,
		MessageInfos:      file_pkg_timeservice_timeservice_proto_msgTypes,
	}.Build()
	File_pkg_timeservice_timeservice_proto = out.File
	file_pkg_timeservice_timeservice_proto_rawDesc = nil
	file_pkg_timeservice_timeservice_proto_goTypes = nil
	file_pkg_timeservice_timeservice_proto_depIdxs = nil
}
