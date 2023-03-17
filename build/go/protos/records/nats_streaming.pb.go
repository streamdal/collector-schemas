// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: nats_streaming.proto

package records

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

type NATSStreamingRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel         string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Value           []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Sequence        uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Subject         string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Redelivered     bool   `protobuf:"varint,5,opt,name=redelivered,proto3" json:"redelivered,omitempty"`
	RedeliveryCount uint32 `protobuf:"varint,6,opt,name=redelivery_count,json=redeliveryCount,proto3" json:"redelivery_count,omitempty"`
	Crc32           uint32 `protobuf:"varint,7,opt,name=crc32,proto3" json:"crc32,omitempty"`
	Timestamp       int64  `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // unix nano
	ForceDeadLetter bool   `protobuf:"varint,9,opt,name=force_dead_letter,json=forceDeadLetter,proto3" json:"force_dead_letter,omitempty"`
}

func (x *NATSStreamingRecord) Reset() {
	*x = NATSStreamingRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nats_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NATSStreamingRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NATSStreamingRecord) ProtoMessage() {}

func (x *NATSStreamingRecord) ProtoReflect() protoreflect.Message {
	mi := &file_nats_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NATSStreamingRecord.ProtoReflect.Descriptor instead.
func (*NATSStreamingRecord) Descriptor() ([]byte, []int) {
	return file_nats_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *NATSStreamingRecord) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *NATSStreamingRecord) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *NATSStreamingRecord) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *NATSStreamingRecord) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *NATSStreamingRecord) GetRedelivered() bool {
	if x != nil {
		return x.Redelivered
	}
	return false
}

func (x *NATSStreamingRecord) GetRedeliveryCount() uint32 {
	if x != nil {
		return x.RedeliveryCount
	}
	return 0
}

func (x *NATSStreamingRecord) GetCrc32() uint32 {
	if x != nil {
		return x.Crc32
	}
	return 0
}

func (x *NATSStreamingRecord) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *NATSStreamingRecord) GetForceDeadLetter() bool {
	if x != nil {
		return x.ForceDeadLetter
	}
	return false
}

var File_nats_streaming_proto protoreflect.FileDescriptor

var file_nats_streaming_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6e, 0x61, 0x74, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0xa8, 0x02, 0x0a, 0x13, 0x4e, 0x41, 0x54, 0x53, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x72, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x72,
	0x63, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x72, 0x63, 0x33, 0x32,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a,
	0x0a, 0x11, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x42, 0x59, 0x0a, 0x17, 0x73, 0x68,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nats_streaming_proto_rawDescOnce sync.Once
	file_nats_streaming_proto_rawDescData = file_nats_streaming_proto_rawDesc
)

func file_nats_streaming_proto_rawDescGZIP() []byte {
	file_nats_streaming_proto_rawDescOnce.Do(func() {
		file_nats_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_nats_streaming_proto_rawDescData)
	})
	return file_nats_streaming_proto_rawDescData
}

var file_nats_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nats_streaming_proto_goTypes = []interface{}{
	(*NATSStreamingRecord)(nil), // 0: records.NATSStreamingRecord
}
var file_nats_streaming_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nats_streaming_proto_init() }
func file_nats_streaming_proto_init() {
	if File_nats_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nats_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NATSStreamingRecord); i {
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
			RawDescriptor: file_nats_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nats_streaming_proto_goTypes,
		DependencyIndexes: file_nats_streaming_proto_depIdxs,
		MessageInfos:      file_nats_streaming_proto_msgTypes,
	}.Build()
	File_nats_streaming_proto = out.File
	file_nats_streaming_proto_rawDesc = nil
	file_nats_streaming_proto_goTypes = nil
	file_nats_streaming_proto_depIdxs = nil
}
