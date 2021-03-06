//
// Copyright 2019, Offchain Labs, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.1
// source: valmessage.proto

package valmessage

import (
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	valprotocol "github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
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

type TokenTypeBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TokenTypeBuf) Reset() {
	*x = TokenTypeBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenTypeBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenTypeBuf) ProtoMessage() {}

func (x *TokenTypeBuf) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenTypeBuf.ProtoReflect.Descriptor instead.
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{0}
}

func (x *TokenTypeBuf) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type VMConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GracePeriod           uint64                `protobuf:"varint,1,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	EscrowRequired        *common.BigIntegerBuf `protobuf:"bytes,2,opt,name=escrow_required,json=escrowRequired,proto3" json:"escrow_required,omitempty"`
	EscrowCurrency        *common.AddressBuf    `protobuf:"bytes,3,opt,name=escrow_currency,json=escrowCurrency,proto3" json:"escrow_currency,omitempty"`
	AssertKeys            []*common.AddressBuf  `protobuf:"bytes,4,rep,name=assert_keys,json=assertKeys,proto3" json:"assert_keys,omitempty"`
	MaxExecutionStepCount uint32                `protobuf:"varint,5,opt,name=max_execution_step_count,json=maxExecutionStepCount,proto3" json:"max_execution_step_count,omitempty"`
	Owner                 *common.AddressBuf    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *VMConfiguration) Reset() {
	*x = VMConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VMConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VMConfiguration) ProtoMessage() {}

func (x *VMConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VMConfiguration.ProtoReflect.Descriptor instead.
func (*VMConfiguration) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{1}
}

func (x *VMConfiguration) GetGracePeriod() uint64 {
	if x != nil {
		return x.GracePeriod
	}
	return 0
}

func (x *VMConfiguration) GetEscrowRequired() *common.BigIntegerBuf {
	if x != nil {
		return x.EscrowRequired
	}
	return nil
}

func (x *VMConfiguration) GetEscrowCurrency() *common.AddressBuf {
	if x != nil {
		return x.EscrowCurrency
	}
	return nil
}

func (x *VMConfiguration) GetAssertKeys() []*common.AddressBuf {
	if x != nil {
		return x.AssertKeys
	}
	return nil
}

func (x *VMConfiguration) GetMaxExecutionStepCount() uint32 {
	if x != nil {
		return x.MaxExecutionStepCount
	}
	return 0
}

func (x *VMConfiguration) GetOwner() *common.AddressBuf {
	if x != nil {
		return x.Owner
	}
	return nil
}

type UnanimousAssertionValidatorNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted   bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signatures [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *UnanimousAssertionValidatorNotification) Reset() {
	*x = UnanimousAssertionValidatorNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnanimousAssertionValidatorNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnanimousAssertionValidatorNotification) ProtoMessage() {}

func (x *UnanimousAssertionValidatorNotification) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnanimousAssertionValidatorNotification.ProtoReflect.Descriptor instead.
func (*UnanimousAssertionValidatorNotification) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{2}
}

func (x *UnanimousAssertionValidatorNotification) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *UnanimousAssertionValidatorNotification) GetSignatures() [][]byte {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type SignedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *valprotocol.MessageBuf `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte                  `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedMessage) Reset() {
	*x = SignedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedMessage) ProtoMessage() {}

func (x *SignedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedMessage.ProtoReflect.Descriptor instead.
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{3}
}

func (x *SignedMessage) GetMessage() *valprotocol.MessageBuf {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type UnanimousAssertionValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeforeHash     *common.HashBuf               `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	BeforeInbox    *common.HashBuf               `protobuf:"bytes,2,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
	SequenceNum    uint64                        `protobuf:"varint,3,opt,name=sequenceNum,proto3" json:"sequenceNum,omitempty"`
	TimeBounds     *protocol.TimeBoundsBlocksBuf `protobuf:"bytes,4,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	SignedMessages []*SignedMessage              `protobuf:"bytes,5,rep,name=signedMessages,proto3" json:"signedMessages,omitempty"`
}

func (x *UnanimousAssertionValidatorRequest) Reset() {
	*x = UnanimousAssertionValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnanimousAssertionValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnanimousAssertionValidatorRequest) ProtoMessage() {}

func (x *UnanimousAssertionValidatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnanimousAssertionValidatorRequest.ProtoReflect.Descriptor instead.
func (*UnanimousAssertionValidatorRequest) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{4}
}

func (x *UnanimousAssertionValidatorRequest) GetBeforeHash() *common.HashBuf {
	if x != nil {
		return x.BeforeHash
	}
	return nil
}

func (x *UnanimousAssertionValidatorRequest) GetBeforeInbox() *common.HashBuf {
	if x != nil {
		return x.BeforeInbox
	}
	return nil
}

func (x *UnanimousAssertionValidatorRequest) GetSequenceNum() uint64 {
	if x != nil {
		return x.SequenceNum
	}
	return 0
}

func (x *UnanimousAssertionValidatorRequest) GetTimeBounds() *protocol.TimeBoundsBlocksBuf {
	if x != nil {
		return x.TimeBounds
	}
	return nil
}

