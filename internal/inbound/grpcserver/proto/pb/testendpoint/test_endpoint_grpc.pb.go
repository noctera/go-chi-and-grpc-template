// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: test_endpoint.proto

package testendpoint

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
	TestEndpoint_GetIdsByName_FullMethodName = "/test_endpoint.TestEndpoint/GetIdsByName"
)

// TestEndpointClient is the client API for TestEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestEndpointClient interface {
	GetIdsByName(ctx context.Context, in *GetIdsByNameRequest, opts ...grpc.CallOption) (*GetIdsByNameResponse, error)
}

type testEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewTestEndpointClient(cc grpc.ClientConnInterface) TestEndpointClient {
	return &testEndpointClient{cc}
}

func (c *testEndpointClient) GetIdsByName(ctx context.Context, in *GetIdsByNameRequest, opts ...grpc.CallOption) (*GetIdsByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdsByNameResponse)
	err := c.cc.Invoke(ctx, TestEndpoint_GetIdsByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestEndpointServer is the server API for TestEndpoint service.
// All implementations must embed UnimplementedTestEndpointServer
// for forward compatibility.
type TestEndpointServer interface {
	GetIdsByName(context.Context, *GetIdsByNameRequest) (*GetIdsByNameResponse, error)
	mustEmbedUnimplementedTestEndpointServer()
}

// UnimplementedTestEndpointServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestEndpointServer struct{}

func (UnimplementedTestEndpointServer) GetIdsByName(context.Context, *GetIdsByNameRequest) (*GetIdsByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdsByName not implemented")
}
func (UnimplementedTestEndpointServer) mustEmbedUnimplementedTestEndpointServer() {}
func (UnimplementedTestEndpointServer) testEmbeddedByValue()                      {}

// UnsafeTestEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestEndpointServer will
// result in compilation errors.
type UnsafeTestEndpointServer interface {
	mustEmbedUnimplementedTestEndpointServer()
}

func RegisterTestEndpointServer(s grpc.ServiceRegistrar, srv TestEndpointServer) {
	// If the following call pancis, it indicates UnimplementedTestEndpointServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestEndpoint_ServiceDesc, srv)
}

func _TestEndpoint_GetIdsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestEndpointServer).GetIdsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestEndpoint_GetIdsByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestEndpointServer).GetIdsByName(ctx, req.(*GetIdsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestEndpoint_ServiceDesc is the grpc.ServiceDesc for TestEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test_endpoint.TestEndpoint",
	HandlerType: (*TestEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdsByName",
			Handler:    _TestEndpoint_GetIdsByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test_endpoint.proto",
}