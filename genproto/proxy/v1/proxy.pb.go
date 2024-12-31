// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: proxy/v1/proxy.proto

package proxyv1

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
	mi := &file_proxy_v1_proxy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[0]
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
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{0}
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

type RPCRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic    string               `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Metadata *RPCRequest_Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Payload  []byte               `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RPCRequest) Reset() {
	*x = RPCRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RPCRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRequest) ProtoMessage() {}

func (x *RPCRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRequest.ProtoReflect.Descriptor instead.
func (*RPCRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *RPCRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RPCRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RPCRequest) GetMetadata() *RPCRequest_Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RPCRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type RPCReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error   *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RPCReply) Reset() {
	*x = RPCReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RPCReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCReply) ProtoMessage() {}

func (x *RPCReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCReply.ProtoReflect.Descriptor instead.
func (*RPCReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *RPCReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RPCReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RPCReply) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username   string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ClientType string `protobuf:"bytes,3,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ClientId   string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticateRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthenticateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthenticateRequest) GetClientType() string {
	if x != nil {
		return x.ClientType
	}
	return ""
}

func (x *AuthenticateRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type AuthenticateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error    *Error    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *AuthenticateReply) Reset() {
	*x = AuthenticateReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateReply) ProtoMessage() {}

func (x *AuthenticateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateReply.ProtoReflect.Descriptor instead.
func (*AuthenticateReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{4}
}

func (x *AuthenticateReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AuthenticateReply) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Token      string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ClientType string `protobuf:"bytes,4,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ClientId   string `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserInfo) GetClientType() string {
	if x != nil {
		return x.ClientType
	}
	return ""
}

func (x *UserInfo) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SubscribeAclRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeAclRequest) Reset() {
	*x = SubscribeAclRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeAclRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeAclRequest) ProtoMessage() {}

func (x *SubscribeAclRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeAclRequest.ProtoReflect.Descriptor instead.
func (*SubscribeAclRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{6}
}

type SubscribeAclReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeAclReply) Reset() {
	*x = SubscribeAclReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeAclReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeAclReply) ProtoMessage() {}

func (x *SubscribeAclReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeAclReply.ProtoReflect.Descriptor instead.
func (*SubscribeAclReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{7}
}

type OnSubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnSubscribeRequest) Reset() {
	*x = OnSubscribeRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnSubscribeRequest) ProtoMessage() {}

func (x *OnSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnSubscribeRequest.ProtoReflect.Descriptor instead.
func (*OnSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{8}
}

type OnSubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnSubscribeReply) Reset() {
	*x = OnSubscribeReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnSubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnSubscribeReply) ProtoMessage() {}

func (x *OnSubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnSubscribeReply.ProtoReflect.Descriptor instead.
func (*OnSubscribeReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{9}
}

type OnUnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnUnsubscribeRequest) Reset() {
	*x = OnUnsubscribeRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnUnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnUnsubscribeRequest) ProtoMessage() {}

func (x *OnUnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnUnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*OnUnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{10}
}

type OnUnsubscribeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnUnsubscribeReply) Reset() {
	*x = OnUnsubscribeReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnUnsubscribeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnUnsubscribeReply) ProtoMessage() {}