func (x *UnanimousAssertionValidatorRequest) GetSignedMessages() []*SignedMessage {
	if x != nil {
		return x.SignedMessages
	}
	return nil
}

type ValidatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *common.HashBuf `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are assignable to Request:
	//	*ValidatorRequest_Unanimous
	//	*ValidatorRequest_UnanimousNotification
	Request isValidatorRequest_Request `protobuf_oneof:"request"`
}

func (x *ValidatorRequest) Reset() {
	*x = ValidatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorRequest) ProtoMessage() {}

func (x *ValidatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorRequest.ProtoReflect.Descriptor instead.
func (*ValidatorRequest) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{5}
}

func (x *ValidatorRequest) GetRequestId() *common.HashBuf {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (m *ValidatorRequest) GetRequest() isValidatorRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ValidatorRequest) GetUnanimous() *UnanimousAssertionValidatorRequest {
	if x, ok := x.GetRequest().(*ValidatorRequest_Unanimous); ok {
		return x.Unanimous
	}
	return nil
}

func (x *ValidatorRequest) GetUnanimousNotification() *UnanimousAssertionValidatorNotification {
	if x, ok := x.GetRequest().(*ValidatorRequest_UnanimousNotification); ok {
		return x.UnanimousNotification
	}
	return nil
}

type isValidatorRequest_Request interface {
	isValidatorRequest_Request()
}

type ValidatorRequest_Unanimous struct {
	Unanimous *UnanimousAssertionValidatorRequest `protobuf:"bytes,2,opt,name=unanimous,proto3,oneof"`
}

type ValidatorRequest_UnanimousNotification struct {
	UnanimousNotification *UnanimousAssertionValidatorNotification `protobuf:"bytes,3,opt,name=unanimousNotification,proto3,oneof"`
}

func (*ValidatorRequest_Unanimous) isValidatorRequest_Request() {}

func (*ValidatorRequest_UnanimousNotification) isValidatorRequest_Request() {}

type UnanimousAssertionFollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted      bool            `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signature     []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	AssertionHash *common.HashBuf `protobuf:"bytes,3,opt,name=assertion_hash,json=assertionHash,proto3" json:"assertion_hash,omitempty"`
}

func (x *UnanimousAssertionFollowerResponse) Reset() {
	*x = UnanimousAssertionFollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnanimousAssertionFollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnanimousAssertionFollowerResponse) ProtoMessage() {}

func (x *UnanimousAssertionFollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnanimousAssertionFollowerResponse.ProtoReflect.Descriptor instead.
func (*UnanimousAssertionFollowerResponse) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{6}
}

func (x *UnanimousAssertionFollowerResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *UnanimousAssertionFollowerResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *UnanimousAssertionFollowerResponse) GetAssertionHash() *common.HashBuf {
	if x != nil {
		return x.AssertionHash
	}
	return nil
}

type FollowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId *common.HashBuf                     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Unanimous *UnanimousAssertionFollowerResponse `protobuf:"bytes,3,opt,name=unanimous,proto3" json:"unanimous,omitempty"`
}

func (x *FollowerResponse) Reset() {
	*x = FollowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valmessage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerResponse) ProtoMessage() {}

func (x *FollowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_valmessage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerResponse.ProtoReflect.Descriptor instead.
func (*FollowerResponse) Descriptor() ([]byte, []int) {
	return file_valmessage_proto_rawDescGZIP(), []int{7}
}

func (x *FollowerResponse) GetRequestId() *common.HashBuf {
	if x != nil {
		return x.RequestId
	}
	return nil
}

func (x *FollowerResponse) GetUnanimous() *UnanimousAssertionFollowerResponse {
	if x != nil {
		return x.Unanimous
	}
	return nil
}

var File_valmessage_proto protoreflect.FileDescriptor

var file_valmessage_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x61, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x76, 0x61, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d,
	0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x61, 0x6c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x75, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xc9, 0x02, 0x0a, 0x0f, 0x56, 0x4d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x67, 0x72, 0x61,
	0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3e, 0x0a, 0x0f, 0x65, 0x73, 0x63, 0x72,
	0x6f, 0x77, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x42, 0x75, 0x66, 0x52, 0x0e, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0f, 0x65, 0x73, 0x63, 0x72,
	0x6f, 0x77, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0e, 0x65, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0a,
	0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x61,
	0x78, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x6d, 0x61,
	0x78, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x65, 0x0a,
	0x27, 0x55, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x61, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x75, 0x66, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xac, 0x02, 0x0a, 0x22, 0x55, 0x6e, 0x61, 0x6e, 0x69,
	0x6d, 0x6f, 0x75, 0x73, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a,
	0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42,
	0x75, 0x66, 0x52, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31,
	0x0a, 0x0b, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x42, 0x75, 0x66, 0x52, 0x0b, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x62, 0x6f,
	0x78, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x61, 0x6c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x8a, 0x02, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x09, 0x75, 0x6e,
	0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x76, 0x61, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x69,
	0x6d, 0x6f, 0x75, 0x73, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x09, 0x75, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x6b, 0x0a, 0x15, 0x75, 0x6e,
	0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x76, 0x61, 0x6c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73,
	0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x15, 0x75, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x22, 0x55, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73,
	0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0d, 0x61, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x90, 0x01, 0x0a, 0x10,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x4c, 0x0a, 0x09, 0x75, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x61, 0x6c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x55, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x75, 0x6e, 0x61, 0x6e, 0x69, 0x6d, 0x6f, 0x75, 0x73, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x66, 0x66,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x69, 0x74, 0x72,
	0x75, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x72, 0x62, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x61, 0x6c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_valmessage_proto_rawDescOnce sync.Once
	file_valmessage_proto_rawDescData = file_valmessage_proto_rawDesc
)

