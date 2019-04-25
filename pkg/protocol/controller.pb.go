// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	controller.proto
	metrics.proto
	image.proto
	util.proto
	auth.proto

It has these top-level messages:
	ChangeContainerStateRequest
	ChangeContainerStateReply
	GetContainersLogsRequest
	GetContainersLogsReply
	MetricsRequest
	MetricsResponse
	DeleteImageEvent
	DeleteImageLayerEvent
	Timestamp
	GenerateTokenRequest
	GenerateTokenReply
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChangeContainerStateRequest_State int32

const (
	ChangeContainerStateRequest_START ChangeContainerStateRequest_State = 0
	ChangeContainerStateRequest_STOP  ChangeContainerStateRequest_State = 1
)

var ChangeContainerStateRequest_State_name = map[int32]string{
	0: "START",
	1: "STOP",
}
var ChangeContainerStateRequest_State_value = map[string]int32{
	"START": 0,
	"STOP":  1,
}

func (x ChangeContainerStateRequest_State) String() string {
	return proto.EnumName(ChangeContainerStateRequest_State_name, int32(x))
}
func (ChangeContainerStateRequest_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type GetContainersLogsReply_Output int32

const (
	GetContainersLogsReply_STDOUT GetContainersLogsReply_Output = 0
	GetContainersLogsReply_STDERR GetContainersLogsReply_Output = 1
)

var GetContainersLogsReply_Output_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
}
var GetContainersLogsReply_Output_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
}

func (x GetContainersLogsReply_Output) String() string {
	return proto.EnumName(GetContainersLogsReply_Output_name, int32(x))
}
func (GetContainersLogsReply_Output) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type ChangeContainerStateRequest struct {
	Id             string                            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RequestedState ChangeContainerStateRequest_State `protobuf:"varint,2,opt,name=requestedState,enum=protocol.ChangeContainerStateRequest_State" json:"requestedState,omitempty"`
}

func (m *ChangeContainerStateRequest) Reset()                    { *m = ChangeContainerStateRequest{} }
func (m *ChangeContainerStateRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangeContainerStateRequest) ProtoMessage()               {}
func (*ChangeContainerStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChangeContainerStateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChangeContainerStateRequest) GetRequestedState() ChangeContainerStateRequest_State {
	if m != nil {
		return m.RequestedState
	}
	return ChangeContainerStateRequest_START
}

type ChangeContainerStateReply struct {
}

