// V3 invite message structure

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: mls/message_contents/message.proto

package message_contents

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

// Wrapper for a MLS welcome message
type WelcomeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//	*WelcomeMessage_V1_
	Version isWelcomeMessage_Version `protobuf_oneof:"version"`
}

func (x *WelcomeMessage) Reset() {
	*x = WelcomeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeMessage) ProtoMessage() {}

func (x *WelcomeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeMessage.ProtoReflect.Descriptor instead.
func (*WelcomeMessage) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_message_proto_rawDescGZIP(), []int{0}
}

func (m *WelcomeMessage) GetVersion() isWelcomeMessage_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *WelcomeMessage) GetV1() *WelcomeMessage_V1 {
	if x, ok := x.GetVersion().(*WelcomeMessage_V1_); ok {
		return x.V1
	}
	return nil
}

type isWelcomeMessage_Version interface {
	isWelcomeMessage_Version()
}

type WelcomeMessage_V1_ struct {
	V1 *WelcomeMessage_V1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*WelcomeMessage_V1_) isWelcomeMessage_Version() {}

// GroupMessage wraps any MLS message sent to a group topic
type GroupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//	*GroupMessage_V1_
	Version isGroupMessage_Version `protobuf_oneof:"version"`
}

func (x *GroupMessage) Reset() {
	*x = GroupMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMessage) ProtoMessage() {}

func (x *GroupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMessage.ProtoReflect.Descriptor instead.
func (*GroupMessage) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_message_proto_rawDescGZIP(), []int{1}
}

func (m *GroupMessage) GetVersion() isGroupMessage_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *GroupMessage) GetV1() *GroupMessage_V1 {
	if x, ok := x.GetVersion().(*GroupMessage_V1_); ok {
		return x.V1
	}
	return nil
}

type isGroupMessage_Version interface {
	isGroupMessage_Version()
}

type GroupMessage_V1_ struct {
	V1 *GroupMessage_V1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*GroupMessage_V1_) isGroupMessage_Version() {}

// Version 1 of the WelcomeMessage format
type WelcomeMessage_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WelcomeMessageTlsSerialized []byte `protobuf:"bytes,1,opt,name=welcome_message_tls_serialized,json=welcomeMessageTlsSerialized,proto3" json:"welcome_message_tls_serialized,omitempty"`
}

func (x *WelcomeMessage_V1) Reset() {
	*x = WelcomeMessage_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WelcomeMessage_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WelcomeMessage_V1) ProtoMessage() {}

func (x *WelcomeMessage_V1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WelcomeMessage_V1.ProtoReflect.Descriptor instead.
func (*WelcomeMessage_V1) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WelcomeMessage_V1) GetWelcomeMessageTlsSerialized() []byte {
	if x != nil {
		return x.WelcomeMessageTlsSerialized
	}
	return nil
}

// Version 1 of the GroupMessage format
type GroupMessage_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MlsMessageTlsSerialized []byte `protobuf:"bytes,1,opt,name=mls_message_tls_serialized,json=mlsMessageTlsSerialized,proto3" json:"mls_message_tls_serialized,omitempty"`
}

func (x *GroupMessage_V1) Reset() {
	*x = GroupMessage_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_message_contents_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupMessage_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMessage_V1) ProtoMessage() {}

func (x *GroupMessage_V1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMessage_V1.ProtoReflect.Descriptor instead.
func (*GroupMessage_V1) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_message_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GroupMessage_V1) GetMlsMessageTlsSerialized() []byte {
	if x != nil {
		return x.MlsMessageTlsSerialized
	}
	return nil
}

var File_mls_message_contents_message_proto protoreflect.FileDescriptor

var file_mls_message_contents_message_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xa6, 0x01, 0x0a, 0x0e, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x3e, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x57, 0x65, 0x6c, 0x63, 0x6f,
	0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02,
	0x76, 0x31, 0x1a, 0x49, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x43, 0x0a, 0x1e, 0x77, 0x65, 0x6c, 0x63,
	0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6c, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x1b, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x09, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56,
	0x31, 0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x1a, 0x41, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x3b, 0x0a,
	0x1a, 0x6d, 0x6c, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6c, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x17, 0x6d, 0x6c, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6c, 0x73,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x57, 0x0a, 0x23, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mls_message_contents_message_proto_rawDescOnce sync.Once
	file_mls_message_contents_message_proto_rawDescData = file_mls_message_contents_message_proto_rawDesc
)

func file_mls_message_contents_message_proto_rawDescGZIP() []byte {
	file_mls_message_contents_message_proto_rawDescOnce.Do(func() {
		file_mls_message_contents_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_mls_message_contents_message_proto_rawDescData)
	})
	return file_mls_message_contents_message_proto_rawDescData
}

var file_mls_message_contents_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mls_message_contents_message_proto_goTypes = []interface{}{
	(*WelcomeMessage)(nil),    // 0: xmtp.mls.message_contents.WelcomeMessage
	(*GroupMessage)(nil),      // 1: xmtp.mls.message_contents.GroupMessage
	(*WelcomeMessage_V1)(nil), // 2: xmtp.mls.message_contents.WelcomeMessage.V1
	(*GroupMessage_V1)(nil),   // 3: xmtp.mls.message_contents.GroupMessage.V1
}
var file_mls_message_contents_message_proto_depIdxs = []int32{
	2, // 0: xmtp.mls.message_contents.WelcomeMessage.v1:type_name -> xmtp.mls.message_contents.WelcomeMessage.V1
	3, // 1: xmtp.mls.message_contents.GroupMessage.v1:type_name -> xmtp.mls.message_contents.GroupMessage.V1
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mls_message_contents_message_proto_init() }
func file_mls_message_contents_message_proto_init() {
	if File_mls_message_contents_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mls_message_contents_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeMessage); i {
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
		file_mls_message_contents_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMessage); i {
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
		file_mls_message_contents_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WelcomeMessage_V1); i {
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
		file_mls_message_contents_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupMessage_V1); i {
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
	file_mls_message_contents_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WelcomeMessage_V1_)(nil),
	}
	file_mls_message_contents_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GroupMessage_V1_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mls_message_contents_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mls_message_contents_message_proto_goTypes,
		DependencyIndexes: file_mls_message_contents_message_proto_depIdxs,
		MessageInfos:      file_mls_message_contents_message_proto_msgTypes,
	}.Build()
	File_mls_message_contents_message_proto = out.File
	file_mls_message_contents_message_proto_rawDesc = nil
	file_mls_message_contents_message_proto_goTypes = nil
	file_mls_message_contents_message_proto_depIdxs = nil
}