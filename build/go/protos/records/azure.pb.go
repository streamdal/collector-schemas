// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: azure.proto

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

// See: https://github.com/Azure/azure-service-bus-go/blob/78c960db000c65f4e32d01371b496d5b33f59c38/message.go#L40
type AzureSinkRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType      string                 `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CorrelationId    string                 `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Data             []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	DeliveryCount    uint32                 `protobuf:"varint,4,opt,name=delivery_count,json=deliveryCount,proto3" json:"delivery_count,omitempty"`
	SessionId        string                 `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	GroupSequence    uint32                 `protobuf:"varint,6,opt,name=group_sequence,json=groupSequence,proto3" json:"group_sequence,omitempty"`
	Id               string                 `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	Label            string                 `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty"`
	ReplyTo          string                 `protobuf:"bytes,9,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	ReplyToGroupId   string                 `protobuf:"bytes,10,opt,name=reply_to_group_id,json=replyToGroupId,proto3" json:"reply_to_group_id,omitempty"`
	To               string                 `protobuf:"bytes,11,opt,name=to,proto3" json:"to,omitempty"`
	Ttl              int64                  `protobuf:"varint,12,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LockToken        string                 `protobuf:"bytes,13,opt,name=lock_token,json=lockToken,proto3" json:"lock_token,omitempty"`
	SystemProperties *AzureSystemProperties `protobuf:"bytes,14,opt,name=system_properties,json=systemProperties,proto3" json:"system_properties,omitempty"`
	UserProperties   map[string]string      `protobuf:"bytes,15,rep,name=user_properties,json=userProperties,proto3" json:"user_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Format           uint32                 `protobuf:"varint,16,opt,name=format,proto3" json:"format,omitempty"`
	ForceDeadLetter  bool                   `protobuf:"varint,17,opt,name=force_dead_letter,json=forceDeadLetter,proto3" json:"force_dead_letter,omitempty"`
}

func (x *AzureSinkRecord) Reset() {
	*x = AzureSinkRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_azure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureSinkRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureSinkRecord) ProtoMessage() {}

func (x *AzureSinkRecord) ProtoReflect() protoreflect.Message {
	mi := &file_azure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureSinkRecord.ProtoReflect.Descriptor instead.
func (*AzureSinkRecord) Descriptor() ([]byte, []int) {
	return file_azure_proto_rawDescGZIP(), []int{0}
}

func (x *AzureSinkRecord) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *AzureSinkRecord) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *AzureSinkRecord) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AzureSinkRecord) GetDeliveryCount() uint32 {
	if x != nil {
		return x.DeliveryCount
	}
	return 0
}

func (x *AzureSinkRecord) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *AzureSinkRecord) GetGroupSequence() uint32 {
	if x != nil {
		return x.GroupSequence
	}
	return 0
}

func (x *AzureSinkRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AzureSinkRecord) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AzureSinkRecord) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *AzureSinkRecord) GetReplyToGroupId() string {
	if x != nil {
		return x.ReplyToGroupId
	}
	return ""
}

func (x *AzureSinkRecord) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *AzureSinkRecord) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *AzureSinkRecord) GetLockToken() string {
	if x != nil {
		return x.LockToken
	}
	return ""
}

func (x *AzureSinkRecord) GetSystemProperties() *AzureSystemProperties {
	if x != nil {
		return x.SystemProperties
	}
	return nil
}

func (x *AzureSinkRecord) GetUserProperties() map[string]string {
	if x != nil {
		return x.UserProperties
	}
	return nil
}

func (x *AzureSinkRecord) GetFormat() uint32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *AzureSinkRecord) GetForceDeadLetter() bool {
	if x != nil {
		return x.ForceDeadLetter
	}
	return false
}

type AzureSystemProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
}

func (x *AzureSystemProperties) Reset() {
	*x = AzureSystemProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_azure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzureSystemProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureSystemProperties) ProtoMessage() {}

func (x *AzureSystemProperties) ProtoReflect() protoreflect.Message {
	mi := &file_azure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureSystemProperties.ProtoReflect.Descriptor instead.
func (*AzureSystemProperties) Descriptor() ([]byte, []int) {
	return file_azure_proto_rawDescGZIP(), []int{1}
}

func (x *AzureSystemProperties) GetLockedUntil() int64 {
	if x != nil {
		return x.LockedUntil
	}
	return 0
}

func (x *AzureSystemProperties) GetSequenceNumber() int64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *AzureSystemProperties) GetPartitionId() int32 {
	if x != nil {
		return x.PartitionId
	}
	return 0
}

func (x *AzureSystemProperties) GetPartitionKey() string {
	if x != nil {
		return x.PartitionKey
	}
	return ""
}

func (x *AzureSystemProperties) GetEnqueuedTime() int64 {
	if x != nil {
		return x.EnqueuedTime
	}
	return 0
}

func (x *AzureSystemProperties) GetDeadLetterSource() string {
	if x != nil {
		return x.DeadLetterSource
	}
	return ""
}

func (x *AzureSystemProperties) GetScheduledEnqueueTime() int64 {
	if x != nil {
		return x.ScheduledEnqueueTime
	}
	return 0
}

func (x *AzureSystemProperties) GetEnqueuedSequenceNumber() int64 {
	if x != nil {
		return x.EnqueuedSequenceNumber
	}
	return 0
}

func (x *AzureSystemProperties) GetViaPartitionKey() string {
	if x != nil {
		return x.ViaPartitionKey
	}
	return ""
}

func (x *AzureSystemProperties) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_azure_proto protoreflect.FileDescriptor

var file_azure_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xb4, 0x05, 0x0a, 0x0f, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x29, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x4b, 0x0a, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x10, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x55, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x1a, 0x41, 0x0a, 0x13, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xad, 0x04,
	0x0a, 0x15, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34,
	0x0a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x45, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x76, 0x69, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x61, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a,
	0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x59, 0x0a,
	0x17, 0x73, 0x68, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_azure_proto_rawDescOnce sync.Once
	file_azure_proto_rawDescData = file_azure_proto_rawDesc
)

func file_azure_proto_rawDescGZIP() []byte {
	file_azure_proto_rawDescOnce.Do(func() {
		file_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_azure_proto_rawDescData)
	})
	return file_azure_proto_rawDescData
}

var file_azure_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_azure_proto_goTypes = []interface{}{
	(*AzureSinkRecord)(nil),       // 0: records.AzureSinkRecord
	(*AzureSystemProperties)(nil), // 1: records.AzureSystemProperties
	nil,                           // 2: records.AzureSinkRecord.UserPropertiesEntry
	nil,                           // 3: records.AzureSystemProperties.AnnotationsEntry
}
var file_azure_proto_depIdxs = []int32{
	1, // 0: records.AzureSinkRecord.system_properties:type_name -> records.AzureSystemProperties
	2, // 1: records.AzureSinkRecord.user_properties:type_name -> records.AzureSinkRecord.UserPropertiesEntry
	3, // 2: records.AzureSystemProperties.annotations:type_name -> records.AzureSystemProperties.AnnotationsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_azure_proto_init() }
func file_azure_proto_init() {
	if File_azure_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_azure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureSinkRecord); i {
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
		file_azure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzureSystemProperties); i {
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
			RawDescriptor: file_azure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_azure_proto_goTypes,
		DependencyIndexes: file_azure_proto_depIdxs,
		MessageInfos:      file_azure_proto_msgTypes,
	}.Build()
	File_azure_proto = out.File
	file_azure_proto_rawDesc = nil
	file_azure_proto_goTypes = nil
	file_azure_proto_depIdxs = nil
}
