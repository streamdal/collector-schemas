// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/azure.proto

package records

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// See: https://github.com/Azure/azure-service-bus-go/blob/78c960db000c65f4e32d01371b496d5b33f59c38/message.go#L40
type AzureSinkRecord struct {
	ContentType          string                 `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CorrelationId        string                 `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Data                 []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	DeliveryCount        uint32                 `protobuf:"varint,4,opt,name=delivery_count,json=deliveryCount,proto3" json:"delivery_count,omitempty"`
	SessionId            string                 `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	GroupSequence        uint32                 `protobuf:"varint,6,opt,name=group_sequence,json=groupSequence,proto3" json:"group_sequence,omitempty"`
	Id                   string                 `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	Label                string                 `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty"`
	ReplyTo              string                 `protobuf:"bytes,9,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	ReplyToGroupId       string                 `protobuf:"bytes,10,opt,name=reply_to_group_id,json=replyToGroupId,proto3" json:"reply_to_group_id,omitempty"`
	To                   string                 `protobuf:"bytes,11,opt,name=to,proto3" json:"to,omitempty"`
	Ttl                  int64                  `protobuf:"varint,12,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LockToken            string                 `protobuf:"bytes,13,opt,name=lock_token,json=lockToken,proto3" json:"lock_token,omitempty"`
	SystemProperties     *AzureSystemProperties `protobuf:"bytes,14,opt,name=system_properties,json=systemProperties,proto3" json:"system_properties,omitempty"`
	UserProperties       map[string]string      `protobuf:"bytes,15,rep,name=user_properties,json=userProperties,proto3" json:"user_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Format               uint32                 `protobuf:"varint,16,opt,name=format,proto3" json:"format,omitempty"`
	ForceDeadLetter      bool                   `protobuf:"varint,17,opt,name=force_dead_letter,json=forceDeadLetter,proto3" json:"force_dead_letter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AzureSinkRecord) Reset()         { *m = AzureSinkRecord{} }
func (m *AzureSinkRecord) String() string { return proto.CompactTextString(m) }
func (*AzureSinkRecord) ProtoMessage()    {}
func (*AzureSinkRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eedd0d64dd1a823, []int{0}
}

func (m *AzureSinkRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSinkRecord.Unmarshal(m, b)
}
func (m *AzureSinkRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSinkRecord.Marshal(b, m, deterministic)
}
func (m *AzureSinkRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSinkRecord.Merge(m, src)
}
func (m *AzureSinkRecord) XXX_Size() int {
	return xxx_messageInfo_AzureSinkRecord.Size(m)
}
func (m *AzureSinkRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSinkRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSinkRecord proto.InternalMessageInfo

func (m *AzureSinkRecord) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *AzureSinkRecord) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *AzureSinkRecord) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AzureSinkRecord) GetDeliveryCount() uint32 {
	if m != nil {
		return m.DeliveryCount
	}
	return 0
}

func (m *AzureSinkRecord) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *AzureSinkRecord) GetGroupSequence() uint32 {
	if m != nil {
		return m.GroupSequence
	}
	return 0
}

func (m *AzureSinkRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AzureSinkRecord) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AzureSinkRecord) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *AzureSinkRecord) GetReplyToGroupId() string {
	if m != nil {
		return m.ReplyToGroupId
	}
	return ""
}

func (m *AzureSinkRecord) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *AzureSinkRecord) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *AzureSinkRecord) GetLockToken() string {
	if m != nil {
		return m.LockToken
	}
	return ""
}

func (m *AzureSinkRecord) GetSystemProperties() *AzureSystemProperties {
	if m != nil {
		return m.SystemProperties
	}
	return nil
}

func (m *AzureSinkRecord) GetUserProperties() map[string]string {
	if m != nil {
		return m.UserProperties
	}
	return nil
}

func (m *AzureSinkRecord) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

func (m *AzureSinkRecord) GetForceDeadLetter() bool {
	if m != nil {
		return m.ForceDeadLetter
	}
	return false
}

