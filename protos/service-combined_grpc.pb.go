// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.1
// source: protos/service-combined.proto

package protos

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
	CombinedCompute_RootMeanSquare_FullMethodName        = "/protos.CombinedCompute/RootMeanSquare"
	CombinedCompute_GeometricMean_FullMethodName         = "/protos.CombinedCompute/GeometricMean"
	CombinedCompute_BodyMassIndex_FullMethodName         = "/protos.CombinedCompute/BodyMassIndex"
	CombinedCompute_PowerLevelDiff_FullMethodName        = "/protos.CombinedCompute/PowerLevelDiff"
	CombinedCompute_PercentageValueChange_FullMethodName = "/protos.CombinedCompute/PercentageValueChange"
)

// CombinedComputeClient is the client API for CombinedCompute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CombinedComputeClient interface {
	RootMeanSquare(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	GeometricMean(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	BodyMassIndex(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	PowerLevelDiff(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
	PercentageValueChange(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
}

type combinedComputeClient struct {
	cc grpc.ClientConnInterface
}

func NewCombinedComputeClient(cc grpc.ClientConnInterface) CombinedComputeClient {
	return &combinedComputeClient{cc}
}

func (c *combinedComputeClient) RootMeanSquare(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, CombinedCompute_RootMeanSquare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *combinedComputeClient) GeometricMean(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, CombinedCompute_GeometricMean_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *combinedComputeClient) BodyMassIndex(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, CombinedCompute_BodyMassIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *combinedComputeClient) PowerLevelDiff(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, CombinedCompute_PowerLevelDiff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *combinedComputeClient) PercentageValueChange(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, CombinedCompute_PercentageValueChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CombinedComputeServer is the server API for CombinedCompute service.
// All implementations must embed UnimplementedCombinedComputeServer
// for forward compatibility.
type CombinedComputeServer interface {
	RootMeanSquare(context.Context, *OperationRequest) (*OperationResponse, error)
	GeometricMean(context.Context, *OperationRequest) (*OperationResponse, error)
	BodyMassIndex(context.Context, *OperationRequest) (*OperationResponse, error)
	PowerLevelDiff(context.Context, *OperationRequest) (*OperationResponse, error)
	PercentageValueChange(context.Context, *OperationRequest) (*OperationResponse, error)
	mustEmbedUnimplementedCombinedComputeServer()
}

// UnimplementedCombinedComputeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCombinedComputeServer struct{}

func (UnimplementedCombinedComputeServer) RootMeanSquare(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RootMeanSquare not implemented")
}
func (UnimplementedCombinedComputeServer) GeometricMean(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeometricMean not implemented")
}
func (UnimplementedCombinedComputeServer) BodyMassIndex(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BodyMassIndex not implemented")
}
func (UnimplementedCombinedComputeServer) PowerLevelDiff(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerLevelDiff not implemented")
}
func (UnimplementedCombinedComputeServer) PercentageValueChange(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PercentageValueChange not implemented")
}
func (UnimplementedCombinedComputeServer) mustEmbedUnimplementedCombinedComputeServer() {}
func (UnimplementedCombinedComputeServer) testEmbeddedByValue()                         {}

// UnsafeCombinedComputeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CombinedComputeServer will
// result in compilation errors.
type UnsafeCombinedComputeServer interface {
	mustEmbedUnimplementedCombinedComputeServer()
}

func RegisterCombinedComputeServer(s grpc.ServiceRegistrar, srv CombinedComputeServer) {
	// If the following call pancis, it indicates UnimplementedCombinedComputeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CombinedCompute_ServiceDesc, srv)
}

func _CombinedCompute_RootMeanSquare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedComputeServer).RootMeanSquare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CombinedCompute_RootMeanSquare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedComputeServer).RootMeanSquare(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CombinedCompute_GeometricMean_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedComputeServer).GeometricMean(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CombinedCompute_GeometricMean_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedComputeServer).GeometricMean(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CombinedCompute_BodyMassIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedComputeServer).BodyMassIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CombinedCompute_BodyMassIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedComputeServer).BodyMassIndex(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CombinedCompute_PowerLevelDiff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedComputeServer).PowerLevelDiff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CombinedCompute_PowerLevelDiff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedComputeServer).PowerLevelDiff(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CombinedCompute_PercentageValueChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedComputeServer).PercentageValueChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CombinedCompute_PercentageValueChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedComputeServer).PercentageValueChange(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CombinedCompute_ServiceDesc is the grpc.ServiceDesc for CombinedCompute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CombinedCompute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CombinedCompute",
	HandlerType: (*CombinedComputeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RootMeanSquare",
			Handler:    _CombinedCompute_RootMeanSquare_Handler,
		},
		{
			MethodName: "GeometricMean",
			Handler:    _CombinedCompute_GeometricMean_Handler,
		},
		{
			MethodName: "BodyMassIndex",
			Handler:    _CombinedCompute_BodyMassIndex_Handler,
		},
		{
			MethodName: "PowerLevelDiff",
			Handler:    _CombinedCompute_PowerLevelDiff_Handler,
		},
		{
			MethodName: "PercentageValueChange",
			Handler:    _CombinedCompute_PercentageValueChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service-combined.proto",
}
