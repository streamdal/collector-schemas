// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: outbound.proto

package events

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

// Emitted by the reader to HSB which is then consumed by the replayer and dProxy
type Outbound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayId            string            `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	Blob                []byte            `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
	Last                bool              `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	Metadata            map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IngestTimestampNano int64             `protobuf:"varint,5,opt,name=ingest_timestamp_nano,json=ingestTimestampNano,proto3" json:"ingest_timestamp_nano,omitempty"`
}

func (x *Outbound) Reset() {
	*x = Outbound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outbound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outbound) ProtoMessage() {}

func (x *Outbound) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outbound.ProtoReflect.Descriptor instead.
func (*Outbound) Descriptor() ([]byte, []int) {
	return file_outbound_proto_rawDescGZIP(), []int{0}
}

func (x *Outbound) GetReplayId() string {
	if x != nil {
		return x.ReplayId
	}
	return ""
}

func (x *Outbound) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *Outbound) GetLast() bool {
	if x != nil {
		return x.Last
	}
	return false
}

func (x *Outbound) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Outbound) GetIngestTimestampNano() int64 {
	if x != nil {
		return x.IngestTimestampNano
	}
	return 0
}

var File_outbound_proto protoreflect.FileDescriptor

var file_outbound_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x08, 0x4f, 0x75, 0x74,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x61, 0x6e, 0x6f, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outbound_proto_rawDescOnce sync.Once
	file_outbound_proto_rawDescData = file_outbound_proto_rawDesc
)

func file_outbound_proto_rawDescGZIP() []byte {
	file_outbound_proto_rawDescOnce.Do(func() {
		file_outbound_proto_rawDescData = protoimpl.X.CompressGZIP(file_outbound_proto_rawDescData)
	})
	return file_outbound_proto_rawDescData
}

var file_outbound_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_outbound_proto_goTypes = []interface{}{
	(*Outbound)(nil), // 0: events.Outbound
	nil,              // 1: events.Outbound.MetadataEntry
}
var file_outbound_proto_depIdxs = []int32{
	1, // 0: events.Outbound.metadata:type_name -> events.Outbound.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_outbound_proto_init() }
func file_outbound_proto_init() {
	if File_outbound_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outbound_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outbound); i {
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
			RawDescriptor: file_outbound_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_outbound_proto_goTypes,
		DependencyIndexes: file_outbound_proto_depIdxs,
		MessageInfos:      file_outbound_proto_msgTypes,
	}.Build()
	File_outbound_proto = out.File
	file_outbound_proto_rawDesc = nil
	file_outbound_proto_goTypes = nil
	file_outbound_proto_depIdxs = nil
}
