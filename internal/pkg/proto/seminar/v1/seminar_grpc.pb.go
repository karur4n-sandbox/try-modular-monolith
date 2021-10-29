// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SeminarServiceClient is the client API for SeminarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeminarServiceClient interface {
	ListSeminars(ctx context.Context, in *ListSeminarsRequest, opts ...grpc.CallOption) (*ListSeminarsResponse, error)
}

type seminarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSeminarServiceClient(cc grpc.ClientConnInterface) SeminarServiceClient {
	return &seminarServiceClient{cc}
}

func (c *seminarServiceClient) ListSeminars(ctx context.Context, in *ListSeminarsRequest, opts ...grpc.CallOption) (*ListSeminarsResponse, error) {
	out := new(ListSeminarsResponse)
	err := c.cc.Invoke(ctx, "/try_modular_monolith.seminar.v1.SeminarService/ListSeminars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeminarServiceServer is the server API for SeminarService service.
// All implementations must embed UnimplementedSeminarServiceServer
// for forward compatibility
type SeminarServiceServer interface {
	ListSeminars(context.Context, *ListSeminarsRequest) (*ListSeminarsResponse, error)
	mustEmbedUnimplementedSeminarServiceServer()
}

// UnimplementedSeminarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSeminarServiceServer struct {
}

func (UnimplementedSeminarServiceServer) ListSeminars(context.Context, *ListSeminarsRequest) (*ListSeminarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSeminars not implemented")
}
func (UnimplementedSeminarServiceServer) mustEmbedUnimplementedSeminarServiceServer() {}

// UnsafeSeminarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeminarServiceServer will
// result in compilation errors.
type UnsafeSeminarServiceServer interface {
	mustEmbedUnimplementedSeminarServiceServer()
}

func RegisterSeminarServiceServer(s grpc.ServiceRegistrar, srv SeminarServiceServer) {
	s.RegisterService(&SeminarService_ServiceDesc, srv)
}

func _SeminarService_ListSeminars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSeminarsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeminarServiceServer).ListSeminars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/try_modular_monolith.seminar.v1.SeminarService/ListSeminars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeminarServiceServer).ListSeminars(ctx, req.(*ListSeminarsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SeminarService_ServiceDesc is the grpc.ServiceDesc for SeminarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeminarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "try_modular_monolith.seminar.v1.SeminarService",
	HandlerType: (*SeminarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSeminars",
			Handler:    _SeminarService_ListSeminars_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/seminar/v1/seminar.proto",
}