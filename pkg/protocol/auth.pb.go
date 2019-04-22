// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package protocol

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

type GenerateTokenRequest struct {
	Image    string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *GenerateTokenRequest) Reset()                    { *m = GenerateTokenRequest{} }
func (m *GenerateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateTokenRequest) ProtoMessage()               {}
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GenerateTokenRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *GenerateTokenRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type GenerateTokenReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GenerateTokenReply) Reset()                    { *m = GenerateTokenReply{} }
func (m *GenerateTokenReply) String() string            { return proto.CompactTextString(m) }
func (*GenerateTokenReply) ProtoMessage()               {}
func (*GenerateTokenReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GenerateTokenReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateTokenRequest)(nil), "protocol.GenerateTokenRequest")
	proto.RegisterType((*GenerateTokenReply)(nil), "protocol.GenerateTokenReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenReply, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenReply, error) {
	out := new(GenerateTokenReply)
	err := grpc.Invoke(ctx, "/protocol.Auth/GenerateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	GenerateToken(context.Context, *GenerateTokenRequest) (*GenerateTokenReply, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Auth/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GenerateToken(ctx, req.(*GenerateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _Auth_GenerateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x1e, 0x5c,
	0x22, 0xee, 0xa9, 0x79, 0xa9, 0x45, 0x89, 0x25, 0xa9, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x99, 0xb9, 0x89, 0xe9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x4a, 0x69, 0x51, 0x62, 0x49,
	0x66, 0x7e, 0x9e, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x9c, 0xaf, 0xa4, 0xc5, 0x25, 0x84,
	0x66, 0x52, 0x41, 0x4e, 0x25, 0xc8, 0x9c, 0x12, 0x10, 0x0f, 0x66, 0x0e, 0x98, 0x63, 0x14, 0xce,
	0xc5, 0xe2, 0x58, 0x5a, 0x92, 0x21, 0xe4, 0xcf, 0xc5, 0x8b, 0xa2, 0x47, 0x48, 0x4e, 0x0f, 0xe6,
	0x32, 0x3d, 0x6c, 0xce, 0x92, 0x92, 0xc1, 0x29, 0x5f, 0x90, 0x53, 0xa9, 0xc4, 0x90, 0xc4, 0x06,
	0x96, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xfe, 0xfc, 0x26, 0xed, 0x00, 0x00, 0x00,
}
