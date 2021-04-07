// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package powerEstimation

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

// PowerEstimationServicePackageClient is the client API for PowerEstimationServicePackage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PowerEstimationServicePackageClient interface {
	PowerEstimatorService(ctx context.Context, in *ServicePackageRequestMessage, opts ...grpc.CallOption) (*EstimateResponseMessage, error)
	PowerEvaluatorService(ctx context.Context, in *ServicePackageRequestMessage, opts ...grpc.CallOption) (*EvaluateResponseMessage, error)
}

type powerEstimationServicePackageClient struct {
	cc grpc.ClientConnInterface
}

func NewPowerEstimationServicePackageClient(cc grpc.ClientConnInterface) PowerEstimationServicePackageClient {
	return &powerEstimationServicePackageClient{cc}
}

func (c *powerEstimationServicePackageClient) PowerEstimatorService(ctx context.Context, in *ServicePackageRequestMessage, opts ...grpc.CallOption) (*EstimateResponseMessage, error) {
	out := new(EstimateResponseMessage)
	err := c.cc.Invoke(ctx, "/powerEstimation.PowerEstimationServicePackage/PowerEstimatorService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerEstimationServicePackageClient) PowerEvaluatorService(ctx context.Context, in *ServicePackageRequestMessage, opts ...grpc.CallOption) (*EvaluateResponseMessage, error) {
	out := new(EvaluateResponseMessage)
	err := c.cc.Invoke(ctx, "/powerEstimation.PowerEstimationServicePackage/PowerEvaluatorService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PowerEstimationServicePackageServer is the server API for PowerEstimationServicePackage service.
// All implementations must embed UnimplementedPowerEstimationServicePackageServer
// for forward compatibility
type PowerEstimationServicePackageServer interface {
	PowerEstimatorService(context.Context, *ServicePackageRequestMessage) (*EstimateResponseMessage, error)
	PowerEvaluatorService(context.Context, *ServicePackageRequestMessage) (*EvaluateResponseMessage, error)
	mustEmbedUnimplementedPowerEstimationServicePackageServer()
}

// UnimplementedPowerEstimationServicePackageServer must be embedded to have forward compatible implementations.
type UnimplementedPowerEstimationServicePackageServer struct {
}

func (UnimplementedPowerEstimationServicePackageServer) PowerEstimatorService(context.Context, *ServicePackageRequestMessage) (*EstimateResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerEstimatorService not implemented")
}
func (UnimplementedPowerEstimationServicePackageServer) PowerEvaluatorService(context.Context, *ServicePackageRequestMessage) (*EvaluateResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerEvaluatorService not implemented")
}
func (UnimplementedPowerEstimationServicePackageServer) mustEmbedUnimplementedPowerEstimationServicePackageServer() {
}

// UnsafePowerEstimationServicePackageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PowerEstimationServicePackageServer will
// result in compilation errors.
type UnsafePowerEstimationServicePackageServer interface {
	mustEmbedUnimplementedPowerEstimationServicePackageServer()
}

func RegisterPowerEstimationServicePackageServer(s grpc.ServiceRegistrar, srv PowerEstimationServicePackageServer) {
	s.RegisterService(&PowerEstimationServicePackage_ServiceDesc, srv)
}

func _PowerEstimationServicePackage_PowerEstimatorService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicePackageRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerEstimationServicePackageServer).PowerEstimatorService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerEstimation.PowerEstimationServicePackage/PowerEstimatorService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerEstimationServicePackageServer).PowerEstimatorService(ctx, req.(*ServicePackageRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerEstimationServicePackage_PowerEvaluatorService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServicePackageRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerEstimationServicePackageServer).PowerEvaluatorService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/powerEstimation.PowerEstimationServicePackage/PowerEvaluatorService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerEstimationServicePackageServer).PowerEvaluatorService(ctx, req.(*ServicePackageRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// PowerEstimationServicePackage_ServiceDesc is the grpc.ServiceDesc for PowerEstimationServicePackage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PowerEstimationServicePackage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "powerEstimation.PowerEstimationServicePackage",
	HandlerType: (*PowerEstimationServicePackageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PowerEstimatorService",
			Handler:    _PowerEstimationServicePackage_PowerEstimatorService_Handler,
		},
		{
			MethodName: "PowerEvaluatorService",
			Handler:    _PowerEstimationServicePackage_PowerEvaluatorService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "powerEstimationSP/proto/powerEstimationAPI.proto",
}