type AzureSystemProperties struct {
	LockedUntil            int64             `protobuf:"varint,1,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	SequenceNumber         int64             `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	PartitionId            int32             `protobuf:"varint,3,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	PartitionKey           string            `protobuf:"bytes,4,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	EnqueuedTime           int64             `protobuf:"varint,5,opt,name=enqueued_time,json=enqueuedTime,proto3" json:"enqueued_time,omitempty"`
	DeadLetterSource       string            `protobuf:"bytes,6,opt,name=dead_letter_source,json=deadLetterSource,proto3" json:"dead_letter_source,omitempty"`
	ScheduledEnqueueTime   int64             `protobuf:"varint,7,opt,name=scheduled_enqueue_time,json=scheduledEnqueueTime,proto3" json:"scheduled_enqueue_time,omitempty"`
	EnqueuedSequenceNumber int64             `protobuf:"varint,8,opt,name=enqueued_sequence_number,json=enqueuedSequenceNumber,proto3" json:"enqueued_sequence_number,omitempty"`
	ViaPartitionKey        string            `protobuf:"bytes,9,opt,name=via_partition_key,json=viaPartitionKey,proto3" json:"via_partition_key,omitempty"`
	Annotations            map[string]string `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *AzureSystemProperties) Reset()         { *m = AzureSystemProperties{} }
func (m *AzureSystemProperties) String() string { return proto.CompactTextString(m) }
func (*AzureSystemProperties) ProtoMessage()    {}
func (*AzureSystemProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eedd0d64dd1a823, []int{1}
}

func (m *AzureSystemProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSystemProperties.Unmarshal(m, b)
}
func (m *AzureSystemProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSystemProperties.Marshal(b, m, deterministic)
}
func (m *AzureSystemProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSystemProperties.Merge(m, src)
}
func (m *AzureSystemProperties) XXX_Size() int {
	return xxx_messageInfo_AzureSystemProperties.Size(m)
}
func (m *AzureSystemProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSystemProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSystemProperties proto.InternalMessageInfo

func (m *AzureSystemProperties) GetLockedUntil() int64 {
	if m != nil {
		return m.LockedUntil
	}
	return 0
}

func (m *AzureSystemProperties) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *AzureSystemProperties) GetPartitionId() int32 {
	if m != nil {
		return m.PartitionId
	}
	return 0
}

func (m *AzureSystemProperties) GetPartitionKey() string {
	if m != nil {
		return m.PartitionKey
	}
	return ""
}

func (m *AzureSystemProperties) GetEnqueuedTime() int64 {
	if m != nil {
		return m.EnqueuedTime
	}
	return 0
}

func (m *AzureSystemProperties) GetDeadLetterSource() string {
	if m != nil {
		return m.DeadLetterSource
	}
	return ""
}

func (m *AzureSystemProperties) GetScheduledEnqueueTime() int64 {
	if m != nil {
		return m.ScheduledEnqueueTime
	}
	return 0
}

func (m *AzureSystemProperties) GetEnqueuedSequenceNumber() int64 {
	if m != nil {
		return m.EnqueuedSequenceNumber
	}
	return 0
}

func (m *AzureSystemProperties) GetViaPartitionKey() string {
	if m != nil {
		return m.ViaPartitionKey
	}
	return ""
}

func (m *AzureSystemProperties) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*AzureSinkRecord)(nil), "records.AzureSinkRecord")
	proto.RegisterMapType((map[string]string)(nil), "records.AzureSinkRecord.UserPropertiesEntry")
	proto.RegisterType((*AzureSystemProperties)(nil), "records.AzureSystemProperties")
	proto.RegisterMapType((map[string]string)(nil), "records.AzureSystemProperties.AnnotationsEntry")
}

func init() { proto.RegisterFile("records/azure.proto", fileDescriptor_8eedd0d64dd1a823) }

var fileDescriptor_8eedd0d64dd1a823 = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x86, 0x21, 0x7f, 0xc4, 0x36, 0xfd, 0xcd, 0x64, 0x19, 0x17, 0x60, 0x83, 0x97, 0x61, 0x98,
	0x17, 0x64, 0x12, 0x90, 0xed, 0x22, 0xd8, 0x45, 0x80, 0xb4, 0x0d, 0x0a, 0x23, 0x45, 0x91, 0xca,
	0xce, 0x45, 0x7b, 0x43, 0xc8, 0xe2, 0x49, 0x4c, 0x58, 0x16, 0x15, 0x92, 0x32, 0xa0, 0xfe, 0xa7,
	0xde, 0xf5, 0x07, 0x16, 0xa4, 0x24, 0xdb, 0x31, 0x82, 0x02, 0xbd, 0x23, 0x9f, 0x73, 0xce, 0x7b,
	0xc8, 0xc3, 0x57, 0x42, 0x87, 0x12, 0x42, 0x21, 0x99, 0xf2, 0x82, 0xcf, 0xa9, 0x04, 0x37, 0x91,
	0x42, 0x0b, 0xdc, 0x28, 0xe0, 0xe9, 0xd7, 0x3a, 0xea, 0x5f, 0x9b, 0xc0, 0x94, 0xc7, 0x4b, 0xdf,
	0x42, 0xfc, 0x3b, 0xea, 0x84, 0x22, 0xd6, 0x10, 0x6b, 0xaa, 0xb3, 0x04, 0x88, 0x33, 0x72, 0xc6,
	0x2d, 0xbf, 0x5d, 0xb0, 0x59, 0x96, 0x00, 0xfe, 0x13, 0xf5, 0x42, 0x21, 0x25, 0x44, 0x81, 0xe6,
	0x22, 0xa6, 0x9c, 0x91, 0x8a, 0x4d, 0xea, 0xee, 0xd0, 0x09, 0xc3, 0x18, 0xd5, 0x58, 0xa0, 0x03,
	0x52, 0x1d, 0x39, 0xe3, 0x8e, 0x6f, 0xd7, 0xa6, 0x94, 0x41, 0xc4, 0xd7, 0x20, 0x33, 0x1a, 0x8a,
	0x34, 0xd6, 0xa4, 0x36, 0x72, 0xc6, 0x5d, 0xbf, 0x5b, 0xd2, 0xd7, 0x06, 0xe2, 0x5f, 0x11, 0x52,
	0xa0, 0x54, 0xa1, 0x5e, 0xb7, 0xea, 0xad, 0x82, 0x4c, 0x98, 0x51, 0x79, 0x94, 0x22, 0x4d, 0xa8,
	0x82, 0xa7, 0x14, 0xe2, 0x10, 0xc8, 0x41, 0xae, 0x62, 0xe9, 0xb4, 0x80, 0xb8, 0x87, 0x2a, 0x9c,
	0x91, 0x86, 0xad, 0xae, 0x70, 0x86, 0x8f, 0x50, 0x3d, 0x0a, 0xe6, 0x10, 0x91, 0xa6, 0x45, 0xf9,
	0x06, 0xff, 0x82, 0x9a, 0x12, 0x92, 0x28, 0xa3, 0x5a, 0x90, 0x96, 0x0d, 0x34, 0xec, 0x7e, 0x26,
	0xf0, 0xdf, 0x68, 0x58, 0x86, 0x68, 0xde, 0x90, 0x33, 0x82, 0x6c, 0x4e, 0xaf, 0xc8, 0x79, 0x6b,
	0xf0, 0x84, 0x99, 0x5e, 0x5a, 0x90, 0x76, 0xde, 0x4b, 0x0b, 0x3c, 0x40, 0x55, 0xad, 0x23, 0xd2,
	0x19, 0x39, 0xe3, 0xaa, 0x6f, 0x96, 0xe6, 0x4e, 0x91, 0x08, 0x97, 0x54, 0x8b, 0x25, 0xc4, 0xa4,
	0x9b, 0xdf, 0xc9, 0x90, 0x99, 0x01, 0xf8, 0x16, 0x0d, 0x55, 0xa6, 0x34, 0xac, 0x68, 0x22, 0x45,
	0x02, 0x52, 0x73, 0x50, 0xa4, 0x37, 0x72, 0xc6, 0xed, 0x8b, 0xdf, 0xdc, 0xe2, 0xc1, 0xdc, 0xfc,
	0xb1, 0x6c, 0xda, 0xdd, 0x26, 0xcb, 0x1f, 0xa8, 0x3d, 0x82, 0xef, 0x51, 0x3f, 0x55, 0x20, 0x77,
	0xa5, 0xfa, 0xa3, 0xea, 0xb8, 0x7d, 0x71, 0xbe, 0x27, 0xb5, 0x79, 0x77, 0xf7, 0x5e, 0x81, 0xdc,
	0x2a, 0xdc, 0xc4, 0x5a, 0x66, 0x7e, 0x2f, 0x7d, 0x06, 0xf1, 0x31, 0x3a, 0x78, 0x10, 0x72, 0x15,
	0x68, 0x32, 0xb0, 0xf3, 0x2e, 0x76, 0xf8, 0x0c, 0x0d, 0x1f, 0x84, 0x0c, 0x81, 0x32, 0x08, 0x18,
	0x8d, 0x40, 0x6b, 0x90, 0x64, 0x38, 0x72, 0xc6, 0x4d, 0xbf, 0x6f, 0x03, 0x6f, 0x20, 0x60, 0xef,
	0x2c, 0x3e, 0xb9, 0x46, 0x87, 0x2f, 0xb4, 0x32, 0xf3, 0x5a, 0x42, 0x56, 0xb8, 0xcd, 0x2c, 0xcd,
	0x6b, 0xad, 0x83, 0x28, 0x85, 0xc2, 0x5c, 0xf9, 0xe6, 0xff, 0xca, 0xa5, 0x73, 0xfa, 0xa5, 0x86,
	0x7e, 0x7a, 0x71, 0x12, 0xc6, 0xbc, 0x66, 0xa2, 0xc0, 0x68, 0x1a, 0x6b, 0x1e, 0x59, 0xb9, 0xaa,
	0xdf, 0xce, 0xd9, 0xbd, 0x41, 0xf8, 0x2f, 0xd4, 0x2f, 0x5d, 0x43, 0xe3, 0x74, 0x35, 0x07, 0x69,
	0x1b, 0x54, 0xfd, 0x5e, 0x89, 0xdf, 0x5b, 0x6a, 0xb4, 0x92, 0x40, 0x6a, 0x5e, 0x7a, 0xdc, 0xd8,
	0xb8, 0xee, 0xb7, 0x37, 0x6c, 0xc2, 0xf0, 0x1f, 0xa8, 0xbb, 0x4d, 0x31, 0xc7, 0xaf, 0xd9, 0xa3,
	0x6e, 0xeb, 0x6e, 0x21, 0x33, 0x49, 0x10, 0x3f, 0xa5, 0x90, 0x02, 0xa3, 0x9a, 0xaf, 0xc0, 0xda,
	0xb9, 0xea, 0x77, 0x4a, 0x38, 0xe3, 0x2b, 0xc0, 0xe7, 0x08, 0xef, 0xcc, 0x8e, 0x2a, 0x91, 0xca,
	0xc2, 0xd5, 0x2d, 0x7f, 0xc0, 0x36, 0xd3, 0x9b, 0x5a, 0x8e, 0xff, 0x43, 0xc7, 0x2a, 0x5c, 0x00,
	0x4b, 0x23, 0x60, 0xb4, 0xd0, 0xc9, 0xb5, 0x1b, 0x56, 0xfb, 0x68, 0x13, 0xbd, 0xc9, 0x83, 0xb6,
	0xc7, 0x25, 0x22, 0x9b, 0x83, 0xec, 0x8f, 0xa0, 0x69, 0xeb, 0x8e, 0xcb, 0xf8, 0xf4, 0xf9, 0x28,
	0xce, 0xd0, 0x70, 0xcd, 0x03, 0xfa, 0xfc, 0xae, 0xf9, 0xb7, 0xd2, 0x5f, 0xf3, 0xe0, 0x6e, 0xf7,
	0xba, 0x1f, 0x50, 0x3b, 0x88, 0x63, 0xa1, 0xed, 0x5f, 0x40, 0x11, 0x64, 0x6d, 0xe7, 0x7d, 0xdf,
	0xc1, 0xee, 0xf5, 0xb6, 0x22, 0x77, 0xde, 0xae, 0xc6, 0xc9, 0x15, 0x1a, 0xec, 0x27, 0xfc, 0x88,
	0x5f, 0x5e, 0x7d, 0x44, 0x3f, 0xab, 0x85, 0x3b, 0x0f, 0x74, 0xb8, 0x70, 0x61, 0x0d, 0xb1, 0x56,
	0xe5, 0x71, 0x3e, 0x5d, 0x3d, 0x72, 0xbd, 0x48, 0xe7, 0x6e, 0x28, 0x56, 0x9e, 0x4d, 0x08, 0x85,
	0x4c, 0xbc, 0x50, 0x44, 0x11, 0x84, 0x5a, 0xc8, 0x7f, 0xcc, 0x20, 0x57, 0x81, 0xf2, 0xe6, 0x29,
	0x8f, 0x98, 0xf7, 0x28, 0x3c, 0xfb, 0x07, 0x55, 0x5e, 0x51, 0x3f, 0x3f, 0xb0, 0xfb, 0x7f, 0xbf,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x10, 0x6b, 0x36, 0xb8, 0x68, 0x05, 0x00, 0x00,
}
