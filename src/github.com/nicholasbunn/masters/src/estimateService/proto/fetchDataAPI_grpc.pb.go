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

// FetchDataClient is the client API for FetchData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FetchDataClient interface {
	FetchDataService(ctx context.Context, in *FetchDataRequestMessage, opts ...grpc.CallOption) (*FetchDataResponseMessage, error)
}

type fetchDataClient struct {
	cc grpc.ClientConnInterface
}

func NewFetchDataClient(cc grpc.ClientConnInterface) FetchDataClient {
	return &fetchDataClient{cc}
}

func (c *fetchDataClient) FetchDataService(ctx context.Context, in *FetchDataRequestMessage, opts ...grpc.CallOption) (*FetchDataResponseMessage, error) {
	out := new(FetchDataResponseMessage)
	err := c.cc.Invoke(ctx, "/fetchData.FetchData/FetchDataService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FetchDataServer is the server API for FetchData service.
// All implementations must embed UnimplementedFetchDataServer
// for forward compatibility
type FetchDataServer interface {
	FetchDataService(context.Context, *FetchDataRequestMessage) (*FetchDataResponseMessage, error)
	mustEmbedUnimplementedFetchDataServer()
}

// UnimplementedFetchDataServer must be embedded to have forward compatible implementations.
type UnimplementedFetchDataServer struct {
}

func (UnimplementedFetchDataServer) FetchDataService(context.Context, *FetchDataRequestMessage) (*FetchDataResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchDataService not implemented")
}
func (UnimplementedFetchDataServer) mustEmbedUnimplementedFetchDataServer() {}

// UnsafeFetchDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FetchDataServer will
// result in compilation errors.
type UnsafeFetchDataServer interface {
	mustEmbedUnimplementedFetchDataServer()
}

func RegisterFetchDataServer(s grpc.ServiceRegistrar, srv FetchDataServer) {
	s.RegisterService(&FetchData_ServiceDesc, srv)
}

func _FetchData_FetchDataService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchDataRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FetchDataServer).FetchDataService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetchData.FetchData/FetchDataService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FetchDataServer).FetchDataService(ctx, req.(*FetchDataRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FetchData_ServiceDesc is the grpc.ServiceDesc for FetchData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FetchData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fetchData.FetchData",
	HandlerType: (*FetchDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchDataService",
			Handler:    _FetchData_FetchDataService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fetchDataService/proto/fetchDataAPI.proto",
}