func file_valmessage_proto_rawDescGZIP() []byte {
	file_valmessage_proto_rawDescOnce.Do(func() {
		file_valmessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_valmessage_proto_rawDescData)
	})
	return file_valmessage_proto_rawDescData
}

var file_valmessage_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_valmessage_proto_goTypes = []interface{}{
	(*TokenTypeBuf)(nil),                            // 0: valmessage.TokenTypeBuf
	(*VMConfiguration)(nil),                         // 1: valmessage.VMConfiguration
	(*UnanimousAssertionValidatorNotification)(nil), // 2: valmessage.UnanimousAssertionValidatorNotification
	(*SignedMessage)(nil),                           // 3: valmessage.SignedMessage
	(*UnanimousAssertionValidatorRequest)(nil),      // 4: valmessage.UnanimousAssertionValidatorRequest
	(*ValidatorRequest)(nil),                        // 5: valmessage.ValidatorRequest
	(*UnanimousAssertionFollowerResponse)(nil),      // 6: valmessage.UnanimousAssertionFollowerResponse
	(*FollowerResponse)(nil),                        // 7: valmessage.FollowerResponse
	(*common.BigIntegerBuf)(nil),                    // 8: common.BigIntegerBuf
	(*common.AddressBuf)(nil),                       // 9: common.AddressBuf
	(*valprotocol.MessageBuf)(nil),                  // 10: valprotocol.MessageBuf
	(*common.HashBuf)(nil),                          // 11: common.HashBuf
	(*protocol.TimeBoundsBlocksBuf)(nil),            // 12: protocol.TimeBoundsBlocksBuf
}
var file_valmessage_proto_depIdxs = []int32{
	8,  // 0: valmessage.VMConfiguration.escrow_required:type_name -> common.BigIntegerBuf
	9,  // 1: valmessage.VMConfiguration.escrow_currency:type_name -> common.AddressBuf
	9,  // 2: valmessage.VMConfiguration.assert_keys:type_name -> common.AddressBuf
	9,  // 3: valmessage.VMConfiguration.owner:type_name -> common.AddressBuf
	10, // 4: valmessage.SignedMessage.message:type_name -> valprotocol.MessageBuf
	11, // 5: valmessage.UnanimousAssertionValidatorRequest.beforeHash:type_name -> common.HashBuf
	11, // 6: valmessage.UnanimousAssertionValidatorRequest.beforeInbox:type_name -> common.HashBuf
	12, // 7: valmessage.UnanimousAssertionValidatorRequest.timeBounds:type_name -> protocol.TimeBoundsBlocksBuf
	3,  // 8: valmessage.UnanimousAssertionValidatorRequest.signedMessages:type_name -> valmessage.SignedMessage
	11, // 9: valmessage.ValidatorRequest.request_id:type_name -> common.HashBuf
	4,  // 10: valmessage.ValidatorRequest.unanimous:type_name -> valmessage.UnanimousAssertionValidatorRequest
	2,  // 11: valmessage.ValidatorRequest.unanimousNotification:type_name -> valmessage.UnanimousAssertionValidatorNotification
	11, // 12: valmessage.UnanimousAssertionFollowerResponse.assertion_hash:type_name -> common.HashBuf
	11, // 13: valmessage.FollowerResponse.request_id:type_name -> common.HashBuf
	6,  // 14: valmessage.FollowerResponse.unanimous:type_name -> valmessage.UnanimousAssertionFollowerResponse
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_valmessage_proto_init() }
func file_valmessage_proto_init() {
	if File_valmessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_valmessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenTypeBuf); i {
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
		file_valmessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VMConfiguration); i {
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
		file_valmessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnanimousAssertionValidatorNotification); i {
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
		file_valmessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedMessage); i {
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
		file_valmessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnanimousAssertionValidatorRequest); i {
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
		file_valmessage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorRequest); i {
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
		file_valmessage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnanimousAssertionFollowerResponse); i {
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
		file_valmessage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerResponse); i {
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
	file_valmessage_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ValidatorRequest_Unanimous)(nil),
		(*ValidatorRequest_UnanimousNotification)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_valmessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_valmessage_proto_goTypes,
		DependencyIndexes: file_valmessage_proto_depIdxs,
		MessageInfos:      file_valmessage_proto_msgTypes,
	}.Build()
	File_valmessage_proto = out.File
	file_valmessage_proto_rawDesc = nil
	file_valmessage_proto_goTypes = nil
	file_valmessage_proto_depIdxs = nil
}
