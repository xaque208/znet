// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: rpc/rpc.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
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

type MACObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mac        string `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Ip         string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	SourceHost string `protobuf:"bytes,3,opt,name=source_host,json=sourceHost,proto3" json:"source_host,omitempty"`
}

func (x *MACObservation) Reset() {
	*x = MACObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MACObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MACObservation) ProtoMessage() {}

func (x *MACObservation) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MACObservation.ProtoReflect.Descriptor instead.
func (*MACObservation) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *MACObservation) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *MACObservation) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *MACObservation) GetSourceHost() string {
	if x != nil {
		return x.SourceHost
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors  bool   `protobuf:"varint,1,opt,name=errors,proto3" json:"errors,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *EventResponse) GetErrors() bool {
	if x != nil {
		return x.Errors
	}
	return false
}

func (x *EventResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EventSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventNames []string `protobuf:"bytes,1,rep,name=event_names,json=eventNames,proto3" json:"event_names,omitempty"`
}

func (x *EventSub) Reset() {
	*x = EventSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSub) ProtoMessage() {}

func (x *EventSub) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSub.ProtoReflect.Descriptor instead.
func (*EventSub) Descriptor() ([]byte, []int) {
	return file_rpc_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *EventSub) GetEventNames() []string {
	if x != nil {
		return x.EventNames
	}
	return nil
}

var File_rpc_rpc_proto protoreflect.FileDescriptor

var file_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x72, 0x70, 0x63, 0x22, 0x53, 0x0a, 0x0e, 0x4d, 0x41, 0x43, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x08,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x67, 0x0a, 0x06, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x12,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x75, 0x62, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_rpc_proto_rawDescOnce sync.Once
	file_rpc_rpc_proto_rawDescData = file_rpc_rpc_proto_rawDesc
)

func file_rpc_rpc_proto_rawDescGZIP() []byte {
	file_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_rpc_proto_rawDescData)
	})
	return file_rpc_rpc_proto_rawDescData
}

var file_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_rpc_proto_goTypes = []interface{}{
	(*MACObservation)(nil), // 0: rpc.MACObservation
	(*Empty)(nil),          // 1: rpc.Empty
	(*Event)(nil),          // 2: rpc.Event
	(*EventResponse)(nil),  // 3: rpc.EventResponse
	(*EventSub)(nil),       // 4: rpc.EventSub
}
var file_rpc_rpc_proto_depIdxs = []int32{
	2, // 0: rpc.Events.NoticeEvent:input_type -> rpc.Event
	4, // 1: rpc.Events.SubscribeEvents:input_type -> rpc.EventSub
	3, // 2: rpc.Events.NoticeEvent:output_type -> rpc.EventResponse
	2, // 3: rpc.Events.SubscribeEvents:output_type -> rpc.Event
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_rpc_proto_init() }
func file_rpc_rpc_proto_init() {
	if File_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MACObservation); i {
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
		file_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_rpc_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_rpc_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_rpc_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSub); i {
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
			RawDescriptor: file_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_rpc_proto_msgTypes,
	}.Build()
	File_rpc_rpc_proto = out.File
	file_rpc_rpc_proto_rawDesc = nil
	file_rpc_rpc_proto_goTypes = nil
	file_rpc_rpc_proto_depIdxs = nil
}
