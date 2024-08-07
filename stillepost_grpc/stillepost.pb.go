// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: stillepost.proto

package stillepost_grpc

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

// The configuration that the Client tells the coordinator about
type ClientMemberConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageLen         uint32 `protobuf:"varint,1,opt,name=message_len,json=messageLen,proto3" json:"message_len,omitempty"`
	MessageCount       uint32 `protobuf:"varint,2,opt,name=message_count,json=messageCount,proto3" json:"message_count,omitempty"`
	ForwardImmediately bool   `protobuf:"varint,3,opt,name=forward_immediately,json=forwardImmediately,proto3" json:"forward_immediately,omitempty"`
	MemberTimeout      uint32 `protobuf:"varint,4,opt,name=member_timeout,json=memberTimeout,proto3" json:"member_timeout,omitempty"`
}

func (x *ClientMemberConfiguration) Reset() {
	*x = ClientMemberConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMemberConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMemberConfiguration) ProtoMessage() {}

func (x *ClientMemberConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMemberConfiguration.ProtoReflect.Descriptor instead.
func (*ClientMemberConfiguration) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMemberConfiguration) GetMessageLen() uint32 {
	if x != nil {
		return x.MessageLen
	}
	return 0
}

func (x *ClientMemberConfiguration) GetMessageCount() uint32 {
	if x != nil {
		return x.MessageCount
	}
	return 0
}

func (x *ClientMemberConfiguration) GetForwardImmediately() bool {
	if x != nil {
		return x.ForwardImmediately
	}
	return false
}

func (x *ClientMemberConfiguration) GetMemberTimeout() uint32 {
	if x != nil {
		return x.MemberTimeout
	}
	return 0
}

// The configuration that the coordinator sets on the members
type MemberConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextMember                string                     `protobuf:"bytes,1,opt,name=next_member,json=nextMember,proto3" json:"next_member,omitempty"`
	IsOrigin                  bool                       `protobuf:"varint,2,opt,name=is_origin,json=isOrigin,proto3" json:"is_origin,omitempty"`
	ClientMemberConfiguration *ClientMemberConfiguration `protobuf:"bytes,3,opt,name=client_member_configuration,json=clientMemberConfiguration,proto3" json:"client_member_configuration,omitempty"`
}

func (x *MemberConfiguration) Reset() {
	*x = MemberConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberConfiguration) ProtoMessage() {}

func (x *MemberConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberConfiguration.ProtoReflect.Descriptor instead.
func (*MemberConfiguration) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{1}
}

func (x *MemberConfiguration) GetNextMember() string {
	if x != nil {
		return x.NextMember
	}
	return ""
}

func (x *MemberConfiguration) GetIsOrigin() bool {
	if x != nil {
		return x.IsOrigin
	}
	return false
}

func (x *MemberConfiguration) GetClientMemberConfiguration() *ClientMemberConfiguration {
	if x != nil {
		return x.ClientMemberConfiguration
	}
	return nil
}

type MemberConfigurationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MemberConfigurationResult) Reset() {
	*x = MemberConfigurationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberConfigurationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberConfigurationResult) ProtoMessage() {}

func (x *MemberConfigurationResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberConfigurationResult.ProtoReflect.Descriptor instead.
func (*MemberConfigurationResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{2}
}

type StillepostOriginResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *StillepostOriginResult) Reset() {
	*x = StillepostOriginResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StillepostOriginResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StillepostOriginResult) ProtoMessage() {}

func (x *StillepostOriginResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StillepostOriginResult.ProtoReflect.Descriptor instead.
func (*StillepostOriginResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{3}
}

func (x *StillepostOriginResult) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type GetPeerEndpointParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPeerEndpointParams) Reset() {
	*x = GetPeerEndpointParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerEndpointParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerEndpointParams) ProtoMessage() {}

func (x *GetPeerEndpointParams) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerEndpointParams.ProtoReflect.Descriptor instead.
func (*GetPeerEndpointParams) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{4}
}

type GetPeerEndpointResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerEndpoint string `protobuf:"bytes,1,opt,name=peer_endpoint,json=peerEndpoint,proto3" json:"peer_endpoint,omitempty"`
}

func (x *GetPeerEndpointResult) Reset() {
	*x = GetPeerEndpointResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeerEndpointResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeerEndpointResult) ProtoMessage() {}

func (x *GetPeerEndpointResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeerEndpointResult.ProtoReflect.Descriptor instead.
func (*GetPeerEndpointResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{5}
}

func (x *GetPeerEndpointResult) GetPeerEndpoint() string {
	if x != nil {
		return x.PeerEndpoint
	}
	return ""
}

type StillepostCoordinatorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StillepostCoordinatorParams) Reset() {
	*x = StillepostCoordinatorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StillepostCoordinatorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StillepostCoordinatorParams) ProtoMessage() {}

