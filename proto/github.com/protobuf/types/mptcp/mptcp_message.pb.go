// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: mptcp_message.proto

package mptcp

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

type MPTCPMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcAddr           string   `protobuf:"bytes,1,opt,name=srcAddr,proto3" json:"srcAddr,omitempty"`
	DstAddr           string   `protobuf:"bytes,2,opt,name=dstAddr,proto3" json:"dstAddr,omitempty"`
	SrcPort           uint32   `protobuf:"varint,3,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	DstPort           uint32   `protobuf:"varint,4,opt,name=dstPort,proto3" json:"dstPort,omitempty"`
	SeqNum            uint32   `protobuf:"varint,5,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
	TimestampCaptured int64    `protobuf:"varint,6,opt,name=timestampCaptured,proto3" json:"timestampCaptured,omitempty"`
	MptcpOptions      []string `protobuf:"bytes,7,rep,name=mptcpOptions,proto3" json:"mptcpOptions,omitempty"`
}

func (x *MPTCPMessage) Reset() {
	*x = MPTCPMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mptcp_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MPTCPMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MPTCPMessage) ProtoMessage() {}

func (x *MPTCPMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mptcp_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MPTCPMessage.ProtoReflect.Descriptor instead.
func (*MPTCPMessage) Descriptor() ([]byte, []int) {
	return file_mptcp_message_proto_rawDescGZIP(), []int{0}
}

func (x *MPTCPMessage) GetSrcAddr() string {
	if x != nil {
		return x.SrcAddr
	}
	return ""
}

func (x *MPTCPMessage) GetDstAddr() string {
	if x != nil {
		return x.DstAddr
	}
	return ""
}

func (x *MPTCPMessage) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *MPTCPMessage) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *MPTCPMessage) GetSeqNum() uint32 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (x *MPTCPMessage) GetTimestampCaptured() int64 {
	if x != nil {
		return x.TimestampCaptured
	}
	return 0
}

func (x *MPTCPMessage) GetMptcpOptions() []string {
	if x != nil {
		return x.MptcpOptions
	}
	return nil
}

var File_mptcp_message_proto protoreflect.FileDescriptor

var file_mptcp_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x70, 0x74, 0x63, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x0c,
	0x4d, 0x50, 0x54, 0x43, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x11,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x70,
	0x74, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x70, 0x74, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x70, 0x74, 0x63,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mptcp_message_proto_rawDescOnce sync.Once
	file_mptcp_message_proto_rawDescData = file_mptcp_message_proto_rawDesc
)

func file_mptcp_message_proto_rawDescGZIP() []byte {
	file_mptcp_message_proto_rawDescOnce.Do(func() {
		file_mptcp_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_mptcp_message_proto_rawDescData)
	})
	return file_mptcp_message_proto_rawDescData
}

var file_mptcp_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mptcp_message_proto_goTypes = []interface{}{
	(*MPTCPMessage)(nil), // 0: main.MPTCPMessage
}
var file_mptcp_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mptcp_message_proto_init() }
func file_mptcp_message_proto_init() {
	if File_mptcp_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mptcp_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MPTCPMessage); i {
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
			RawDescriptor: file_mptcp_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mptcp_message_proto_goTypes,
		DependencyIndexes: file_mptcp_message_proto_depIdxs,
		MessageInfos:      file_mptcp_message_proto_msgTypes,
	}.Build()
	File_mptcp_message_proto = out.File
	file_mptcp_message_proto_rawDesc = nil
	file_mptcp_message_proto_goTypes = nil
	file_mptcp_message_proto_depIdxs = nil
}
