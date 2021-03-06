// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BasicServiceClient is the client API for BasicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicServiceClient interface {
	GetTimeAndName(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerInfo, error)
}

type basicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicServiceClient(cc grpc.ClientConnInterface) BasicServiceClient {
	return &basicServiceClient{cc}
}

func (c *basicServiceClient) GetTimeAndName(ctx context.Context, in *ServerInfo, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := c.cc.Invoke(ctx, "/proto.BasicService/GetTimeAndName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServiceServer is the server API for BasicService service.
// All implementations must embed UnimplementedBasicServiceServer
// for forward compatibility
type BasicServiceServer interface {
	GetTimeAndName(context.Context, *ServerInfo) (*ServerInfo, error)
	mustEmbedUnimplementedBasicServiceServer()
}

// UnimplementedBasicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBasicServiceServer struct {
}

func (UnimplementedBasicServiceServer) GetTimeAndName(context.Context, *ServerInfo) (*ServerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeAndName not implemented")
}
func (UnimplementedBasicServiceServer) mustEmbedUnimplementedBasicServiceServer() {}

// UnsafeBasicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicServiceServer will
// result in compilation errors.
type UnsafeBasicServiceServer interface {
	mustEmbedUnimplementedBasicServiceServer()
}

func RegisterBasicServiceServer(s grpc.ServiceRegistrar, srv BasicServiceServer) {
	s.RegisterService(&BasicService_ServiceDesc, srv)
}

func _BasicService_GetTimeAndName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServiceServer).GetTimeAndName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BasicService/GetTimeAndName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServiceServer).GetTimeAndName(ctx, req.(*ServerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// BasicService_ServiceDesc is the grpc.ServiceDesc for BasicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BasicService",
	HandlerType: (*BasicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeAndName",
			Handler:    _BasicService_GetTimeAndName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
