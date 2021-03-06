// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.1
// source: common.proto

package common

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

type BigIntegerBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BigIntegerBuf) Reset() {
	*x = BigIntegerBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigIntegerBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigIntegerBuf) ProtoMessage() {}

func (x *BigIntegerBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigIntegerBuf.ProtoReflect.Descriptor instead.
func (*BigIntegerBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *BigIntegerBuf) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type HashBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HashBuf) Reset() {
	*x = HashBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashBuf) ProtoMessage() {}

func (x *HashBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashBuf.ProtoReflect.Descriptor instead.
func (*HashBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *HashBuf) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type TimeBlocksBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *BigIntegerBuf `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *TimeBlocksBuf) Reset() {
	*x = TimeBlocksBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeBlocksBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeBlocksBuf) ProtoMessage() {}

func (x *TimeBlocksBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeBlocksBuf.ProtoReflect.Descriptor instead.
func (*TimeBlocksBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *TimeBlocksBuf) GetVal() *BigIntegerBuf {
	if x != nil {
		return x.Val
	}
	return nil
}

type TimeTicksBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *BigIntegerBuf `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *TimeTicksBuf) Reset() {
	*x = TimeTicksBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeTicksBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeTicksBuf) ProtoMessage() {}

func (x *TimeTicksBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeTicksBuf.ProtoReflect.Descriptor instead.
func (*TimeTicksBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *TimeTicksBuf) GetVal() *BigIntegerBuf {
	if x != nil {
		return x.Val
	}
	return nil
}

type AddressBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddressBuf) Reset() {
	*x = AddressBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBuf) ProtoMessage() {}

func (x *AddressBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBuf.ProtoReflect.Descriptor instead.
func (*AddressBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *AddressBuf) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type BlockIdBuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height     *TimeBlocksBuf `protobuf:"bytes,1,opt,name=height,proto3" json:"height,omitempty"`
	HeaderHash *HashBuf       `protobuf:"bytes,2,opt,name=headerHash,proto3" json:"headerHash,omitempty"`
}

func (x *BlockIdBuf) Reset() {
	*x = BlockIdBuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockIdBuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockIdBuf) ProtoMessage() {}

func (x *BlockIdBuf) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockIdBuf.ProtoReflect.Descriptor instead.
func (*BlockIdBuf) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *BlockIdBuf) GetHeight() *TimeBlocksBuf {
	if x != nil {
		return x.Height
	}
	return nil
}

func (x *BlockIdBuf) GetHeaderHash() *HashBuf {
	if x != nil {
		return x.HeaderHash
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0d, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x42, 0x75, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a,
	0x07, 0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38,
	0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x12,
	0x27, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x37, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x69, 0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x12, 0x27, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x69, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x42, 0x75, 0x66, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x75, 0x66, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6c, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x42, 0x75, 0x66, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x75, 0x66, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x42, 0x75, 0x66, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48,
	0x61, 0x73, 0x68, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61,
	0x72, 0x62, 0x69, 0x74, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x61, 0x72, 0x62, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(*BigIntegerBuf)(nil), // 0: common.BigIntegerBuf
	(*HashBuf)(nil),       // 1: common.HashBuf
	(*TimeBlocksBuf)(nil), // 2: common.TimeBlocksBuf
	(*TimeTicksBuf)(nil),  // 3: common.TimeTicksBuf
	(*AddressBuf)(nil),    // 4: common.AddressBuf
	(*BlockIdBuf)(nil),    // 5: common.BlockIdBuf
}
var file_common_proto_depIdxs = []int32{
	0, // 0: common.TimeBlocksBuf.val:type_name -> common.BigIntegerBuf
	0, // 1: common.TimeTicksBuf.val:type_name -> common.BigIntegerBuf
	2, // 2: common.BlockIdBuf.height:type_name -> common.TimeBlocksBuf
	1, // 3: common.BlockIdBuf.headerHash:type_name -> common.HashBuf
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigIntegerBuf); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashBuf); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeBlocksBuf); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeTicksBuf); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBuf); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockIdBuf); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
