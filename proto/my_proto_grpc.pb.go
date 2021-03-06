// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package t3

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

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerServiceClient interface {
	// Simple rpc
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	AddCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error)
	UpdateName(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error)
	UpdateNumber(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error)
	DeleteCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error)
	GetNumberRebelds(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) AddCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) UpdateName(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) UpdateNumber(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) DeleteCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerServiceClient) GetNumberRebelds(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaBroker, error) {
	out := new(RespuestaBroker)
	err := c.cc.Invoke(ctx, "/t3.BrokerService/GetNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
// All implementations must embed UnimplementedBrokerServiceServer
// for forward compatibility
type BrokerServiceServer interface {
	// Simple rpc
	SayHello(context.Context, *Message) (*Message, error)
	AddCity(context.Context, *Comando) (*RespuestaBroker, error)
	UpdateName(context.Context, *Comando) (*RespuestaBroker, error)
	UpdateNumber(context.Context, *Comando) (*RespuestaBroker, error)
	DeleteCity(context.Context, *Comando) (*RespuestaBroker, error)
	GetNumberRebelds(context.Context, *Comando) (*RespuestaBroker, error)
	mustEmbedUnimplementedBrokerServiceServer()
}

// UnimplementedBrokerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (UnimplementedBrokerServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBrokerServiceServer) AddCity(context.Context, *Comando) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedBrokerServiceServer) UpdateName(context.Context, *Comando) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedBrokerServiceServer) UpdateNumber(context.Context, *Comando) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedBrokerServiceServer) DeleteCity(context.Context, *Comando) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedBrokerServiceServer) GetNumberRebelds(context.Context, *Comando) (*RespuestaBroker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebelds not implemented")
}
func (UnimplementedBrokerServiceServer) mustEmbedUnimplementedBrokerServiceServer() {}

// UnsafeBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServiceServer will
// result in compilation errors.
type UnsafeBrokerServiceServer interface {
	mustEmbedUnimplementedBrokerServiceServer()
}

func RegisterBrokerServiceServer(s grpc.ServiceRegistrar, srv BrokerServiceServer) {
	s.RegisterService(&BrokerService_ServiceDesc, srv)
}

func _BrokerService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).AddCity(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).UpdateName(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).UpdateNumber(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).DeleteCity(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerService_GetNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).GetNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.BrokerService/GetNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).GetNumberRebelds(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerService_ServiceDesc is the grpc.ServiceDesc for BrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "t3.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BrokerService_SayHello_Handler,
		},
		{
			MethodName: "AddCity",
			Handler:    _BrokerService_AddCity_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _BrokerService_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _BrokerService_UpdateNumber_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _BrokerService_DeleteCity_Handler,
		},
		{
			MethodName: "GetNumberRebelds",
			Handler:    _BrokerService_GetNumberRebelds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "my_proto.proto",
}

// FulcrumServiceClient is the client API for FulcrumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulcrumServiceClient interface {
	// Simple rpc
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	AddCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error)
	UpdateName(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error)
	UpdateNumber(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error)
	DeleteCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error)
	GetNumberRebelds(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error)
	// Message contendra el nombre del planeta
	GetClockVector(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RespuestaReplica, error)
}

type fulcrumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrumServiceClient(cc grpc.ClientConnInterface) FulcrumServiceClient {
	return &fulcrumServiceClient{cc}
}

func (c *fulcrumServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) AddCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) UpdateName(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/UpdateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) UpdateNumber(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/UpdateNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) DeleteCity(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) GetNumberRebelds(ctx context.Context, in *Comando, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/GetNumberRebelds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServiceClient) GetClockVector(ctx context.Context, in *Message, opts ...grpc.CallOption) (*RespuestaReplica, error) {
	out := new(RespuestaReplica)
	err := c.cc.Invoke(ctx, "/t3.FulcrumService/GetClockVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulcrumServiceServer is the server API for FulcrumService service.
// All implementations must embed UnimplementedFulcrumServiceServer
// for forward compatibility
type FulcrumServiceServer interface {
	// Simple rpc
	SayHello(context.Context, *Message) (*Message, error)
	AddCity(context.Context, *Comando) (*RespuestaReplica, error)
	UpdateName(context.Context, *Comando) (*RespuestaReplica, error)
	UpdateNumber(context.Context, *Comando) (*RespuestaReplica, error)
	DeleteCity(context.Context, *Comando) (*RespuestaReplica, error)
	GetNumberRebelds(context.Context, *Comando) (*RespuestaReplica, error)
	// Message contendra el nombre del planeta
	GetClockVector(context.Context, *Message) (*RespuestaReplica, error)
	mustEmbedUnimplementedFulcrumServiceServer()
}

// UnimplementedFulcrumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrumServiceServer struct {
}

func (UnimplementedFulcrumServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedFulcrumServiceServer) AddCity(context.Context, *Comando) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedFulcrumServiceServer) UpdateName(context.Context, *Comando) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateName not implemented")
}
func (UnimplementedFulcrumServiceServer) UpdateNumber(context.Context, *Comando) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNumber not implemented")
}
func (UnimplementedFulcrumServiceServer) DeleteCity(context.Context, *Comando) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedFulcrumServiceServer) GetNumberRebelds(context.Context, *Comando) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumberRebelds not implemented")
}
func (UnimplementedFulcrumServiceServer) GetClockVector(context.Context, *Message) (*RespuestaReplica, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClockVector not implemented")
}
func (UnimplementedFulcrumServiceServer) mustEmbedUnimplementedFulcrumServiceServer() {}

// UnsafeFulcrumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulcrumServiceServer will
// result in compilation errors.
type UnsafeFulcrumServiceServer interface {
	mustEmbedUnimplementedFulcrumServiceServer()
}

func RegisterFulcrumServiceServer(s grpc.ServiceRegistrar, srv FulcrumServiceServer) {
	s.RegisterService(&FulcrumService_ServiceDesc, srv)
}

func _FulcrumService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).AddCity(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_UpdateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).UpdateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/UpdateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).UpdateName(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_UpdateNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).UpdateNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/UpdateNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).UpdateNumber(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).DeleteCity(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_GetNumberRebelds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comando)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).GetNumberRebelds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/GetNumberRebelds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).GetNumberRebelds(ctx, req.(*Comando))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumService_GetClockVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServiceServer).GetClockVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t3.FulcrumService/GetClockVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServiceServer).GetClockVector(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// FulcrumService_ServiceDesc is the grpc.ServiceDesc for FulcrumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FulcrumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "t3.FulcrumService",
	HandlerType: (*FulcrumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _FulcrumService_SayHello_Handler,
		},
		{
			MethodName: "AddCity",
			Handler:    _FulcrumService_AddCity_Handler,
		},
		{
			MethodName: "UpdateName",
			Handler:    _FulcrumService_UpdateName_Handler,
		},
		{
			MethodName: "UpdateNumber",
			Handler:    _FulcrumService_UpdateNumber_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _FulcrumService_DeleteCity_Handler,
		},
		{
			MethodName: "GetNumberRebelds",
			Handler:    _FulcrumService_GetNumberRebelds_Handler,
		},
		{
			MethodName: "GetClockVector",
			Handler:    _FulcrumService_GetClockVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "my_proto.proto",
}
