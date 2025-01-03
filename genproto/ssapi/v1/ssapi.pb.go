// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: ssapi/v1/ssapi.proto

package ssapiv1

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

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic   string            `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Type    string            `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Headers map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body    []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[0]
	if x != nil {
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
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PublishRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *PublishRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type PublishReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishReply) Reset() {
	*x = PublishReply{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReply) ProtoMessage() {}

func (x *PublishReply) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReply.ProtoReflect.Descriptor instead.
func (*PublishReply) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{1}
}

type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{2}
}

type DisconnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectReply) Reset() {
	*x = DisconnectReply{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DisconnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectReply) ProtoMessage() {}

func (x *DisconnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectReply.ProtoReflect.Descriptor instead.
func (*DisconnectReply) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{3}
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{4}
}

type SubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeReply) Reset() {
	*x = SubscribeReply{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeReply) ProtoMessage() {}

func (x *SubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeReply.ProtoReflect.Descriptor instead.
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{5}
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{6}
}

type UnsubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeReply) Reset() {
	*x = UnsubscribeReply{}
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeReply) ProtoMessage() {}

func (x *UnsubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_ssapi_v1_ssapi_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeReply.ProtoReflect.Descriptor instead.
func (*UnsubscribeReply) Descriptor() ([]byte, []int) {
	return file_ssapi_v1_ssapi_proto_rawDescGZIP(), []int{7}
}

var File_ssapi_v1_ssapi_proto protoreflect.FileDescriptor

var file_ssapi_v1_ssapi_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xe4, 0x01, 0x0a, 0x0e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x13, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xe1, 0x02, 0x0a, 0x0a, 0x41, 0x50, 0x49,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x21, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x73,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x56, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x65, 0x65,
	0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x23, 0x2e, 0x64, 0x65,
	0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x73, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x23, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x73, 0x73,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x44, 0x5a, 0x42,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c,
	0x6f, 0x6f, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ssapi_v1_ssapi_proto_rawDescOnce sync.Once
	file_ssapi_v1_ssapi_proto_rawDescData = file_ssapi_v1_ssapi_proto_rawDesc
)

func file_ssapi_v1_ssapi_proto_rawDescGZIP() []byte {
	file_ssapi_v1_ssapi_proto_rawDescOnce.Do(func() {
		file_ssapi_v1_ssapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssapi_v1_ssapi_proto_rawDescData)
	})
	return file_ssapi_v1_ssapi_proto_rawDescData
}

var file_ssapi_v1_ssapi_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ssapi_v1_ssapi_proto_goTypes = []any{
	(*PublishRequest)(nil),     // 0: deeploop.ssapi.v1.PublishRequest
	(*PublishReply)(nil),       // 1: deeploop.ssapi.v1.PublishReply
	(*DisconnectRequest)(nil),  // 2: deeploop.ssapi.v1.DisconnectRequest
	(*DisconnectReply)(nil),    // 3: deeploop.ssapi.v1.DisconnectReply
	(*SubscribeRequest)(nil),   // 4: deeploop.ssapi.v1.SubscribeRequest
	(*SubscribeReply)(nil),     // 5: deeploop.ssapi.v1.SubscribeReply
	(*UnsubscribeRequest)(nil), // 6: deeploop.ssapi.v1.UnsubscribeRequest
	(*UnsubscribeReply)(nil),   // 7: deeploop.ssapi.v1.UnsubscribeReply
	nil,                        // 8: deeploop.ssapi.v1.PublishRequest.HeadersEntry
}
var file_ssapi_v1_ssapi_proto_depIdxs = []int32{
	8, // 0: deeploop.ssapi.v1.PublishRequest.headers:type_name -> deeploop.ssapi.v1.PublishRequest.HeadersEntry
	0, // 1: deeploop.ssapi.v1.APIService.Publish:input_type -> deeploop.ssapi.v1.PublishRequest
	2, // 2: deeploop.ssapi.v1.APIService.Disconnect:input_type -> deeploop.ssapi.v1.DisconnectRequest
	4, // 3: deeploop.ssapi.v1.APIService.Subscribe:input_type -> deeploop.ssapi.v1.SubscribeRequest
	4, // 4: deeploop.ssapi.v1.APIService.Unsubscribe:input_type -> deeploop.ssapi.v1.SubscribeRequest
	1, // 5: deeploop.ssapi.v1.APIService.Publish:output_type -> deeploop.ssapi.v1.PublishReply
	3, // 6: deeploop.ssapi.v1.APIService.Disconnect:output_type -> deeploop.ssapi.v1.DisconnectReply
	5, // 7: deeploop.ssapi.v1.APIService.Subscribe:output_type -> deeploop.ssapi.v1.SubscribeReply
	7, // 8: deeploop.ssapi.v1.APIService.Unsubscribe:output_type -> deeploop.ssapi.v1.UnsubscribeReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ssapi_v1_ssapi_proto_init() }
func file_ssapi_v1_ssapi_proto_init() {
	if File_ssapi_v1_ssapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssapi_v1_ssapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ssapi_v1_ssapi_proto_goTypes,
		DependencyIndexes: file_ssapi_v1_ssapi_proto_depIdxs,
		MessageInfos:      file_ssapi_v1_ssapi_proto_msgTypes,
	}.Build()
	File_ssapi_v1_ssapi_proto = out.File
	file_ssapi_v1_ssapi_proto_rawDesc = nil
	file_ssapi_v1_ssapi_proto_goTypes = nil
	file_ssapi_v1_ssapi_proto_depIdxs = nil
}
