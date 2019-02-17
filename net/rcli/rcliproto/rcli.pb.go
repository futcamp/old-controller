// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rcli.proto

package rcli

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	LoginHash            string   `protobuf:"bytes,1,opt,name=LoginHash,proto3" json:"LoginHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLoginHash() string {
	if m != nil {
		return m.LoginHash
	}
	return ""
}

type LoginResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type CmdRequest struct {
	LoginHash            string   `protobuf:"bytes,1,opt,name=LoginHash,proto3" json:"LoginHash,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdRequest) Reset()         { *m = CmdRequest{} }
func (m *CmdRequest) String() string { return proto.CompactTextString(m) }
func (*CmdRequest) ProtoMessage()    {}
func (*CmdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{2}
}
func (m *CmdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdRequest.Unmarshal(m, b)
}
func (m *CmdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdRequest.Marshal(b, m, deterministic)
}
func (m *CmdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdRequest.Merge(m, src)
}
func (m *CmdRequest) XXX_Size() int {
	return xxx_messageInfo_CmdRequest.Size(m)
}
func (m *CmdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CmdRequest proto.InternalMessageInfo

func (m *CmdRequest) GetLoginHash() string {
	if m != nil {
		return m.LoginHash
	}
	return ""
}

func (m *CmdRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type CmdResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdResponse) Reset()         { *m = CmdResponse{} }
func (m *CmdResponse) String() string { return proto.CompactTextString(m) }
func (*CmdResponse) ProtoMessage()    {}
func (*CmdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{3}
}
func (m *CmdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdResponse.Unmarshal(m, b)
}
func (m *CmdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdResponse.Marshal(b, m, deterministic)
}
func (m *CmdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdResponse.Merge(m, src)
}
func (m *CmdResponse) XXX_Size() int {
	return xxx_messageInfo_CmdResponse.Size(m)
}
func (m *CmdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CmdResponse proto.InternalMessageInfo

func (m *CmdResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type CmdListRequest struct {
	LoginHash            string   `protobuf:"bytes,1,opt,name=LoginHash,proto3" json:"LoginHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdListRequest) Reset()         { *m = CmdListRequest{} }
func (m *CmdListRequest) String() string { return proto.CompactTextString(m) }
func (*CmdListRequest) ProtoMessage()    {}
func (*CmdListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{4}
}
func (m *CmdListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdListRequest.Unmarshal(m, b)
}
func (m *CmdListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdListRequest.Marshal(b, m, deterministic)
}
func (m *CmdListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdListRequest.Merge(m, src)
}
func (m *CmdListRequest) XXX_Size() int {
	return xxx_messageInfo_CmdListRequest.Size(m)
}
func (m *CmdListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CmdListRequest proto.InternalMessageInfo

func (m *CmdListRequest) GetLoginHash() string {
	if m != nil {
		return m.LoginHash
	}
	return ""
}

type CmdListResponse struct {
	Cmds                 []string `protobuf:"bytes,1,rep,name=Cmds,proto3" json:"Cmds,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdListResponse) Reset()         { *m = CmdListResponse{} }
func (m *CmdListResponse) String() string { return proto.CompactTextString(m) }
func (*CmdListResponse) ProtoMessage()    {}
func (*CmdListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{5}
}
func (m *CmdListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdListResponse.Unmarshal(m, b)
}
func (m *CmdListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdListResponse.Marshal(b, m, deterministic)
}
func (m *CmdListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdListResponse.Merge(m, src)
}
func (m *CmdListResponse) XXX_Size() int {
	return xxx_messageInfo_CmdListResponse.Size(m)
}
func (m *CmdListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CmdListResponse proto.InternalMessageInfo

func (m *CmdListResponse) GetCmds() []string {
	if m != nil {
		return m.Cmds
	}
	return nil
}

func (m *CmdListResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type UpdateStartupRequest struct {
	LoginHash            string   `protobuf:"bytes,1,opt,name=LoginHash,proto3" json:"LoginHash,omitempty"`
	Cmds                 []string `protobuf:"bytes,2,rep,name=Cmds,proto3" json:"Cmds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStartupRequest) Reset()         { *m = UpdateStartupRequest{} }
func (m *UpdateStartupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStartupRequest) ProtoMessage()    {}
func (*UpdateStartupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{6}
}
func (m *UpdateStartupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStartupRequest.Unmarshal(m, b)
}
func (m *UpdateStartupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStartupRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStartupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStartupRequest.Merge(m, src)
}
func (m *UpdateStartupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStartupRequest.Size(m)
}
func (m *UpdateStartupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStartupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStartupRequest proto.InternalMessageInfo

func (m *UpdateStartupRequest) GetLoginHash() string {
	if m != nil {
		return m.LoginHash
	}
	return ""
}

func (m *UpdateStartupRequest) GetCmds() []string {
	if m != nil {
		return m.Cmds
	}
	return nil
}

type UpdateStartupResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStartupResponse) Reset()         { *m = UpdateStartupResponse{} }
func (m *UpdateStartupResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateStartupResponse) ProtoMessage()    {}
func (*UpdateStartupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{7}
}
func (m *UpdateStartupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStartupResponse.Unmarshal(m, b)
}
func (m *UpdateStartupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStartupResponse.Marshal(b, m, deterministic)
}
func (m *UpdateStartupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStartupResponse.Merge(m, src)
}
func (m *UpdateStartupResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateStartupResponse.Size(m)
}
func (m *UpdateStartupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStartupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStartupResponse proto.InternalMessageInfo

func (m *UpdateStartupResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type LinuxCmdRequest struct {
	LoginHash            string   `protobuf:"bytes,1,opt,name=LoginHash,proto3" json:"LoginHash,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	Args                 string   `protobuf:"bytes,3,opt,name=Args,proto3" json:"Args,omitempty"`
	SubArgs              string   `protobuf:"bytes,4,opt,name=SubArgs,proto3" json:"SubArgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxCmdRequest) Reset()         { *m = LinuxCmdRequest{} }
func (m *LinuxCmdRequest) String() string { return proto.CompactTextString(m) }
func (*LinuxCmdRequest) ProtoMessage()    {}
func (*LinuxCmdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{8}
}
func (m *LinuxCmdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxCmdRequest.Unmarshal(m, b)
}
func (m *LinuxCmdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxCmdRequest.Marshal(b, m, deterministic)
}
func (m *LinuxCmdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxCmdRequest.Merge(m, src)
}
func (m *LinuxCmdRequest) XXX_Size() int {
	return xxx_messageInfo_LinuxCmdRequest.Size(m)
}
func (m *LinuxCmdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxCmdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxCmdRequest proto.InternalMessageInfo

func (m *LinuxCmdRequest) GetLoginHash() string {
	if m != nil {
		return m.LoginHash
	}
	return ""
}

func (m *LinuxCmdRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *LinuxCmdRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *LinuxCmdRequest) GetSubArgs() string {
	if m != nil {
		return m.SubArgs
	}
	return ""
}

type LinuxCmdResponse struct {
	Output               string   `protobuf:"bytes,1,opt,name=Output,proto3" json:"Output,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxCmdResponse) Reset()         { *m = LinuxCmdResponse{} }
func (m *LinuxCmdResponse) String() string { return proto.CompactTextString(m) }
func (*LinuxCmdResponse) ProtoMessage()    {}
func (*LinuxCmdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d14d85b7ce07a3f0, []int{9}
}
func (m *LinuxCmdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxCmdResponse.Unmarshal(m, b)
}
func (m *LinuxCmdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxCmdResponse.Marshal(b, m, deterministic)
}
func (m *LinuxCmdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxCmdResponse.Merge(m, src)
}
func (m *LinuxCmdResponse) XXX_Size() int {
	return xxx_messageInfo_LinuxCmdResponse.Size(m)
}
func (m *LinuxCmdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxCmdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxCmdResponse proto.InternalMessageInfo

func (m *LinuxCmdResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *LinuxCmdResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "rcli.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "rcli.LoginResponse")
	proto.RegisterType((*CmdRequest)(nil), "rcli.CmdRequest")
	proto.RegisterType((*CmdResponse)(nil), "rcli.CmdResponse")
	proto.RegisterType((*CmdListRequest)(nil), "rcli.CmdListRequest")
	proto.RegisterType((*CmdListResponse)(nil), "rcli.CmdListResponse")
	proto.RegisterType((*UpdateStartupRequest)(nil), "rcli.UpdateStartupRequest")
	proto.RegisterType((*UpdateStartupResponse)(nil), "rcli.UpdateStartupResponse")
	proto.RegisterType((*LinuxCmdRequest)(nil), "rcli.LinuxCmdRequest")
	proto.RegisterType((*LinuxCmdResponse)(nil), "rcli.LinuxCmdResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteCliClient is the client API for RemoteCli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteCliClient interface {
	LoginCheck(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SendCmd(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdResponse, error)
	ComandsList(ctx context.Context, in *CmdListRequest, opts ...grpc.CallOption) (*CmdListResponse, error)
	UpdateStartup(ctx context.Context, in *UpdateStartupRequest, opts ...grpc.CallOption) (*UpdateStartupResponse, error)
	SendLinuxCmd(ctx context.Context, in *LinuxCmdRequest, opts ...grpc.CallOption) (*LinuxCmdResponse, error)
}

type remoteCliClient struct {
	cc *grpc.ClientConn
}

func NewRemoteCliClient(cc *grpc.ClientConn) RemoteCliClient {
	return &remoteCliClient{cc}
}

func (c *remoteCliClient) LoginCheck(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/rcli.RemoteCli/LoginCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteCliClient) SendCmd(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdResponse, error) {
	out := new(CmdResponse)
	err := c.cc.Invoke(ctx, "/rcli.RemoteCli/SendCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteCliClient) ComandsList(ctx context.Context, in *CmdListRequest, opts ...grpc.CallOption) (*CmdListResponse, error) {
	out := new(CmdListResponse)
	err := c.cc.Invoke(ctx, "/rcli.RemoteCli/ComandsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteCliClient) UpdateStartup(ctx context.Context, in *UpdateStartupRequest, opts ...grpc.CallOption) (*UpdateStartupResponse, error) {
	out := new(UpdateStartupResponse)
	err := c.cc.Invoke(ctx, "/rcli.RemoteCli/UpdateStartup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteCliClient) SendLinuxCmd(ctx context.Context, in *LinuxCmdRequest, opts ...grpc.CallOption) (*LinuxCmdResponse, error) {
	out := new(LinuxCmdResponse)
	err := c.cc.Invoke(ctx, "/rcli.RemoteCli/SendLinuxCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteCliServer is the server API for RemoteCli service.
type RemoteCliServer interface {
	LoginCheck(context.Context, *LoginRequest) (*LoginResponse, error)
	SendCmd(context.Context, *CmdRequest) (*CmdResponse, error)
	ComandsList(context.Context, *CmdListRequest) (*CmdListResponse, error)
	UpdateStartup(context.Context, *UpdateStartupRequest) (*UpdateStartupResponse, error)
	SendLinuxCmd(context.Context, *LinuxCmdRequest) (*LinuxCmdResponse, error)
}

func RegisterRemoteCliServer(s *grpc.Server, srv RemoteCliServer) {
	s.RegisterService(&_RemoteCli_serviceDesc, srv)
}

func _RemoteCli_LoginCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCliServer).LoginCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcli.RemoteCli/LoginCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCliServer).LoginCheck(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteCli_SendCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCliServer).SendCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcli.RemoteCli/SendCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCliServer).SendCmd(ctx, req.(*CmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteCli_ComandsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCliServer).ComandsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcli.RemoteCli/ComandsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCliServer).ComandsList(ctx, req.(*CmdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteCli_UpdateStartup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStartupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCliServer).UpdateStartup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcli.RemoteCli/UpdateStartup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCliServer).UpdateStartup(ctx, req.(*UpdateStartupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteCli_SendLinuxCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinuxCmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteCliServer).SendLinuxCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rcli.RemoteCli/SendLinuxCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteCliServer).SendLinuxCmd(ctx, req.(*LinuxCmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteCli_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rcli.RemoteCli",
	HandlerType: (*RemoteCliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginCheck",
			Handler:    _RemoteCli_LoginCheck_Handler,
		},
		{
			MethodName: "SendCmd",
			Handler:    _RemoteCli_SendCmd_Handler,
		},
		{
			MethodName: "ComandsList",
			Handler:    _RemoteCli_ComandsList_Handler,
		},
		{
			MethodName: "UpdateStartup",
			Handler:    _RemoteCli_UpdateStartup_Handler,
		},
		{
			MethodName: "SendLinuxCmd",
			Handler:    _RemoteCli_SendLinuxCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rcli.proto",
}

func init() { proto.RegisterFile("rcli.proto", fileDescriptor_d14d85b7ce07a3f0) }

var fileDescriptor_d14d85b7ce07a3f0 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0x95, 0x42, 0xc0, 0x5e, 0x40, 0x70, 0x04, 0xd2, 0x54, 0x17, 0xa4, 0x89, 0x91, 0x85, 0x41,
	0xa3, 0x0b, 0x37, 0x1a, 0xa3, 0xe3, 0x82, 0x18, 0x12, 0x93, 0x12, 0x3f, 0xa0, 0xd0, 0x09, 0x34,
	0xd2, 0x4e, 0xed, 0xcc, 0x44, 0x7f, 0xcf, 0x3f, 0x33, 0x9d, 0x4e, 0x1f, 0x34, 0x18, 0x48, 0xdc,
	0xdd, 0x7b, 0x7a, 0xe6, 0x9c, 0xfb, 0x2a, 0x40, 0xb4, 0x58, 0x7b, 0xe3, 0x30, 0xa2, 0x9c, 0xa2,
	0x5a, 0x1c, 0x5b, 0x97, 0xd0, 0x9a, 0xd2, 0xa5, 0x17, 0xd8, 0xe4, 0x53, 0x10, 0xc6, 0xd1, 0x19,
	0xe8, 0x32, 0x9f, 0x38, 0x6c, 0x65, 0x54, 0x86, 0x95, 0x91, 0x6e, 0xe7, 0x80, 0x75, 0x01, 0x6d,
	0xc5, 0x66, 0x21, 0x0d, 0x18, 0x41, 0x03, 0xa8, 0xdb, 0x84, 0x89, 0x35, 0x97, 0xdc, 0x43, 0x5b,
	0x65, 0xd6, 0x0b, 0x00, 0xf6, 0xdd, 0xbd, 0x44, 0x91, 0x01, 0x0d, 0x4c, 0x7d, 0xdf, 0x09, 0x5c,
	0x43, 0x93, 0xdf, 0xd2, 0xd4, 0x3a, 0x87, 0xa6, 0x54, 0xd9, 0x61, 0x36, 0x86, 0x23, 0xec, 0xbb,
	0x53, 0x8f, 0xf1, 0xfd, 0xba, 0x78, 0x80, 0x4e, 0xc6, 0x57, 0xd2, 0x08, 0x6a, 0xd8, 0x77, 0x99,
	0x51, 0x19, 0x56, 0x47, 0xba, 0x2d, 0xe3, 0x82, 0x9d, 0xb6, 0x61, 0x37, 0x81, 0xde, 0x7b, 0xe8,
	0x3a, 0x9c, 0xcc, 0xb8, 0x13, 0x71, 0x11, 0xee, 0xd7, 0x65, 0xea, 0xa0, 0xe5, 0x0e, 0xd6, 0x15,
	0xf4, 0x4b, 0x4a, 0x3b, 0x3a, 0xfd, 0x82, 0xce, 0xd4, 0x0b, 0xc4, 0xf7, 0xff, 0x67, 0x1b, 0xd7,
	0xf3, 0x14, 0x2d, 0x99, 0x51, 0x95, 0xb0, 0x8c, 0x63, 0xf6, 0x4c, 0xcc, 0x25, 0x5c, 0x4b, 0xd8,
	0x2a, 0xb5, 0x9e, 0xa1, 0x9b, 0x1b, 0xe7, 0x45, 0xbe, 0x09, 0x1e, 0x0a, 0xae, 0x6c, 0x55, 0xf6,
	0xd7, 0xdc, 0x6e, 0x7e, 0x34, 0xd0, 0x6d, 0xe2, 0x53, 0x4e, 0xf0, 0xda, 0x43, 0x77, 0x00, 0xb2,
	0x4c, 0xbc, 0x22, 0x8b, 0x0f, 0x84, 0xc6, 0xf2, 0x32, 0x8b, 0xa7, 0x68, 0x9e, 0x6c, 0x60, 0x89,
	0xa9, 0x75, 0x80, 0xae, 0xa1, 0x31, 0x23, 0x81, 0x8b, 0x7d, 0x17, 0x75, 0x13, 0x46, 0x3e, 0x0d,
	0xf3, 0xb8, 0x80, 0x64, 0x2f, 0xee, 0xa1, 0x89, 0x69, 0xdc, 0x34, 0x8b, 0x77, 0x8e, 0x7a, 0x19,
	0xa7, 0x70, 0x32, 0x66, 0xbf, 0x84, 0x66, 0xaf, 0x5f, 0xa1, 0xbd, 0xb1, 0x24, 0x64, 0x26, 0xcc,
	0x6d, 0x37, 0x60, 0x9e, 0x6e, 0xfd, 0x96, 0x69, 0x3d, 0x42, 0x2b, 0xae, 0x3d, 0x1d, 0x25, 0x52,
	0xa6, 0xa5, 0x9d, 0x9a, 0x83, 0x32, 0x9c, 0x0a, 0xcc, 0xeb, 0xf2, 0xdf, 0xbd, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xdd, 0x16, 0x58, 0x05, 0xc9, 0x03, 0x00, 0x00,
}
