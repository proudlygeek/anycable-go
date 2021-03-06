// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package anycable

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_ERROR   Status = 0
	Status_SUCCESS Status = 1
	Status_FAILURE Status = 2
)

var Status_name = map[int32]string{
	0: "ERROR",
	1: "SUCCESS",
	2: "FAILURE",
}

var Status_value = map[string]int32{
	"ERROR":   0,
	"SUCCESS": 1,
	"FAILURE": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

type Env struct {
	Url                  string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cstate               map[string]string `protobuf:"bytes,3,rep,name=cstate,proto3" json:"cstate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Istate               map[string]string `protobuf:"bytes,4,rep,name=istate,proto3" json:"istate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Env) Reset()         { *m = Env{} }
func (m *Env) String() string { return proto.CompactTextString(m) }
func (*Env) ProtoMessage()    {}
func (*Env) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Env) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Env.Unmarshal(m, b)
}
func (m *Env) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Env.Marshal(b, m, deterministic)
}
func (m *Env) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Env.Merge(m, src)
}
func (m *Env) XXX_Size() int {
	return xxx_messageInfo_Env.Size(m)
}
func (m *Env) XXX_DiscardUnknown() {
	xxx_messageInfo_Env.DiscardUnknown(m)
}

var xxx_messageInfo_Env proto.InternalMessageInfo

func (m *Env) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Env) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Env) GetCstate() map[string]string {
	if m != nil {
		return m.Cstate
	}
	return nil
}

func (m *Env) GetIstate() map[string]string {
	if m != nil {
		return m.Istate
	}
	return nil
}

