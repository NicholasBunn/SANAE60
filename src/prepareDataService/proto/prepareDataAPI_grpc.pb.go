// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package prepareData

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

// PrepareDataClient is the client API for PrepareData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrepareDataClient interface {
	PrepareEstimateDataService(ctx context.Context, in *PrepareRequestMessage, opts ...grpc.CallOption) (*PrepareResponseMessage, error)
}

type prepareDataClient struct {
	cc grpc.ClientConnInterface
}

func NewPrepareDataClient(cc grpc.ClientConnInterface) PrepareDataClient {
	return &prepareDataClient{cc}
}

func (c *prepareDataClient) PrepareEstimateDataService(ctx context.Context, in *PrepareRequestMessage, opts ...grpc.CallOption) (*PrepareResponseMessage, error) {
	out := new(PrepareResponseMessage)
	err := c.cc.Invoke(ctx, "/prepareData.PrepareData/PrepareEstimateDataService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrepareDataServer is the server API for PrepareData service.
// All implementations must embed UnimplementedPrepareDataServer
// for forward compatibility
type PrepareDataServer interface {
	PrepareEstimateDataService(context.Context, *PrepareRequestMessage) (*PrepareResponseMessage, error)
	mustEmbedUnimplementedPrepareDataServer()
}

// UnimplementedPrepareDataServer must be embedded to have forward compatible implementations.
type UnimplementedPrepareDataServer struct {
}

func (UnimplementedPrepareDataServer) PrepareEstimateDataService(context.Context, *PrepareRequestMessage) (*PrepareResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareEstimateDataService not implemented")
}
func (UnimplementedPrepareDataServer) mustEmbedUnimplementedPrepareDataServer() {}

// UnsafePrepareDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrepareDataServer will
// result in compilation errors.
type UnsafePrepareDataServer interface {
	mustEmbedUnimplementedPrepareDataServer()
}

func RegisterPrepareDataServer(s grpc.ServiceRegistrar, srv PrepareDataServer) {
	s.RegisterService(&PrepareData_ServiceDesc, srv)
}

func _PrepareData_PrepareEstimateDataService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrepareDataServer).PrepareEstimateDataService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prepareData.PrepareData/PrepareEstimateDataService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrepareDataServer).PrepareEstimateDataService(ctx, req.(*PrepareRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// PrepareData_ServiceDesc is the grpc.ServiceDesc for PrepareData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrepareData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prepareData.PrepareData",
	HandlerType: (*PrepareDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareEstimateDataService",
			Handler:    _PrepareData_PrepareEstimateDataService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prepareDataService/proto/prepareDataAPI.proto",
}
