// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package timeservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TimeServiceClient is the client API for TimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeServiceClient interface {
	// Send current wall clock time.
	SendClock(ctx context.Context, in *ClockRequest, opts ...grpc.CallOption) (*ClockReply, error)
}

type timeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeServiceClient(cc grpc.ClientConnInterface) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) SendClock(ctx context.Context, in *ClockRequest, opts ...grpc.CallOption) (*ClockReply, error) {
	out := new(ClockReply)
	err := c.cc.Invoke(ctx, "/timeservice.TimeService/SendClock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeServiceServer is the server API for TimeService service.
// All implementations must embed UnimplementedTimeServiceServer
// for forward compatibility
type TimeServiceServer interface {
	// Send current wall clock time.
	SendClock(context.Context, *ClockRequest) (*ClockReply, error)
	mustEmbedUnimplementedTimeServiceServer()
}

// UnimplementedTimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeServiceServer struct {
}

func (UnimplementedTimeServiceServer) SendClock(context.Context, *ClockRequest) (*ClockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClock not implemented")
}
func (UnimplementedTimeServiceServer) mustEmbedUnimplementedTimeServiceServer() {}

// UnsafeTimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeServiceServer will
// result in compilation errors.
type UnsafeTimeServiceServer interface {
	mustEmbedUnimplementedTimeServiceServer()
}

func RegisterTimeServiceServer(s grpc.ServiceRegistrar, srv TimeServiceServer) {
	s.RegisterService(&_TimeService_serviceDesc, srv)
}

func _TimeService_SendClock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServiceServer).SendClock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timeservice.TimeService/SendClock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServiceServer).SendClock(ctx, req.(*ClockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timeservice.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClock",
			Handler:    _TimeService_SendClock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/timeservice/timeservice.proto",
}