func (x *OnUnsubscribeReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnUnsubscribeReply.ProtoReflect.Descriptor instead.
func (*OnUnsubscribeReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{11}
}

type OnDisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnDisconnectRequest) Reset() {
	*x = OnDisconnectRequest{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnDisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDisconnectRequest) ProtoMessage() {}

func (x *OnDisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDisconnectRequest.ProtoReflect.Descriptor instead.
func (*OnDisconnectRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{12}
}

type OnDisconnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnDisconnectReply) Reset() {
	*x = OnDisconnectReply{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnDisconnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnDisconnectReply) ProtoMessage() {}

func (x *OnDisconnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnDisconnectReply.ProtoReflect.Descriptor instead.
func (*OnDisconnectReply) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{13}
}

type RPCRequest_Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balancer string `protobuf:"bytes,1,opt,name=balancer,proto3" json:"balancer,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RPCRequest_Metadata) Reset() {
	*x = RPCRequest_Metadata{}
	mi := &file_proxy_v1_proxy_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RPCRequest_Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRequest_Metadata) ProtoMessage() {}

func (x *RPCRequest_Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRequest_Metadata.ProtoReflect.Descriptor instead.
func (*RPCRequest_Metadata) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RPCRequest_Metadata) GetBalancer() string {
	if x != nil {
		return x.Balancer
	}
	return ""
}

func (x *RPCRequest_Metadata) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_proxy_v1_proxy_proto protoreflect.FileDescriptor

var file_proxy_v1_proxy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xd4, 0x01, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x12,
	0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xca, 0x01, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x50, 0x43, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x1a, 0x38, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x64, 0x0a,
	0x08, 0x52, 0x50, 0x43, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x7d, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a,
	0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x41, 0x63, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4f, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x12, 0x0a, 0x10, 0x4f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x4f, 0x6e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4f,
	0x6e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x15, 0x0a, 0x13, 0x4f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4f, 0x6e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa7, 0x04,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x50, 0x43, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x50, 0x43, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x5c, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x5c, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6c, 0x12,
	0x26, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a,
	0x0b, 0x4f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25, 0x2e, 0x64,
	0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0d, 0x4f, 0x6e, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x27, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5c, 0x0a, 0x0c, 0x4f, 0x6e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x26, 0x2e, 0x64, 0x65, 0x65, 0x70,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x69,
	0x6f, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_v1_proxy_proto_rawDescOnce sync.Once
	file_proxy_v1_proxy_proto_rawDescData = file_proxy_v1_proxy_proto_rawDesc
)

func file_proxy_v1_proxy_proto_rawDescGZIP() []byte {
	file_proxy_v1_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_v1_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_v1_proxy_proto_rawDescData)
	})
	return file_proxy_v1_proxy_proto_rawDescData
}

var file_proxy_v1_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_proxy_v1_proxy_proto_goTypes = []any{
	(*Error)(nil),                // 0: deeploop.proxy.v1.Error
	(*RPCRequest)(nil),           // 1: deeploop.proxy.v1.RPCRequest
	(*RPCReply)(nil),             // 2: deeploop.proxy.v1.RPCReply
	(*AuthenticateRequest)(nil),  // 3: deeploop.proxy.v1.AuthenticateRequest
	(*AuthenticateReply)(nil),    // 4: deeploop.proxy.v1.AuthenticateReply
	(*UserInfo)(nil),             // 5: deeploop.proxy.v1.UserInfo
	(*SubscribeAclRequest)(nil),  // 6: deeploop.proxy.v1.SubscribeAclRequest
	(*SubscribeAclReply)(nil),    // 7: deeploop.proxy.v1.SubscribeAclReply
	(*OnSubscribeRequest)(nil),   // 8: deeploop.proxy.v1.OnSubscribeRequest
	(*OnSubscribeReply)(nil),     // 9: deeploop.proxy.v1.OnSubscribeReply
	(*OnUnsubscribeRequest)(nil), // 10: deeploop.proxy.v1.OnUnsubscribeRequest
	(*OnUnsubscribeReply)(nil),   // 11: deeploop.proxy.v1.OnUnsubscribeReply
	(*OnDisconnectRequest)(nil),  // 12: deeploop.proxy.v1.OnDisconnectRequest
	(*OnDisconnectReply)(nil),    // 13: deeploop.proxy.v1.OnDisconnectReply
	nil,                          // 14: deeploop.proxy.v1.Error.MetadataEntry
	(*RPCRequest_Metadata)(nil),  // 15: deeploop.proxy.v1.RPCRequest.Metadata
}
var file_proxy_v1_proxy_proto_depIdxs = []int32{
	14, // 0: deeploop.proxy.v1.Error.metadata:type_name -> deeploop.proxy.v1.Error.MetadataEntry
	15, // 1: deeploop.proxy.v1.RPCRequest.metadata:type_name -> deeploop.proxy.v1.RPCRequest.Metadata
	0,  // 2: deeploop.proxy.v1.RPCReply.error:type_name -> deeploop.proxy.v1.Error
	0,  // 3: deeploop.proxy.v1.AuthenticateReply.error:type_name -> deeploop.proxy.v1.Error
	5,  // 4: deeploop.proxy.v1.AuthenticateReply.user_info:type_name -> deeploop.proxy.v1.UserInfo
	1,  // 5: deeploop.proxy.v1.ProxyService.RPC:input_type -> deeploop.proxy.v1.RPCRequest
	3,  // 6: deeploop.proxy.v1.ProxyService.Authenticate:input_type -> deeploop.proxy.v1.AuthenticateRequest
	6,  // 7: deeploop.proxy.v1.ProxyService.SubscribeAcl:input_type -> deeploop.proxy.v1.SubscribeAclRequest
	8,  // 8: deeploop.proxy.v1.ProxyService.OnSubscribe:input_type -> deeploop.proxy.v1.OnSubscribeRequest
	10, // 9: deeploop.proxy.v1.ProxyService.OnUnsubscribe:input_type -> deeploop.proxy.v1.OnUnsubscribeRequest
	12, // 10: deeploop.proxy.v1.ProxyService.OnDisconnect:input_type -> deeploop.proxy.v1.OnDisconnectRequest
	2,  // 11: deeploop.proxy.v1.ProxyService.RPC:output_type -> deeploop.proxy.v1.RPCReply
	4,  // 12: deeploop.proxy.v1.ProxyService.Authenticate:output_type -> deeploop.proxy.v1.AuthenticateReply
	7,  // 13: deeploop.proxy.v1.ProxyService.SubscribeAcl:output_type -> deeploop.proxy.v1.SubscribeAclReply
	9,  // 14: deeploop.proxy.v1.ProxyService.OnSubscribe:output_type -> deeploop.proxy.v1.OnSubscribeReply
	11, // 15: deeploop.proxy.v1.ProxyService.OnUnsubscribe:output_type -> deeploop.proxy.v1.OnUnsubscribeReply
	13, // 16: deeploop.proxy.v1.ProxyService.OnDisconnect:output_type -> deeploop.proxy.v1.OnDisconnectReply
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proxy_v1_proxy_proto_init() }
func file_proxy_v1_proxy_proto_init() {
	if File_proxy_v1_proxy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_v1_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_v1_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_v1_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_v1_proxy_proto_msgTypes,
	}.Build()
	File_proxy_v1_proxy_proto = out.File
	file_proxy_v1_proxy_proto_rawDesc = nil
	file_proxy_v1_proxy_proto_goTypes = nil
	file_proxy_v1_proxy_proto_depIdxs = nil
}
