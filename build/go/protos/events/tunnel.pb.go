// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tunnel.proto

package events

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

type Tunnel_Type int32

const (
	Tunnel_UNSET Tunnel_Type = 0
	// Sent by plumber to dProxy
	Tunnel_AUTH_REQUEST Tunnel_Type = 1
	// Sent by dProxy to plumber
	Tunnel_AUTH_RESPONSE Tunnel_Type = 2
	// Sent by dProxy to plumber
	// Contains an events.Outbound message with a replay payload
	Tunnel_OUTBOUND_MESSAGE Tunnel_Type = 3
	// Sent by dProxy to plumber
	// Replicates an ISB replay event message for display in plumber logs
	Tunnel_REPLAY_EVENT Tunnel_Type = 4
	// Sent by dProxy to plumber
	// Currently only used for force-disconnecting an ephemeral tunnel
	Tunnel_TUNNEL_EVENT Tunnel_Type = 5
)

var Tunnel_Type_name = map[int32]string{
	0: "UNSET",
	1: "AUTH_REQUEST",
	2: "AUTH_RESPONSE",
	3: "OUTBOUND_MESSAGE",
	4: "REPLAY_EVENT",
	5: "TUNNEL_EVENT",
}

var Tunnel_Type_value = map[string]int32{
	"UNSET":            0,
	"AUTH_REQUEST":     1,
	"AUTH_RESPONSE":    2,
	"OUTBOUND_MESSAGE": 3,
	"REPLAY_EVENT":     4,
	"TUNNEL_EVENT":     5,
}

func (x Tunnel_Type) String() string {
	return proto.EnumName(Tunnel_Type_name, int32(x))
}

func (Tunnel_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0, 0}
}

type ReplayEvent_Type int32

const (
	ReplayEvent_UNSET         ReplayEvent_Type = 0
	ReplayEvent_CREATE_REPLAY ReplayEvent_Type = 1
	ReplayEvent_PAUSE_REPLAY  ReplayEvent_Type = 2
	ReplayEvent_RESUME_REPLAY ReplayEvent_Type = 3
	ReplayEvent_ABORT_REPLAY  ReplayEvent_Type = 4
	ReplayEvent_FINISH_REPLAY ReplayEvent_Type = 5
	ReplayEvent_DISCONNECT    ReplayEvent_Type = 6
)

var ReplayEvent_Type_name = map[int32]string{
	0: "UNSET",
	1: "CREATE_REPLAY",
	2: "PAUSE_REPLAY",
	3: "RESUME_REPLAY",
	4: "ABORT_REPLAY",
	5: "FINISH_REPLAY",
	6: "DISCONNECT",
}

var ReplayEvent_Type_value = map[string]int32{
	"UNSET":         0,
	"CREATE_REPLAY": 1,
	"PAUSE_REPLAY":  2,
	"RESUME_REPLAY": 3,
	"ABORT_REPLAY":  4,
	"FINISH_REPLAY": 5,
	"DISCONNECT":    6,
}

func (x ReplayEvent_Type) String() string {
	return proto.EnumName(ReplayEvent_Type_name, int32(x))
}

func (ReplayEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{5, 0}
}

type TunnelEvent_Type int32

const (
	TunnelEvent_UNSET TunnelEvent_Type = 0
	// Used when an ephemeral tunnel destination is deleted from ui-bff
	// This action will cause dProxy to send this event, forcing plumber to disconnect
	TunnelEvent_FORCE_DISCONNECT TunnelEvent_Type = 1
)

var TunnelEvent_Type_name = map[int32]string{
	0: "UNSET",
	1: "FORCE_DISCONNECT",
}

var TunnelEvent_Type_value = map[string]int32{
	"UNSET":            0,
	"FORCE_DISCONNECT": 1,
}

func (x TunnelEvent_Type) String() string {
	return proto.EnumName(TunnelEvent_Type_name, int32(x))
}

func (TunnelEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{6, 0}
}

