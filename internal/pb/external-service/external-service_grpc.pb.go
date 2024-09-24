// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.2
// source: external-service.proto

package external_service

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
	FooService_GetFoo_FullMethodName = "/external_service.FooService/GetFoo"
)

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// FooService provides operations on Foo entities
type FooServiceClient interface {
	// GetFoo retrieves a Foo by its ID
	GetFoo(ctx context.Context, in *GetFooRequest, opts ...grpc.CallOption) (*Foo, error)
}

type fooServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFooServiceClient(cc grpc.ClientConnInterface) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) GetFoo(ctx context.Context, in *GetFooRequest, opts ...grpc.CallOption) (*Foo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Foo)
	err := c.cc.Invoke(ctx, FooService_GetFoo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
// All implementations must embed UnimplementedFooServiceServer
// for forward compatibility.
//
// FooService provides operations on Foo entities
type FooServiceServer interface {
	// GetFoo retrieves a Foo by its ID
	GetFoo(context.Context, *GetFooRequest) (*Foo, error)
	mustEmbedUnimplementedFooServiceServer()
}

// UnimplementedFooServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFooServiceServer struct{}

func (UnimplementedFooServiceServer) GetFoo(context.Context, *GetFooRequest) (*Foo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoo not implemented")
}
func (UnimplementedFooServiceServer) mustEmbedUnimplementedFooServiceServer() {}
func (UnimplementedFooServiceServer) testEmbeddedByValue()                    {}

// UnsafeFooServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FooServiceServer will
// result in compilation errors.
type UnsafeFooServiceServer interface {
	mustEmbedUnimplementedFooServiceServer()
}

func RegisterFooServiceServer(s grpc.ServiceRegistrar, srv FooServiceServer) {
	// If the following call pancis, it indicates UnimplementedFooServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FooService_ServiceDesc, srv)
}

func _FooService_GetFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).GetFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FooService_GetFoo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).GetFoo(ctx, req.(*GetFooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FooService_ServiceDesc is the grpc.ServiceDesc for FooService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FooService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "external_service.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFoo",
			Handler:    _FooService_GetFoo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external-service.proto",
}
