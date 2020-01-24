// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	connection "github.com/networkservicemesh/api/pkg/api/connection"
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

type NetworkServiceRequest struct {
	Connection           *connection.Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	MechanismPreferences []*connection.Mechanism `protobuf:"bytes,2,rep,name=mechanism_preferences,json=mechanismPreferences,proto3" json:"mechanism_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkServiceRequest) Reset()         { *m = NetworkServiceRequest{} }
func (m *NetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceRequest) ProtoMessage()    {}
func (*NetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_361e8247d5a9945c, []int{0}
}

func (m *NetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceRequest.Unmarshal(m, b)
}
func (m *NetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceRequest.Marshal(b, m, deterministic)
}
func (m *NetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceRequest.Merge(m, src)
}
func (m *NetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceRequest.Size(m)
}
func (m *NetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceRequest proto.InternalMessageInfo

func (m *NetworkServiceRequest) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceRequest) GetMechanismPreferences() []*connection.Mechanism {
	if m != nil {
		return m.MechanismPreferences
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "networkservice.NetworkServiceRequest")
}

func init() { proto.RegisterFile("networkservice.proto", fileDescriptor_361e8247d5a9945c) }

var fileDescriptor_361e8247d5a9945c = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xa9, 0xa2, 0x42, 0x06, 0x43, 0xc2, 0x3a, 0x4a, 0x9f, 0x86, 0x20, 0xec, 0x29, 0x85,
	0x0a, 0xfa, 0xee, 0x10, 0x44, 0x50, 0xa4, 0xbe, 0x09, 0x22, 0x6d, 0xb8, 0x6b, 0x83, 0x4d, 0x6e,
	0x4c, 0x52, 0x65, 0x9f, 0xc3, 0x47, 0xbf, 0xac, 0xcc, 0xec, 0x4f, 0x3a, 0xb6, 0x97, 0x72, 0x39,
	0xf7, 0xdc, 0x5f, 0x4f, 0xee, 0x25, 0x23, 0x05, 0xee, 0x1b, 0xcd, 0x87, 0x05, 0xf3, 0x25, 0x38,
	0x30, 0x6d, 0xd0, 0x21, 0x1d, 0xf6, 0xd5, 0xf4, 0xad, 0x16, 0xae, 0xe9, 0x2a, 0xc6, 0x51, 0x66,
	0xfd, 0x96, 0x04, 0xdb, 0xec, 0x93, 0x38, 0x2a, 0x67, 0xb0, 0xd5, 0x6d, 0xa9, 0x20, 0x2b, 0xb5,
	0x58, 0x0a, 0x0a, 0xb8, 0x13, 0xa8, 0x82, 0xd2, 0xff, 0x2e, 0x4d, 0xb4, 0x5b, 0x68, 0xb0, 0x19,
	0x48, 0xed, 0x16, 0xfe, 0xeb, 0x3b, 0x17, 0xbf, 0x11, 0x89, 0x9f, 0x3c, 0xfd, 0xc5, 0xd3, 0x0b,
	0xf8, 0xec, 0xc0, 0x3a, 0x7a, 0x4d, 0xc8, 0x96, 0x93, 0x44, 0x93, 0x68, 0x3a, 0xc8, 0xc7, 0x2c,
	0x40, 0xcf, 0x36, 0x65, 0x11, 0x38, 0xe9, 0x03, 0x89, 0x25, 0xf0, 0xa6, 0x54, 0xc2, 0xca, 0x77,
	0x6d, 0x60, 0x0e, 0x06, 0x14, 0x07, 0x9b, 0x1c, 0x4d, 0x8e, 0xa7, 0x83, 0x3c, 0x0e, 0x11, 0x8f,
	0x6b, 0x63, 0x31, 0xda, 0xcc, 0x3c, 0x6f, 0x47, 0xf2, 0x9f, 0x88, 0x0c, 0xfb, 0xe9, 0xe8, 0x3d,
	0x39, 0x5b, 0x27, 0xbc, 0x64, 0x3b, 0xbb, 0xdd, 0xfb, 0x90, 0xf4, 0x40, 0x68, 0x7a, 0x43, 0x4e,
	0x66, 0x2d, 0x5a, 0xa0, 0x07, 0x0c, 0xe9, 0x98, 0xd5, 0x88, 0x75, 0xbb, 0xba, 0x59, 0xd5, 0xcd,
	0xd9, 0xdd, 0x72, 0x73, 0xb7, 0xe7, 0xaf, 0x3b, 0xe7, 0xab, 0x4e, 0xff, 0x1d, 0x57, 0x7f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0x2d, 0x6f, 0xf7, 0xed, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error) {
	out := new(connection.Connection)
	err := c.cc.Invoke(ctx, "/networkservice.NetworkService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/networkservice.NetworkService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Request(context.Context, *NetworkServiceRequest) (*connection.Connection, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) Request(ctx context.Context, req *NetworkServiceRequest) (*connection.Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedNetworkServiceServer) Close(ctx context.Context, req *connection.Connection) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkservice.NetworkService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Request(ctx, req.(*NetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connection.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/networkservice.NetworkService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Close(ctx, req.(*connection.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "networkservice.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _NetworkService_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _NetworkService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkservice.proto",
}