// Tunnel is an envelope message for tunnel communication between dproxy and plumber
type Tunnel struct {
	Type     Tunnel_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.Tunnel_Type" json:"type,omitempty"`
	ReplayId string      `protobuf:"bytes,2,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Tunnel_AuthRequest
	//	*Tunnel_AuthResponse
	//	*Tunnel_OutboundMessage
	//	*Tunnel_ReplayMessage
	//	*Tunnel_TunnelEvent
	Payload              isTunnel_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Tunnel) Reset()         { *m = Tunnel{} }
func (m *Tunnel) String() string { return proto.CompactTextString(m) }
func (*Tunnel) ProtoMessage()    {}
func (*Tunnel) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}

func (m *Tunnel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tunnel.Unmarshal(m, b)
}
func (m *Tunnel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tunnel.Marshal(b, m, deterministic)
}
func (m *Tunnel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tunnel.Merge(m, src)
}
func (m *Tunnel) XXX_Size() int {
	return xxx_messageInfo_Tunnel.Size(m)
}
func (m *Tunnel) XXX_DiscardUnknown() {
	xxx_messageInfo_Tunnel.DiscardUnknown(m)
}

var xxx_messageInfo_Tunnel proto.InternalMessageInfo

func (m *Tunnel) GetType() Tunnel_Type {
	if m != nil {
		return m.Type
	}
	return Tunnel_UNSET
}

func (m *Tunnel) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

type isTunnel_Payload interface {
	isTunnel_Payload()
}

type Tunnel_AuthRequest struct {
	AuthRequest *AuthRequest `protobuf:"bytes,100,opt,name=auth_request,json=authRequest,proto3,oneof"`
}

type Tunnel_AuthResponse struct {
	AuthResponse *AuthResponse `protobuf:"bytes,101,opt,name=auth_response,json=authResponse,proto3,oneof"`
}

type Tunnel_OutboundMessage struct {
	OutboundMessage *Outbound `protobuf:"bytes,102,opt,name=outbound_message,json=outboundMessage,proto3,oneof"`
}

type Tunnel_ReplayMessage struct {
	ReplayMessage *ReplayEvent `protobuf:"bytes,103,opt,name=replay_message,json=replayMessage,proto3,oneof"`
}

type Tunnel_TunnelEvent struct {
	TunnelEvent *TunnelEvent `protobuf:"bytes,104,opt,name=tunnel_event,json=tunnelEvent,proto3,oneof"`
}

func (*Tunnel_AuthRequest) isTunnel_Payload() {}

func (*Tunnel_AuthResponse) isTunnel_Payload() {}

func (*Tunnel_OutboundMessage) isTunnel_Payload() {}

func (*Tunnel_ReplayMessage) isTunnel_Payload() {}

func (*Tunnel_TunnelEvent) isTunnel_Payload() {}

func (m *Tunnel) GetPayload() isTunnel_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Tunnel) GetAuthRequest() *AuthRequest {
	if x, ok := m.GetPayload().(*Tunnel_AuthRequest); ok {
		return x.AuthRequest
	}
	return nil
}

func (m *Tunnel) GetAuthResponse() *AuthResponse {
	if x, ok := m.GetPayload().(*Tunnel_AuthResponse); ok {
		return x.AuthResponse
	}
	return nil
}

func (m *Tunnel) GetOutboundMessage() *Outbound {
	if x, ok := m.GetPayload().(*Tunnel_OutboundMessage); ok {
		return x.OutboundMessage
	}
	return nil
}

func (m *Tunnel) GetReplayMessage() *ReplayEvent {
	if x, ok := m.GetPayload().(*Tunnel_ReplayMessage); ok {
		return x.ReplayMessage
	}
	return nil
}

func (m *Tunnel) GetTunnelEvent() *TunnelEvent {
	if x, ok := m.GetPayload().(*Tunnel_TunnelEvent); ok {
		return x.TunnelEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Tunnel) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Tunnel_AuthRequest)(nil),
		(*Tunnel_AuthResponse)(nil),
		(*Tunnel_OutboundMessage)(nil),
		(*Tunnel_ReplayMessage)(nil),
		(*Tunnel_TunnelEvent)(nil),
	}
}

type AuthRequest struct {
	// API Token generated in https://console.batch.sh
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Message bus type name, ex: kafka, rabbitmq, etc.
	MessageBus string `protobuf:"bytes,2,opt,name=message_bus,json=messageBus,proto3" json:"message_bus,omitempty"`
	// Tunnels can be given a custom name to help identify them in the plumber server.
	// If this value is empty, a default name will be generated based off of message_bus and the
	// connecting IP address
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Plumber server mode needs permanent tunnels without creating multiple
	// destinations in console for all running plumber instances  in a cluster.
	// If cluster ID is specified, we will use that + team ID + connection ID = unique identifier
	// so that we can always reconnect to the same tunnel.
	//
	// Tunnels without this value specified will be created as ephemeral tunnels.
	// Ephemeral tunnels will be deleted after 3 hours of inactivity.
	PlumberClusterId     string   `protobuf:"bytes,4,opt,name=plumber_cluster_id,json=plumberClusterId,proto3" json:"plumber_cluster_id,omitempty"`
	TunnelId             string   `protobuf:"bytes,5,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthRequest) GetMessageBus() string {
	if m != nil {
		return m.MessageBus
	}
	return ""
}

