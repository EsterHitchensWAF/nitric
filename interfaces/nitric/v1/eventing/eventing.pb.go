// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: v1/eventing.proto

package eventing

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request message contains the name to be greeted
	TopicName string `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	// The message to be published to the topic
	Event *NitricEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_eventing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_eventing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_v1_eventing_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetTopicName() string {
	if x != nil {
		return x.TopicName
	}
	return ""
}

func (x *PublishRequest) GetEvent() *NitricEvent {
	if x != nil {
		return x.Event
	}
	return nil
}

type NitricEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   string          `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	PayloadType string          `protobuf:"bytes,2,opt,name=payloadType,proto3" json:"payloadType,omitempty"`
	Payload     *_struct.Struct `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *NitricEvent) Reset() {
	*x = NitricEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_eventing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NitricEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NitricEvent) ProtoMessage() {}

func (x *NitricEvent) ProtoReflect() protoreflect.Message {
	mi := &file_v1_eventing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NitricEvent.ProtoReflect.Descriptor instead.
func (*NitricEvent) Descriptor() ([]byte, []int) {
	return file_v1_eventing_proto_rawDescGZIP(), []int{1}
}

func (x *NitricEvent) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *NitricEvent) GetPayloadType() string {
	if x != nil {
		return x.PayloadType
	}
	return ""
}

func (x *NitricEvent) GetPayload() *_struct.Struct {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GetTopicsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *GetTopicsReply) Reset() {
	*x = GetTopicsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_eventing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopicsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopicsReply) ProtoMessage() {}

func (x *GetTopicsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_eventing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopicsReply.ProtoReflect.Descriptor instead.
func (*GetTopicsReply) Descriptor() ([]byte, []int) {
	return file_v1_eventing_proto_rawDescGZIP(), []int{2}
}

func (x *GetTopicsReply) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

var File_v1_eventing_proto protoreflect.FileDescriptor

var file_v1_eventing_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x4e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x22,
	0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x22, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x3d, 0x0a, 0x15, 0x69, 0x6f, 0x2e, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x0e, 0x4e, 0x69,
	0x74, 0x72, 0x69, 0x63, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x01, 0x5a, 0x12,
	0x6e, 0x69, 0x74, 0x72, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_eventing_proto_rawDescOnce sync.Once
	file_v1_eventing_proto_rawDescData = file_v1_eventing_proto_rawDesc
)

func file_v1_eventing_proto_rawDescGZIP() []byte {
	file_v1_eventing_proto_rawDescOnce.Do(func() {
		file_v1_eventing_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_eventing_proto_rawDescData)
	})
	return file_v1_eventing_proto_rawDescData
}

var file_v1_eventing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_eventing_proto_goTypes = []interface{}{
	(*PublishRequest)(nil), // 0: nitric.v1.eventing.PublishRequest
	(*NitricEvent)(nil),    // 1: nitric.v1.eventing.NitricEvent
	(*GetTopicsReply)(nil), // 2: nitric.v1.eventing.GetTopicsReply
	(*_struct.Struct)(nil), // 3: google.protobuf.Struct
	(*empty.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_v1_eventing_proto_depIdxs = []int32{
	1, // 0: nitric.v1.eventing.PublishRequest.event:type_name -> nitric.v1.eventing.NitricEvent
	3, // 1: nitric.v1.eventing.NitricEvent.payload:type_name -> google.protobuf.Struct
	0, // 2: nitric.v1.eventing.Eventing.Publish:input_type -> nitric.v1.eventing.PublishRequest
	4, // 3: nitric.v1.eventing.Eventing.GetTopics:input_type -> google.protobuf.Empty
	4, // 4: nitric.v1.eventing.Eventing.Publish:output_type -> google.protobuf.Empty
	2, // 5: nitric.v1.eventing.Eventing.GetTopics:output_type -> nitric.v1.eventing.GetTopicsReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_eventing_proto_init() }
func file_v1_eventing_proto_init() {
	if File_v1_eventing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_eventing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_v1_eventing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NitricEvent); i {
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
		file_v1_eventing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopicsReply); i {
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
			RawDescriptor: file_v1_eventing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_eventing_proto_goTypes,
		DependencyIndexes: file_v1_eventing_proto_depIdxs,
		MessageInfos:      file_v1_eventing_proto_msgTypes,
	}.Build()
	File_v1_eventing_proto = out.File
	file_v1_eventing_proto_rawDesc = nil
	file_v1_eventing_proto_goTypes = nil
	file_v1_eventing_proto_depIdxs = nil
}
