// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: hiveos.proto

package pb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

// HiveosServiceClient is the client API for HiveosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HiveosServiceClient interface {
}

type hiveosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHiveosServiceClient(cc grpc.ClientConnInterface) HiveosServiceClient {
	return &hiveosServiceClient{cc}
}

// HiveosServiceServer is the server API for HiveosService service.
// All implementations must embed UnimplementedHiveosServiceServer
// for forward compatibility.
type HiveosServiceServer interface {
	mustEmbedUnimplementedHiveosServiceServer()
}

// UnimplementedHiveosServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHiveosServiceServer struct{}

func (UnimplementedHiveosServiceServer) mustEmbedUnimplementedHiveosServiceServer() {}
func (UnimplementedHiveosServiceServer) testEmbeddedByValue()                       {}

// UnsafeHiveosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HiveosServiceServer will
// result in compilation errors.
type UnsafeHiveosServiceServer interface {
	mustEmbedUnimplementedHiveosServiceServer()
}

func RegisterHiveosServiceServer(s grpc.ServiceRegistrar, srv HiveosServiceServer) {
	// If the following call pancis, it indicates UnimplementedHiveosServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HiveosService_ServiceDesc, srv)
}

// HiveosService_ServiceDesc is the grpc.ServiceDesc for HiveosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HiveosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.HiveosService",
	HandlerType: (*HiveosServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "hiveos.proto",
}