func (m *AuthRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRequest) GetPlumberClusterId() string {
	if m != nil {
		return m.PlumberClusterId
	}
	return ""
}

func (m *AuthRequest) GetTunnelId() string {
	if m != nil {
		return m.TunnelId
	}
	return ""
}

type AuthResponse struct {
	// Whether or not the connection is authorized
	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	// Error message if any
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteTunnelRequest struct {
	// API Token generated in https://console.batch.sh
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PlumberClusterId     string   `protobuf:"bytes,2,opt,name=plumber_cluster_id,json=plumberClusterId,proto3" json:"plumber_cluster_id,omitempty"`
	TunnelId             string   `protobuf:"bytes,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTunnelRequest) Reset()         { *m = DeleteTunnelRequest{} }
func (m *DeleteTunnelRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTunnelRequest) ProtoMessage()    {}
func (*DeleteTunnelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{3}
}

func (m *DeleteTunnelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTunnelRequest.Unmarshal(m, b)
}
func (m *DeleteTunnelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTunnelRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTunnelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTunnelRequest.Merge(m, src)
}
func (m *DeleteTunnelRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTunnelRequest.Size(m)
}
func (m *DeleteTunnelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTunnelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTunnelRequest proto.InternalMessageInfo

func (m *DeleteTunnelRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeleteTunnelRequest) GetPlumberClusterId() string {
	if m != nil {
		return m.PlumberClusterId
	}
	return ""
}

func (m *DeleteTunnelRequest) GetTunnelId() string {
	if m != nil {
		return m.TunnelId
	}
	return ""
}

type DeleteTunnelResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTunnelResponse) Reset()         { *m = DeleteTunnelResponse{} }
func (m *DeleteTunnelResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTunnelResponse) ProtoMessage()    {}
func (*DeleteTunnelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{4}
}

func (m *DeleteTunnelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTunnelResponse.Unmarshal(m, b)
}
func (m *DeleteTunnelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTunnelResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTunnelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTunnelResponse.Merge(m, src)
}
func (m *DeleteTunnelResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTunnelResponse.Size(m)
}
func (m *DeleteTunnelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTunnelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTunnelResponse proto.InternalMessageInfo

func (m *DeleteTunnelResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteTunnelResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReplayEvent struct {
	Type                 ReplayEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.ReplayEvent_Type" json:"type,omitempty"`
	ReplayId             string           `protobuf:"bytes,2,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplayEvent) Reset()         { *m = ReplayEvent{} }
func (m *ReplayEvent) String() string { return proto.CompactTextString(m) }
func (*ReplayEvent) ProtoMessage()    {}
func (*ReplayEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{5}
}

func (m *ReplayEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplayEvent.Unmarshal(m, b)
}
func (m *ReplayEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplayEvent.Marshal(b, m, deterministic)
}
func (m *ReplayEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplayEvent.Merge(m, src)
}
func (m *ReplayEvent) XXX_Size() int {
	return xxx_messageInfo_ReplayEvent.Size(m)
}
func (m *ReplayEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplayEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReplayEvent proto.InternalMessageInfo

func (m *ReplayEvent) GetType() ReplayEvent_Type {
	if m != nil {
		return m.Type
	}
	return ReplayEvent_UNSET
}

func (m *ReplayEvent) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

type TunnelEvent struct {
	TunnelId             string   `protobuf:"bytes,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TunnelEvent) Reset()         { *m = TunnelEvent{} }
func (m *TunnelEvent) String() string { return proto.CompactTextString(m) }
func (*TunnelEvent) ProtoMessage()    {}
func (*TunnelEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{6}
}

func (m *TunnelEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelEvent.Unmarshal(m, b)
}
func (m *TunnelEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelEvent.Marshal(b, m, deterministic)
}
func (m *TunnelEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelEvent.Merge(m, src)
}
func (m *TunnelEvent) XXX_Size() int {
	return xxx_messageInfo_TunnelEvent.Size(m)
}
func (m *TunnelEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelEvent proto.InternalMessageInfo

func (m *TunnelEvent) GetTunnelId() string {
	if m != nil {
		return m.TunnelId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.Tunnel_Type", Tunnel_Type_name, Tunnel_Type_value)
	proto.RegisterEnum("events.ReplayEvent_Type", ReplayEvent_Type_name, ReplayEvent_Type_value)
	proto.RegisterEnum("events.TunnelEvent_Type", TunnelEvent_Type_name, TunnelEvent_Type_value)
	proto.RegisterType((*Tunnel)(nil), "events.Tunnel")
	proto.RegisterType((*AuthRequest)(nil), "events.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "events.AuthResponse")
	proto.RegisterType((*DeleteTunnelRequest)(nil), "events.DeleteTunnelRequest")
	proto.RegisterType((*DeleteTunnelResponse)(nil), "events.DeleteTunnelResponse")
	proto.RegisterType((*ReplayEvent)(nil), "events.ReplayEvent")
	proto.RegisterType((*TunnelEvent)(nil), "events.TunnelEvent")
}

func init() { proto.RegisterFile("tunnel.proto", fileDescriptor_6f51ddaa7891a711) }

var fileDescriptor_6f51ddaa7891a711 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0xc5, 0xfc, 0x25, 0x5c, 0x7e, 0x3e, 0x67, 0xc2, 0xc2, 0xfa, 0x2a, 0xb5, 0x11, 0x9b, 0x64,
	0x91, 0x82, 0xd4, 0x6e, 0x22, 0xb5, 0x59, 0x00, 0x99, 0x14, 0xaa, 0xc4, 0xd0, 0xb1, 0x5d, 0xa9,
	0xdd, 0x58, 0xc6, 0x9e, 0x02, 0xaa, 0xb1, 0x5d, 0xcf, 0x38, 0x12, 0x5d, 0xf4, 0x65, 0xfa, 0x52,
	0x7d, 0x95, 0xee, 0x2a, 0xcf, 0xd8, 0x89, 0x49, 0xa2, 0x28, 0x3b, 0xee, 0xb9, 0xe7, 0x0c, 0x73,
	0xee, 0x3d, 0x63, 0x68, 0xf1, 0x24, 0x08, 0xa8, 0xdf, 0x8f, 0xe2, 0x90, 0x87, 0xa8, 0x4e, 0x6f,
	0x68, 0xc0, 0xd9, 0xff, 0x9d, 0x30, 0xe1, 0x8b, 0x30, 0x09, 0x3c, 0x89, 0xf7, 0xfe, 0x56, 0xa0,
	0x6e, 0x0a, 0x22, 0x3a, 0x86, 0x2a, 0xdf, 0x46, 0x54, 0x53, 0x8e, 0x94, 0x93, 0xce, 0x9b, 0xc3,
	0xbe, 0x54, 0xf4, 0x65, 0xb7, 0x6f, 0x6e, 0x23, 0x4a, 0x04, 0x01, 0xbd, 0x80, 0x46, 0x4c, 0x23,
	0xdf, 0xd9, 0xda, 0x6b, 0x4f, 0x2b, 0x1f, 0x29, 0x27, 0x0d, 0xb2, 0x2f, 0x81, 0xa9, 0x87, 0xce,
	0xa0, 0xe5, 0x24, 0x7c, 0x65, 0xc7, 0xf4, 0x47, 0x42, 0x19, 0xd7, 0xbc, 0x23, 0xe5, 0xa4, 0x79,
	0x77, 0xda, 0x30, 0xe1, 0x2b, 0x22, 0x5b, 0x93, 0x12, 0x69, 0x3a, 0x77, 0x25, 0x7a, 0x07, 0xed,
	0x4c, 0xc9, 0xa2, 0x30, 0x60, 0x54, 0xa3, 0x42, 0xda, 0xdd, 0x95, 0xca, 0xde, 0xa4, 0x44, 0x5a,
	0x4e, 0xa1, 0x46, 0xe7, 0xa0, 0xe6, 0xce, 0xec, 0x0d, 0x65, 0xcc, 0x59, 0x52, 0xed, 0x9b, 0xd0,
	0xab, 0xb9, 0x7e, 0x96, 0xf5, 0x27, 0x25, 0xf2, 0x5f, 0xce, 0xbd, 0x96, 0x54, 0xf4, 0x1e, 0x3a,
	0x99, 0xa5, 0x5c, 0xbc, 0xdc, 0xbd, 0x37, 0x11, 0x5d, 0x9c, 0x16, 0x93, 0x12, 0x69, 0x4b, 0x72,
	0xae, 0x3e, 0xcb, 0x87, 0x6d, 0x0b, 0xb6, 0xb6, 0xda, 0xd5, 0xca, 0x09, 0xe6, 0xda, 0x26, 0xbf,
	0x2b, 0x7b, 0x11, 0x54, 0xd3, 0xc1, 0xa2, 0x06, 0xd4, 0x2c, 0xdd, 0xc0, 0xa6, 0x5a, 0x42, 0x2a,
	0xb4, 0x86, 0x96, 0x39, 0xb1, 0x09, 0xfe, 0x64, 0x61, 0xc3, 0x54, 0x15, 0x74, 0x00, 0xed, 0x0c,
	0x31, 0xe6, 0x33, 0xdd, 0xc0, 0x6a, 0x19, 0x75, 0x41, 0x9d, 0x59, 0xe6, 0x68, 0x66, 0xe9, 0x17,
	0xf6, 0x35, 0x36, 0x8c, 0xe1, 0x07, 0xac, 0x56, 0x52, 0x29, 0xc1, 0xf3, 0xab, 0xe1, 0x17, 0x1b,
	0x7f, 0xc6, 0xba, 0xa9, 0x56, 0x53, 0xc4, 0xb4, 0x74, 0x1d, 0x5f, 0x65, 0x48, 0x6d, 0xd4, 0x80,
	0xbd, 0xc8, 0xd9, 0xfa, 0xa1, 0xe3, 0xf5, 0x7e, 0x2b, 0xd0, 0x2c, 0xec, 0x03, 0x75, 0xa1, 0xc6,
	0xc3, 0xef, 0x34, 0x10, 0x09, 0x68, 0x10, 0x59, 0xa0, 0x57, 0xd0, 0xcc, 0x66, 0x62, 0x2f, 0x12,
	0x96, 0xed, 0x1b, 0x32, 0x68, 0x94, 0x30, 0x84, 0xa0, 0x1a, 0x38, 0x1b, 0xaa, 0x55, 0x44, 0x47,
	0xfc, 0x46, 0xa7, 0x80, 0x22, 0x3f, 0xd9, 0x2c, 0x68, 0x6c, 0xbb, 0x7e, 0xc2, 0x38, 0x8d, 0xd3,
	0xac, 0x54, 0x05, 0x43, 0xcd, 0x3a, 0x63, 0xd9, 0x98, 0x7a, 0x69, 0xa0, 0xb2, 0xf9, 0xad, 0x3d,
	0xad, 0x26, 0x03, 0x25, 0x81, 0xa9, 0xd7, 0x9b, 0x40, 0xab, 0xb8, 0x79, 0xf4, 0x12, 0x20, 0xdd,
	0x7c, 0x18, 0xaf, 0x7f, 0x52, 0x4f, 0x5c, 0x75, 0x9f, 0x14, 0x10, 0xa4, 0xc1, 0x5e, 0xbe, 0x43,
	0x79, 0xd7, 0xbc, 0xec, 0xdd, 0xc0, 0xe1, 0x05, 0xf5, 0x29, 0xa7, 0x72, 0x21, 0x4f, 0xdb, 0x7e,
	0xdc, 0x41, 0xf9, 0x39, 0x0e, 0x2a, 0xf7, 0x1c, 0x7c, 0x84, 0xee, 0xee, 0xff, 0x66, 0x4e, 0x34,
	0xd8, 0x63, 0x89, 0xeb, 0x52, 0xc6, 0x32, 0x1b, 0x79, 0xf9, 0x84, 0x87, 0x3f, 0x0a, 0x34, 0x0b,
	0x59, 0x44, 0xa7, 0x3b, 0x8f, 0x56, 0x7b, 0x24, 0xae, 0xcf, 0x7d, 0xb9, 0xbd, 0x5f, 0x0f, 0xb3,
	0x78, 0x00, 0xed, 0x31, 0xc1, 0x43, 0x13, 0xdb, 0x32, 0x57, 0xaa, 0x92, 0x26, 0x6a, 0x3e, 0xb4,
	0x8c, 0x5b, 0xa4, 0x9c, 0x92, 0x08, 0x36, 0xac, 0xeb, 0x5b, 0x48, 0x04, 0x71, 0x38, 0x9a, 0x11,
	0x33, 0x47, 0xaa, 0x29, 0xe9, 0x72, 0xaa, 0x4f, 0x8d, 0x49, 0x0e, 0xd5, 0x50, 0x07, 0xe0, 0x62,
	0x6a, 0x8c, 0x67, 0xba, 0x8e, 0xc7, 0xa6, 0x5a, 0xef, 0x19, 0xd0, 0x2c, 0xbc, 0x94, 0xdd, 0x91,
	0x2a, 0xf7, 0x46, 0x7a, 0xfc, 0xf0, 0xae, 0x5d, 0x50, 0x2f, 0x67, 0x64, 0x8c, 0xed, 0xc2, 0xa1,
	0xca, 0xc8, 0x84, 0x03, 0xb6, 0xea, 0x2f, 0x1c, 0xee, 0xae, 0xfa, 0x8c, 0xc6, 0x37, 0x6b, 0x97,
	0xb2, 0xb9, 0xf2, 0xf5, 0x7c, 0xb9, 0xe6, 0xab, 0x64, 0xd1, 0x77, 0xc3, 0xcd, 0x40, 0x34, 0xdd,
	0x30, 0x8e, 0x06, 0x6e, 0xe8, 0xfb, 0xd4, 0xe5, 0x61, 0xfc, 0x9a, 0xb9, 0x2b, 0xba, 0x71, 0xd8,
	0x60, 0x91, 0xac, 0x7d, 0x6f, 0xb0, 0x0c, 0x07, 0xe2, 0x6b, 0xc9, 0x06, 0x72, 0xbc, 0x8b, 0xba,
	0x28, 0xdf, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xde, 0xc0, 0xda, 0x4f, 0x64, 0x05, 0x00, 0x00,
}
