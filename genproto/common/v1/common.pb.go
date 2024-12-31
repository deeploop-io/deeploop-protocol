// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: common/v1/common.proto

package commonv1

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

type QoS int32

const (
	QoS_AtMostOnce  QoS = 0
	QoS_AtLeastOnce QoS = 1
)

// Enum value maps for QoS.
var (
	QoS_name = map[int32]string{
		0: "AtMostOnce",
		1: "AtLeastOnce",
	}
	QoS_value = map[string]int32{
		"AtMostOnce":  0,
		"AtLeastOnce": 1,
	}
)

func (x QoS) Enum() *QoS {
	p := new(QoS)
	*p = x
	return p
}

func (x QoS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QoS) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (QoS) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[0]
}

func (x QoS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QoS.Descriptor instead.
func (QoS) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message   string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Temporary bool              `protobuf:"varint,3,opt,name=temporary,proto3" json:"temporary,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_common_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetTemporary() bool {
	if x != nil {
		return x.Temporary
	}
	return false
}

func (x *Error) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_common_v1_common_proto protoreflect.FileDescriptor

var file_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xd5, 0x01, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x72, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x2a, 0x26, 0x0a, 0x03, 0x51, 0x6f, 0x53, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x74, 0x4d, 0x6f, 0x73, 0x74, 0x4f, 0x6e, 0x63, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x74, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x4f, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c,
	0x6f, 0x6f, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_common_proto_rawDescOnce sync.Once
	file_common_v1_common_proto_rawDescData = file_common_v1_common_proto_rawDesc
)

func file_common_v1_common_proto_rawDescGZIP() []byte {
	file_common_v1_common_proto_rawDescOnce.Do(func() {
		file_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_common_proto_rawDescData)
	})
	return file_common_v1_common_proto_rawDescData
}

var file_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_v1_common_proto_goTypes = []any{
	(QoS)(0),      // 0: deeploop.common.v1.QoS
	(*Error)(nil), // 1: deeploop.common.v1.Error
	nil,           // 2: deeploop.common.v1.Error.MetadataEntry
}
var file_common_v1_common_proto_depIdxs = []int32{
	2, // 0: deeploop.common.v1.Error.metadata:type_name -> deeploop.common.v1.Error.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_v1_common_proto_init() }
func file_common_v1_common_proto_init() {
	if File_common_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_common_proto_goTypes,
		DependencyIndexes: file_common_v1_common_proto_depIdxs,
		EnumInfos:         file_common_v1_common_proto_enumTypes,
		MessageInfos:      file_common_v1_common_proto_msgTypes,
	}.Build()
	File_common_v1_common_proto = out.File
	file_common_v1_common_proto_rawDesc = nil
	file_common_v1_common_proto_goTypes = nil
	file_common_v1_common_proto_depIdxs = nil
}