func (x *StillepostCoordinatorParams) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StillepostCoordinatorParams.ProtoReflect.Descriptor instead.
func (*StillepostCoordinatorParams) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{6}
}

type RundgangConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RundgangId string `protobuf:"bytes,1,opt,name=rundgangId,proto3" json:"rundgangId,omitempty"`
}

func (x *RundgangConfiguration) Reset() {
	*x = RundgangConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RundgangConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RundgangConfiguration) ProtoMessage() {}

func (x *RundgangConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RundgangConfiguration.ProtoReflect.Descriptor instead.
func (*RundgangConfiguration) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{7}
}

func (x *RundgangConfiguration) GetRundgangId() string {
	if x != nil {
		return x.RundgangId
	}
	return ""
}

type RundgangConfigurationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RundgangConfigurationResult) Reset() {
	*x = RundgangConfigurationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RundgangConfigurationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RundgangConfigurationResult) ProtoMessage() {}

func (x *RundgangConfigurationResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RundgangConfigurationResult.ProtoReflect.Descriptor instead.
func (*RundgangConfigurationResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{8}
}

// Coordinator and Client RPCs
type StillepostResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *StillepostResult) Reset() {
	*x = StillepostResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StillepostResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StillepostResult) ProtoMessage() {}

func (x *StillepostResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StillepostResult.ProtoReflect.Descriptor instead.
func (*StillepostResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{9}
}

func (x *StillepostResult) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type StillepostClientParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StillepostClientParams) Reset() {
	*x = StillepostClientParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StillepostClientParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StillepostClientParams) ProtoMessage() {}

func (x *StillepostClientParams) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StillepostClientParams.ProtoReflect.Descriptor instead.
func (*StillepostClientParams) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{10}
}

type ClientRingConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster            string                     `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	ClientMemberConfig *ClientMemberConfiguration `protobuf:"bytes,2,opt,name=client_member_config,json=clientMemberConfig,proto3" json:"client_member_config,omitempty"`
}

func (x *ClientRingConfiguration) Reset() {
	*x = ClientRingConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRingConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRingConfiguration) ProtoMessage() {}

func (x *ClientRingConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRingConfiguration.ProtoReflect.Descriptor instead.
func (*ClientRingConfiguration) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{11}
}

func (x *ClientRingConfiguration) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ClientRingConfiguration) GetClientMemberConfig() *ClientMemberConfiguration {
	if x != nil {
		return x.ClientMemberConfig
	}
	return nil
}

type ClientRingConfigurationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientRingConfigurationResult) Reset() {
	*x = ClientRingConfigurationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRingConfigurationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRingConfigurationResult) ProtoMessage() {}

func (x *ClientRingConfigurationResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRingConfigurationResult.ProtoReflect.Descriptor instead.
func (*ClientRingConfigurationResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{12}
}

type ClientPingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *ClientPingParams) Reset() {
	*x = ClientPingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPingParams) ProtoMessage() {}

func (x *ClientPingParams) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPingParams.ProtoReflect.Descriptor instead.
func (*ClientPingParams) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{13}
}

func (x *ClientPingParams) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type ClientPingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results string `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
}

func (x *ClientPingResult) Reset() {
	*x = ClientPingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stillepost_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPingResult) ProtoMessage() {}

func (x *ClientPingResult) ProtoReflect() protoreflect.Message {
	mi := &file_stillepost_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPingResult.ProtoReflect.Descriptor instead.
func (*ClientPingResult) Descriptor() ([]byte, []int) {
	return file_stillepost_proto_rawDescGZIP(), []int{14}
}

func (x *ClientPingResult) GetResults() string {
	if x != nil {
		return x.Results
	}
	return ""
}

var File_stillepost_proto protoreflect.FileDescriptor

var file_stillepost_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x12, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6d, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xaf,
	0x01, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x12, 0x5a, 0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1b, 0x0a, 0x19, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a,
	0x16, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x37, 0x0a, 0x15, 0x52, 0x75, 0x6e, 0x64, 0x67, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75,
	0x6e, 0x64, 0x67, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x75, 0x6e, 0x64, 0x67, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x52, 0x75,
	0x6e, 0x64, 0x67, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x53, 0x74, 0x69,
	0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x17,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x4c, 0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x1f, 0x0a, 0x1d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x2c, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x2c,
	0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xb2, 0x02, 0x0a,
	0x10, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1a, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x75, 0x6e, 0x64, 0x67, 0x61,
	0x6e, 0x67, 0x12, 0x16, 0x2e, 0x52, 0x75, 0x6e, 0x64, 0x67, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x52, 0x75, 0x6e,
	0x64, 0x67, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x74,
	0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x53, 0x74, 0x69, 0x6c, 0x6c,
	0x65, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70,
	0x6f, 0x73, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x32, 0xda, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x4b, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x53, 0x74, 0x69, 0x6c,
	0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70,
	0x6f, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x11, 0x2e, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0e, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x6c, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1c,
	0x5a, 0x1a, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x69,
	0x6c, 0x6c, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stillepost_proto_rawDescOnce sync.Once
	file_stillepost_proto_rawDescData = file_stillepost_proto_rawDesc
)