func (m *ChangeContainerStateReply) Reset()                    { *m = ChangeContainerStateReply{} }
func (m *ChangeContainerStateReply) String() string            { return proto.CompactTextString(m) }
func (*ChangeContainerStateReply) ProtoMessage()               {}
func (*ChangeContainerStateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetContainersLogsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetContainersLogsRequest) Reset()                    { *m = GetContainersLogsRequest{} }
func (m *GetContainersLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetContainersLogsRequest) ProtoMessage()               {}
func (*GetContainersLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetContainersLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetContainersLogsReply struct {
	Output  GetContainersLogsReply_Output `protobuf:"varint,1,opt,name=output,enum=protocol.GetContainersLogsReply_Output" json:"output,omitempty"`
	Time    *Timestamp                    `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	TaskId  string                        `protobuf:"bytes,3,opt,name=taskId" json:"taskId,omitempty"`
	Message string                        `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *GetContainersLogsReply) Reset()                    { *m = GetContainersLogsReply{} }
func (m *GetContainersLogsReply) String() string            { return proto.CompactTextString(m) }
func (*GetContainersLogsReply) ProtoMessage()               {}
func (*GetContainersLogsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetContainersLogsReply) GetOutput() GetContainersLogsReply_Output {
	if m != nil {
		return m.Output
	}
	return GetContainersLogsReply_STDOUT
}

func (m *GetContainersLogsReply) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *GetContainersLogsReply) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *GetContainersLogsReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ChangeContainerStateRequest)(nil), "protocol.ChangeContainerStateRequest")
	proto.RegisterType((*ChangeContainerStateReply)(nil), "protocol.ChangeContainerStateReply")
	proto.RegisterType((*GetContainersLogsRequest)(nil), "protocol.GetContainersLogsRequest")
	proto.RegisterType((*GetContainersLogsReply)(nil), "protocol.GetContainersLogsReply")
	proto.RegisterEnum("protocol.ChangeContainerStateRequest_State", ChangeContainerStateRequest_State_name, ChangeContainerStateRequest_State_value)
	proto.RegisterEnum("protocol.GetContainersLogsReply_Output", GetContainersLogsReply_Output_name, GetContainersLogsReply_Output_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Controller service

type ControllerClient interface {
	ChangeContainerState(ctx context.Context, in *ChangeContainerStateRequest, opts ...grpc.CallOption) (*ChangeContainerStateReply, error)
	GetContainerLogs(ctx context.Context, in *GetContainersLogsRequest, opts ...grpc.CallOption) (Controller_GetContainerLogsClient, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) ChangeContainerState(ctx context.Context, in *ChangeContainerStateRequest, opts ...grpc.CallOption) (*ChangeContainerStateReply, error) {
	out := new(ChangeContainerStateReply)
	err := grpc.Invoke(ctx, "/protocol.Controller/ChangeContainerState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) GetContainerLogs(ctx context.Context, in *GetContainersLogsRequest, opts ...grpc.CallOption) (Controller_GetContainerLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Controller_serviceDesc.Streams[0], c.cc, "/protocol.Controller/GetContainerLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerGetContainerLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Controller_GetContainerLogsClient interface {
	Recv() (*GetContainersLogsReply, error)
	grpc.ClientStream
}

type controllerGetContainerLogsClient struct {
	grpc.ClientStream
}

func (x *controllerGetContainerLogsClient) Recv() (*GetContainersLogsReply, error) {
	m := new(GetContainersLogsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Controller service

type ControllerServer interface {
	ChangeContainerState(context.Context, *ChangeContainerStateRequest) (*ChangeContainerStateReply, error)
	GetContainerLogs(*GetContainersLogsRequest, Controller_GetContainerLogsServer) error
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_ChangeContainerState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeContainerStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).ChangeContainerState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Controller/ChangeContainerState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).ChangeContainerState(ctx, req.(*ChangeContainerStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_GetContainerLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetContainersLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServer).GetContainerLogs(m, &controllerGetContainerLogsServer{stream})
}

type Controller_GetContainerLogsServer interface {
	Send(*GetContainersLogsReply) error
	grpc.ServerStream
}

type controllerGetContainerLogsServer struct {
	grpc.ServerStream
}

func (x *controllerGetContainerLogsServer) Send(m *GetContainersLogsReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeContainerState",
			Handler:    _Controller_ChangeContainerState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetContainerLogs",
			Handler:       _Controller_GetContainerLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "controller.proto",
}

func init() { proto.RegisterFile("controller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xc7, 0xbb, 0xfd, 0xb5, 0xf9, 0xb5, 0x23, 0x84, 0x30, 0x4a, 0x89, 0xad, 0x87, 0xb0, 0x22,
	0x2d, 0x0a, 0x41, 0xe2, 0x03, 0x88, 0x54, 0x11, 0x41, 0xa8, 0x24, 0xf1, 0xe6, 0x25, 0x36, 0x4b,
	0x0d, 0x26, 0xd9, 0x98, 0xdd, 0x1c, 0xfa, 0x36, 0x3e, 0x92, 0x27, 0x9f, 0x47, 0xba, 0xdb, 0x3f,
	0x22, 0xb1, 0xf5, 0xb4, 0x33, 0x3b, 0xdf, 0xf9, 0xee, 0xec, 0x67, 0xc0, 0x9a, 0xf2, 0x5c, 0x96,
	0x3c, 0x4d, 0x59, 0xe9, 0x16, 0x25, 0x97, 0x1c, 0x3b, 0xea, 0x98, 0xf2, 0xb4, 0x0f, 0x95, 0x4c,
	0x52, 0x7d, 0x4b, 0xdf, 0x09, 0x0c, 0xc6, 0x2f, 0x51, 0x3e, 0x63, 0x63, 0x9e, 0xcb, 0x28, 0xc9,
	0x59, 0x19, 0xc8, 0x48, 0x32, 0x9f, 0xbd, 0x55, 0x4c, 0x48, 0x34, 0xa1, 0x99, 0xc4, 0x36, 0x71,
	0xc8, 0xa8, 0xeb, 0x37, 0x93, 0x18, 0x03, 0x30, 0x4b, 0x5d, 0x62, 0xb1, 0x12, 0xda, 0x4d, 0x87,
	0x8c, 0x4c, 0xef, 0xcc, 0x5d, 0xd9, 0xbb, 0x5b, 0xec, 0x5c, 0x9d, 0xfc, 0xb0, 0xa0, 0x47, 0xd0,
	0x56, 0x01, 0x76, 0xa1, 0x1d, 0x84, 0x57, 0x7e, 0x68, 0x35, 0xb0, 0x03, 0xad, 0x20, 0x9c, 0x3c,
	0x58, 0x84, 0x0e, 0xe0, 0xb0, 0xde, 0xb2, 0x48, 0xe7, 0xf4, 0x14, 0xec, 0x5b, 0x26, 0xd7, 0x15,
	0x71, 0xcf, 0x67, 0xe2, 0x97, 0xd9, 0xe9, 0x27, 0x81, 0x5e, 0x8d, 0xb8, 0x48, 0xe7, 0x78, 0x09,
	0x06, 0xaf, 0x64, 0x51, 0x49, 0x25, 0x37, 0xbd, 0xe1, 0xe6, 0x3b, 0xf5, 0x1d, 0xee, 0x44, 0xc9,
	0xfd, 0x65, 0x1b, 0x0e, 0xa1, 0x25, 0x93, 0x4c, 0xd3, 0xd8, 0xf3, 0xf6, 0x37, 0xed, 0x61, 0x92,
	0x31, 0x21, 0xa3, 0xac, 0xf0, 0x95, 0x00, 0x7b, 0x60, 0xc8, 0x48, 0xbc, 0xde, 0xc5, 0xf6, 0x3f,
	0x35, 0xd8, 0x32, 0x43, 0x1b, 0xfe, 0x67, 0x4c, 0x88, 0x68, 0xc6, 0xec, 0x96, 0x2a, 0xac, 0x52,
	0xea, 0x80, 0xa1, 0x1f, 0x43, 0x00, 0x23, 0x08, 0xaf, 0x27, 0x8f, 0x0b, 0x3e, 0x3a, 0xbe, 0xf1,
	0x7d, 0x8b, 0x78, 0x1f, 0x04, 0x60, 0xbc, 0xde, 0x37, 0xc6, 0x70, 0x50, 0x07, 0x0c, 0x4f, 0xfe,
	0xb4, 0xa3, 0xfe, 0xf1, 0x2e, 0xd9, 0x82, 0x7b, 0x03, 0x9f, 0xc0, 0xfa, 0x8e, 0x66, 0x41, 0x06,
	0xe9, 0x56, 0x6c, 0xda, 0xde, 0xd9, 0x85, 0x96, 0x36, 0xce, 0xc9, 0xb3, 0xa1, 0x44, 0x17, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x2d, 0x13, 0x21, 0xc8, 0x02, 0x00, 0x00,
}