type EnvResponse struct {
	Cstate               map[string]string `protobuf:"bytes,1,rep,name=cstate,proto3" json:"cstate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Istate               map[string]string `protobuf:"bytes,2,rep,name=istate,proto3" json:"istate,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EnvResponse) Reset()         { *m = EnvResponse{} }
func (m *EnvResponse) String() string { return proto.CompactTextString(m) }
func (*EnvResponse) ProtoMessage()    {}
func (*EnvResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *EnvResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvResponse.Unmarshal(m, b)
}
func (m *EnvResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvResponse.Marshal(b, m, deterministic)
}
func (m *EnvResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvResponse.Merge(m, src)
}
func (m *EnvResponse) XXX_Size() int {
	return xxx_messageInfo_EnvResponse.Size(m)
}
func (m *EnvResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnvResponse proto.InternalMessageInfo

func (m *EnvResponse) GetCstate() map[string]string {
	if m != nil {
		return m.Cstate
	}
	return nil
}

func (m *EnvResponse) GetIstate() map[string]string {
	if m != nil {
		return m.Istate
	}
	return nil
}

type ConnectionRequest struct {
	Path                 string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env                  *Env              `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(m, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ConnectionRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *ConnectionRequest) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type ConnectionResponse struct {
	Status               Status       `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	Identifiers          string       `protobuf:"bytes,2,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	Transmissions        []string     `protobuf:"bytes,3,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	ErrorMsg             string       `protobuf:"bytes,4,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	Env                  *EnvResponse `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConnectionResponse) Reset()         { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()    {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResponse.Unmarshal(m, b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionResponse.Size(m)
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *ConnectionResponse) GetIdentifiers() string {
	if m != nil {
		return m.Identifiers
	}
	return ""
}

func (m *ConnectionResponse) GetTransmissions() []string {
	if m != nil {
		return m.Transmissions
	}
	return nil
}

func (m *ConnectionResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *ConnectionResponse) GetEnv() *EnvResponse {
	if m != nil {
		return m.Env
	}
	return nil
}

type CommandMessage struct {
	Command               string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Identifier            string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ConnectionIdentifiers string   `protobuf:"bytes,3,opt,name=connection_identifiers,json=connectionIdentifiers,proto3" json:"connection_identifiers,omitempty"`
	Data                  string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Env                   *Env     `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *CommandMessage) Reset()         { *m = CommandMessage{} }
func (m *CommandMessage) String() string { return proto.CompactTextString(m) }
func (*CommandMessage) ProtoMessage()    {}
func (*CommandMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *CommandMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandMessage.Unmarshal(m, b)
}
func (m *CommandMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandMessage.Marshal(b, m, deterministic)
}
func (m *CommandMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMessage.Merge(m, src)
}
func (m *CommandMessage) XXX_Size() int {
	return xxx_messageInfo_CommandMessage.Size(m)
}
func (m *CommandMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMessage proto.InternalMessageInfo

func (m *CommandMessage) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandMessage) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CommandMessage) GetConnectionIdentifiers() string {
	if m != nil {
		return m.ConnectionIdentifiers
	}
	return ""
}

func (m *CommandMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *CommandMessage) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type CommandResponse struct {
	Status               Status       `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	Disconnect           bool         `protobuf:"varint,2,opt,name=disconnect,proto3" json:"disconnect,omitempty"`
	StopStreams          bool         `protobuf:"varint,3,opt,name=stop_streams,json=stopStreams,proto3" json:"stop_streams,omitempty"`
	Streams              []string     `protobuf:"bytes,4,rep,name=streams,proto3" json:"streams,omitempty"`
	Transmissions        []string     `protobuf:"bytes,5,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	ErrorMsg             string       `protobuf:"bytes,6,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	Env                  *EnvResponse `protobuf:"bytes,7,opt,name=env,proto3" json:"env,omitempty"`
	StoppedStreams       []string     `protobuf:"bytes,8,rep,name=stopped_streams,json=stoppedStreams,proto3" json:"stopped_streams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (m *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(m, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func (m *CommandResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *CommandResponse) GetDisconnect() bool {
	if m != nil {
		return m.Disconnect
	}
	return false
}

func (m *CommandResponse) GetStopStreams() bool {
	if m != nil {
		return m.StopStreams
	}
	return false
}

func (m *CommandResponse) GetStreams() []string {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *CommandResponse) GetTransmissions() []string {
	if m != nil {
		return m.Transmissions
	}
	return nil
}

func (m *CommandResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func (m *CommandResponse) GetEnv() *EnvResponse {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *CommandResponse) GetStoppedStreams() []string {
	if m != nil {
		return m.StoppedStreams
	}
	return nil
}

type DisconnectRequest struct {
	Identifiers          string            `protobuf:"bytes,1,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	Subscriptions        []string          `protobuf:"bytes,2,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Path                 string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Headers              map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env                  *Env              `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetIdentifiers() string {
	if m != nil {
		return m.Identifiers
	}
	return ""
}

func (m *DisconnectRequest) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *DisconnectRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DisconnectRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DisconnectRequest) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type DisconnectResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (m *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(m, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *DisconnectResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("anycable.Status", Status_name, Status_value)
	proto.RegisterType((*Env)(nil), "anycable.Env")
	proto.RegisterMapType((map[string]string)(nil), "anycable.Env.CstateEntry")
	proto.RegisterMapType((map[string]string)(nil), "anycable.Env.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "anycable.Env.IstateEntry")
	proto.RegisterType((*EnvResponse)(nil), "anycable.EnvResponse")
	proto.RegisterMapType((map[string]string)(nil), "anycable.EnvResponse.CstateEntry")
	proto.RegisterMapType((map[string]string)(nil), "anycable.EnvResponse.IstateEntry")
	proto.RegisterType((*ConnectionRequest)(nil), "anycable.ConnectionRequest")
	proto.RegisterMapType((map[string]string)(nil), "anycable.ConnectionRequest.HeadersEntry")
	proto.RegisterType((*ConnectionResponse)(nil), "anycable.ConnectionResponse")
	proto.RegisterType((*CommandMessage)(nil), "anycable.CommandMessage")
	proto.RegisterType((*CommandResponse)(nil), "anycable.CommandResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "anycable.DisconnectRequest")
	proto.RegisterMapType((map[string]string)(nil), "anycable.DisconnectRequest.HeadersEntry")
	proto.RegisterType((*DisconnectResponse)(nil), "anycable.DisconnectResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xae, 0xed, 0x5c, 0x8f, 0x7b, 0x49, 0x8f, 0xfe, 0xfe, 0x72, 0x53, 0x54, 0x52, 0x0b, 0xa9,
	0x11, 0x12, 0x91, 0x08, 0x20, 0xd1, 0xae, 0x80, 0xe0, 0x8a, 0x48, 0x54, 0x20, 0x47, 0x5d, 0xb1,
	0xa8, 0x5c, 0x67, 0x68, 0xad, 0x36, 0x63, 0xe3, 0x99, 0x44, 0xca, 0x5b, 0xf0, 0x22, 0x6c, 0x79,
	0x06, 0x16, 0x3c, 0x01, 0x7b, 0x1e, 0x82, 0x1d, 0xf2, 0x8c, 0x9d, 0x8c, 0x1b, 0xb7, 0x55, 0x8b,
	0x10, 0x3b, 0xcf, 0xb9, 0xcc, 0x7c, 0x97, 0x73, 0xa2, 0x40, 0x3d, 0x8e, 0xfc, 0x4e, 0x14, 0x87,
	0x3c, 0xc4, 0x9a, 0x47, 0xa7, 0xbe, 0x77, 0x72, 0x41, 0xec, 0x9f, 0x3a, 0x18, 0x0e, 0x9d, 0x60,
	0x03, 0x8c, 0x71, 0x7c, 0x61, 0x69, 0x2d, 0xad, 0x5d, 0x77, 0x93, 0x4f, 0x7c, 0x0a, 0xd5, 0x33,
	0xe2, 0x0d, 0x49, 0xcc, 0x2c, 0xbd, 0x65, 0xb4, 0xcd, 0x6e, 0xb3, 0x93, 0x75, 0x75, 0x1c, 0x3a,
	0xe9, 0xbc, 0x91, 0x49, 0x87, 0xf2, 0x78, 0xea, 0x66, 0xa5, 0xf8, 0x18, 0x2a, 0x3e, 0xe3, 0x1e,
	0x27, 0x96, 0x21, 0x9a, 0x36, 0xf3, 0x4d, 0x3d, 0x91, 0x93, 0x3d, 0x69, 0x61, 0xd2, 0x12, 0xc8,
	0x96, 0x52, 0x51, 0x4b, 0x5f, 0x6d, 0x91, 0x85, 0xcd, 0x7d, 0x58, 0x56, 0x9f, 0x4f, 0xd0, 0x9f,
	0x93, 0x69, 0x86, 0xfe, 0x9c, 0x4c, 0xf1, 0x3f, 0x28, 0x4f, 0xbc, 0x8b, 0x31, 0xb1, 0x74, 0x11,
	0x93, 0x87, 0x7d, 0xfd, 0xb9, 0xd6, 0xdc, 0x03, 0x53, 0x41, 0x71, 0xdb, 0xd6, 0xfe, 0xdd, 0x5a,
	0xed, 0x5f, 0x1a, 0x98, 0x0e, 0x9d, 0xb8, 0x84, 0x45, 0x21, 0x65, 0x04, 0xf7, 0x66, 0x3a, 0x69,
	0x82, 0xf4, 0x4e, 0x8e, 0x74, 0x56, 0x56, 0xa8, 0xd7, 0xde, 0x4c, 0x2f, 0xfd, 0xba, 0xd6, 0x22,
	0xdd, 0xfe, 0x0d, 0xf7, 0x6f, 0x1a, 0xac, 0xf7, 0x42, 0x4a, 0x89, 0xcf, 0x83, 0x90, 0xba, 0xe4,
	0xd3, 0x98, 0x30, 0x8e, 0x08, 0xa5, 0xc8, 0xe3, 0x67, 0xe9, 0x15, 0xe2, 0x1b, 0x5f, 0x5d, 0x9e,
	0xb9, 0xf6, 0x9c, 0xdb, 0xc2, 0x0d, 0x57, 0x4c, 0xe0, 0x7d, 0x30, 0x08, 0x9d, 0x58, 0x46, 0x4b,
	0x6b, 0x9b, 0xdd, 0x95, 0xbc, 0x36, 0x49, 0xe6, 0x4f, 0x86, 0xc7, 0xfe, 0xae, 0x01, 0xaa, 0x40,
	0x52, 0x37, 0xdb, 0x50, 0x49, 0xb4, 0x19, 0x33, 0x71, 0xcb, 0x6a, 0xb7, 0x31, 0x7f, 0x76, 0x20,
	0xe2, 0x6e, 0x9a, 0xc7, 0x16, 0x98, 0xc1, 0x90, 0x50, 0x1e, 0x7c, 0x0c, 0x24, 0xcb, 0xe4, 0x01,
	0x35, 0x84, 0x0f, 0x60, 0x85, 0xc7, 0x1e, 0x65, 0xa3, 0x80, 0xb1, 0x20, 0xa4, 0x4c, 0x2c, 0x52,
	0xdd, 0xcd, 0x07, 0x71, 0x0b, 0xea, 0x24, 0x8e, 0xc3, 0xf8, 0x78, 0xc4, 0x4e, 0xad, 0x92, 0xb8,
	0xa5, 0x26, 0x02, 0x87, 0xec, 0x14, 0x77, 0xa5, 0x04, 0x65, 0x21, 0xc1, 0x46, 0xe1, 0x78, 0x08,
	0x29, 0xec, 0xaf, 0x1a, 0xac, 0xf6, 0xc2, 0xd1, 0xc8, 0xa3, 0xc3, 0x43, 0xc2, 0x98, 0x77, 0x4a,
	0xd0, 0x82, 0xaa, 0x2f, 0x23, 0xa9, 0x22, 0xd9, 0x11, 0xb7, 0x01, 0xe6, 0x38, 0x53, 0xe4, 0x4a,
	0x04, 0x9f, 0xc1, 0xff, 0xfe, 0x4c, 0x9a, 0x63, 0x95, 0xa5, 0x21, 0x6a, 0x37, 0xe6, 0xd9, 0xbe,
	0xc2, 0x17, 0xa1, 0x34, 0xf4, 0xb8, 0x97, 0x92, 0x10, 0xdf, 0x99, 0x87, 0xe5, 0xab, 0x3c, 0xb4,
	0xbf, 0xe8, 0xb0, 0x96, 0x02, 0xbf, 0x83, 0x09, 0xdb, 0x00, 0xc3, 0x80, 0xa5, 0x70, 0x04, 0x93,
	0x9a, 0xab, 0x44, 0x70, 0x07, 0x96, 0x19, 0x0f, 0xa3, 0x63, 0xc6, 0x63, 0xe2, 0x8d, 0x24, 0xfe,
	0x9a, 0x6b, 0x26, 0xb1, 0x81, 0x0c, 0x25, 0x32, 0x65, 0xd9, 0x92, 0xf0, 0x27, 0x3b, 0x2e, 0xfa,
	0x57, 0xbe, 0xd1, 0xbf, 0x4a, 0xb1, 0x7f, 0xd5, 0x9b, 0xfc, 0xc3, 0x5d, 0x58, 0x4b, 0x40, 0x45,
	0x64, 0x38, 0xc3, 0x5a, 0x13, 0xaf, 0xad, 0xa6, 0xe1, 0x14, 0xae, 0xfd, 0x59, 0x87, 0xf5, 0xd7,
	0x33, 0x82, 0xd9, 0x0a, 0x5e, 0x1a, 0x46, 0xad, 0x70, 0x18, 0xd9, 0xf8, 0x84, 0xf9, 0x71, 0x10,
	0x71, 0x41, 0x46, 0x97, 0x64, 0x72, 0xc1, 0xd9, 0x2a, 0x1b, 0xc5, 0xab, 0x5c, 0xba, 0xbc, 0xca,
	0x0b, 0x48, 0xae, 0x5f, 0xe5, 0xf2, 0x5f, 0x59, 0xe5, 0x0f, 0x80, 0x2a, 0x8e, 0x5b, 0x0f, 0x51,
	0xce, 0x41, 0x3d, 0xef, 0xe0, 0xc3, 0x47, 0x50, 0x91, 0xe5, 0x58, 0x87, 0xb2, 0xe3, 0xba, 0xef,
	0xdc, 0xc6, 0x12, 0x9a, 0x50, 0x1d, 0x1c, 0xf5, 0x7a, 0xce, 0x60, 0xd0, 0xd0, 0x92, 0xc3, 0xc1,
	0xcb, 0xfe, 0xdb, 0x23, 0xd7, 0x69, 0xe8, 0xdd, 0x1f, 0x1a, 0x18, 0xee, 0xfb, 0x1e, 0x1e, 0x40,
	0x35, 0xfd, 0x75, 0xc1, 0xad, 0x6b, 0x7e, 0xf9, 0x9a, 0xf7, 0x8a, 0x93, 0x92, 0x83, 0xbd, 0x84,
	0x2f, 0x92, 0x7b, 0xe4, 0xd6, 0x5a, 0x6a, 0xa9, 0xba, 0xe9, 0xcd, 0xcd, 0x85, 0x8c, 0x72, 0x43,
	0x1f, 0x60, 0xae, 0x8e, 0x0a, 0x66, 0xc1, 0x3b, 0x15, 0xcc, 0xa2, 0xa0, 0xf6, 0xd2, 0x49, 0x45,
	0xfc, 0xe7, 0x78, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x45, 0xe1, 0x26, 0x95, 0x80, 0x08, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCClient interface {
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	Command(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
}

type rPCClient struct {
	cc *grpc.ClientConn
}

func NewRPCClient(cc *grpc.ClientConn) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Command(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
type RPCServer interface {
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	Command(context.Context, *CommandMessage) (*CommandResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
}

// UnimplementedRPCServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (*UnimplementedRPCServer) Connect(ctx context.Context, req *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedRPCServer) Command(ctx context.Context, req *CommandMessage) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Command not implemented")
}
func (*UnimplementedRPCServer) Disconnect(ctx context.Context, req *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterRPCServer(s *grpc.Server, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
}

func _RPC_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Command(ctx, req.(*CommandMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "anycable.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _RPC_Connect_Handler,
		},
		{
			MethodName: "Command",
			Handler:    _RPC_Command_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _RPC_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
