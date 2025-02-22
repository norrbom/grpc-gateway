// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: server/server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProxyPing_SayPing_FullMethodName = "/server.ProxyPing/SayPing"
)

// ProxyPingClient is the client API for ProxyPing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyPingClient interface {
	SayPing(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessageReply, error)
}

type proxyPingClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyPingClient(cc grpc.ClientConnInterface) ProxyPingClient {
	return &proxyPingClient{cc}
}

func (c *proxyPingClient) SayPing(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingMessageReply)
	err := c.cc.Invoke(ctx, ProxyPing_SayPing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyPingServer is the server API for ProxyPing service.
// All implementations must embed UnimplementedProxyPingServer
// for forward compatibility.
type ProxyPingServer interface {
	SayPing(context.Context, *PingMessage) (*PingMessageReply, error)
	mustEmbedUnimplementedProxyPingServer()
}

// UnimplementedProxyPingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProxyPingServer struct{}

func (UnimplementedProxyPingServer) SayPing(context.Context, *PingMessage) (*PingMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayPing not implemented")
}
func (UnimplementedProxyPingServer) mustEmbedUnimplementedProxyPingServer() {}
func (UnimplementedProxyPingServer) testEmbeddedByValue()                   {}

// UnsafeProxyPingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyPingServer will
// result in compilation errors.
type UnsafeProxyPingServer interface {
	mustEmbedUnimplementedProxyPingServer()
}

func RegisterProxyPingServer(s grpc.ServiceRegistrar, srv ProxyPingServer) {
	// If the following call pancis, it indicates UnimplementedProxyPingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProxyPing_ServiceDesc, srv)
}

func _ProxyPing_SayPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyPingServer).SayPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyPing_SayPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyPingServer).SayPing(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ProxyPing_ServiceDesc is the grpc.ServiceDesc for ProxyPing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyPing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.ProxyPing",
	HandlerType: (*ProxyPingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayPing",
			Handler:    _ProxyPing_SayPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/server.proto",
}

const (
	BackendPing_SayPing_FullMethodName = "/server.BackendPing/SayPing"
)

// BackendPingClient is the client API for BackendPing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackendPingClient interface {
	SayPing(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessageReply, error)
}

type backendPingClient struct {
	cc grpc.ClientConnInterface
}

func NewBackendPingClient(cc grpc.ClientConnInterface) BackendPingClient {
	return &backendPingClient{cc}
}

func (c *backendPingClient) SayPing(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessageReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingMessageReply)
	err := c.cc.Invoke(ctx, BackendPing_SayPing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendPingServer is the server API for BackendPing service.
// All implementations must embed UnimplementedBackendPingServer
// for forward compatibility.
type BackendPingServer interface {
	SayPing(context.Context, *PingMessage) (*PingMessageReply, error)
	mustEmbedUnimplementedBackendPingServer()
}

// UnimplementedBackendPingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBackendPingServer struct{}

func (UnimplementedBackendPingServer) SayPing(context.Context, *PingMessage) (*PingMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayPing not implemented")
}
func (UnimplementedBackendPingServer) mustEmbedUnimplementedBackendPingServer() {}
func (UnimplementedBackendPingServer) testEmbeddedByValue()                     {}

// UnsafeBackendPingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackendPingServer will
// result in compilation errors.
type UnsafeBackendPingServer interface {
	mustEmbedUnimplementedBackendPingServer()
}

func RegisterBackendPingServer(s grpc.ServiceRegistrar, srv BackendPingServer) {
	// If the following call pancis, it indicates UnimplementedBackendPingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BackendPing_ServiceDesc, srv)
}

func _BackendPing_SayPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendPingServer).SayPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackendPing_SayPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendPingServer).SayPing(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// BackendPing_ServiceDesc is the grpc.ServiceDesc for BackendPing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackendPing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.BackendPing",
	HandlerType: (*BackendPingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayPing",
			Handler:    _BackendPing_SayPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/server.proto",
}
