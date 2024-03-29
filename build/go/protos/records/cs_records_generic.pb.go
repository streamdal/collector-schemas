// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: cs_records_generic.proto

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

type GenericRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body            []byte            `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Source          string            `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Timestamp       int64             `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // client unix nano
	Metadata        map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ForceDeadLetter bool              `protobuf:"varint,5,opt,name=force_dead_letter,json=forceDeadLetter,proto3" json:"force_dead_letter,omitempty"`
}

func (x *GenericRecord) Reset() {
	*x = GenericRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs_records_generic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericRecord) ProtoMessage() {}

func (x *GenericRecord) ProtoReflect() protoreflect.Message {
	mi := &file_cs_records_generic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericRecord.ProtoReflect.Descriptor instead.
func (*GenericRecord) Descriptor() ([]byte, []int) {
	return file_cs_records_generic_proto_rawDescGZIP(), []int{0}
}

func (x *GenericRecord) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *GenericRecord) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GenericRecord) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GenericRecord) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GenericRecord) GetForceDeadLetter() bool {
	if x != nil {
		return x.ForceDeadLetter
	}
	return false
}

var File_cs_records_generic_proto protoreflect.FileDescriptor

var file_cs_records_generic_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f,
	0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x3b, 0x0a,
	0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x59, 0x0a, 0x17, 0x73, 0x68,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cs_records_generic_proto_rawDescOnce sync.Once
	file_cs_records_generic_proto_rawDescData = file_cs_records_generic_proto_rawDesc
)

func file_cs_records_generic_proto_rawDescGZIP() []byte {
	file_cs_records_generic_proto_rawDescOnce.Do(func() {
		file_cs_records_generic_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs_records_generic_proto_rawDescData)
	})
	return file_cs_records_generic_proto_rawDescData
}

var file_cs_records_generic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cs_records_generic_proto_goTypes = []interface{}{
	(*GenericRecord)(nil), // 0: records.GenericRecord
	nil,                   // 1: records.GenericRecord.MetadataEntry
}
var file_cs_records_generic_proto_depIdxs = []int32{
	1, // 0: records.GenericRecord.metadata:type_name -> records.GenericRecord.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cs_records_generic_proto_init() }
func file_cs_records_generic_proto_init() {
	if File_cs_records_generic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cs_records_generic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericRecord); i {
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
			RawDescriptor: file_cs_records_generic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs_records_generic_proto_goTypes,
		DependencyIndexes: file_cs_records_generic_proto_depIdxs,
		MessageInfos:      file_cs_records_generic_proto_msgTypes,
	}.Build()
	File_cs_records_generic_proto = out.File
	file_cs_records_generic_proto_rawDesc = nil
	file_cs_records_generic_proto_goTypes = nil
	file_cs_records_generic_proto_depIdxs = nil
}