func file_stillepost_proto_rawDescGZIP() []byte {
	file_stillepost_proto_rawDescOnce.Do(func() {
		file_stillepost_proto_rawDescData = protoimpl.X.CompressGZIP(file_stillepost_proto_rawDescData)
	})
	return file_stillepost_proto_rawDescData
}

var file_stillepost_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_stillepost_proto_goTypes = []interface{}{
	(*ClientMemberConfiguration)(nil),     // 0: ClientMemberConfiguration
	(*MemberConfiguration)(nil),           // 1: MemberConfiguration
	(*MemberConfigurationResult)(nil),     // 2: MemberConfigurationResult
	(*StillepostOriginResult)(nil),        // 3: StillepostOriginResult
	(*GetPeerEndpointParams)(nil),         // 4: GetPeerEndpointParams
	(*GetPeerEndpointResult)(nil),         // 5: GetPeerEndpointResult
	(*StillepostCoordinatorParams)(nil),   // 6: StillepostCoordinatorParams
	(*RundgangConfiguration)(nil),         // 7: RundgangConfiguration
	(*RundgangConfigurationResult)(nil),   // 8: RundgangConfigurationResult
	(*StillepostResult)(nil),              // 9: StillepostResult
	(*StillepostClientParams)(nil),        // 10: StillepostClientParams
	(*ClientRingConfiguration)(nil),       // 11: ClientRingConfiguration
	(*ClientRingConfigurationResult)(nil), // 12: ClientRingConfigurationResult
	(*ClientPingParams)(nil),              // 13: ClientPingParams
	(*ClientPingResult)(nil),              // 14: ClientPingResult
}
var file_stillepost_proto_depIdxs = []int32{
	0,  // 0: MemberConfiguration.client_member_configuration:type_name -> ClientMemberConfiguration
	0,  // 1: ClientRingConfiguration.client_member_config:type_name -> ClientMemberConfiguration
	4,  // 2: StillepostMember.GetPeerEndpoint:input_type -> GetPeerEndpointParams
	1,  // 3: StillepostMember.ConfigureMember:input_type -> MemberConfiguration
	7,  // 4: StillepostMember.ConfigureRundgang:input_type -> RundgangConfiguration
	6,  // 5: StillepostMember.Stillepost:input_type -> StillepostCoordinatorParams
	11, // 6: StillepostCoordinator.ConfigureRing:input_type -> ClientRingConfiguration
	10, // 7: StillepostCoordinator.Stillepost:input_type -> StillepostClientParams
	13, // 8: StillepostCoordinator.PingAllSystems:input_type -> ClientPingParams
	5,  // 9: StillepostMember.GetPeerEndpoint:output_type -> GetPeerEndpointResult
	2,  // 10: StillepostMember.ConfigureMember:output_type -> MemberConfigurationResult
	8,  // 11: StillepostMember.ConfigureRundgang:output_type -> RundgangConfigurationResult
	3,  // 12: StillepostMember.Stillepost:output_type -> StillepostOriginResult
	12, // 13: StillepostCoordinator.ConfigureRing:output_type -> ClientRingConfigurationResult
	9,  // 14: StillepostCoordinator.Stillepost:output_type -> StillepostResult
	14, // 15: StillepostCoordinator.PingAllSystems:output_type -> ClientPingResult
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_stillepost_proto_init() }
func file_stillepost_proto_init() {
	if File_stillepost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stillepost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMemberConfiguration); i {
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
		file_stillepost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberConfiguration); i {
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
		file_stillepost_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberConfigurationResult); i {
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
		file_stillepost_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StillepostOriginResult); i {
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
		file_stillepost_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeerEndpointParams); i {
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
		file_stillepost_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeerEndpointResult); i {
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
		file_stillepost_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StillepostCoordinatorParams); i {
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
		file_stillepost_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RundgangConfiguration); i {
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
		file_stillepost_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RundgangConfigurationResult); i {
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
		file_stillepost_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StillepostResult); i {
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
		file_stillepost_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StillepostClientParams); i {
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
		file_stillepost_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRingConfiguration); i {
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
		file_stillepost_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRingConfigurationResult); i {
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
		file_stillepost_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPingParams); i {
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
		file_stillepost_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPingResult); i {
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
			RawDescriptor: file_stillepost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_stillepost_proto_goTypes,
		DependencyIndexes: file_stillepost_proto_depIdxs,
		MessageInfos:      file_stillepost_proto_msgTypes,
	}.Build()
	File_stillepost_proto = out.File
	file_stillepost_proto_rawDesc = nil
	file_stillepost_proto_goTypes = nil
	file_stillepost_proto_depIdxs = nil
}
