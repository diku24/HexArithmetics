// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: arithmetic_svc.proto

package pb

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

// ArtihmeticServiceClient is the client API for ArtihmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtihmeticServiceClient interface {
	GetAddtion(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetSubstraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
	GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error)
}

type artihmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtihmeticServiceClient(cc grpc.ClientConnInterface) ArtihmeticServiceClient {
	return &artihmeticServiceClient{cc}
}

func (c *artihmeticServiceClient) GetAddtion(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArtihmeticService/GetAddtion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artihmeticServiceClient) GetSubstraction(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArtihmeticService/GetSubstraction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artihmeticServiceClient) GetMultiplication(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArtihmeticService/GetMultiplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artihmeticServiceClient) GetDivision(ctx context.Context, in *OperationParameters, opts ...grpc.CallOption) (*Answer, error) {
	out := new(Answer)
	err := c.cc.Invoke(ctx, "/pb.ArtihmeticService/GetDivision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtihmeticServiceServer is the server API for ArtihmeticService service.
// All implementations should embed UnimplementedArtihmeticServiceServer
// for forward compatibility
type ArtihmeticServiceServer interface {
	GetAddtion(context.Context, *OperationParameters) (*Answer, error)
	GetSubstraction(context.Context, *OperationParameters) (*Answer, error)
	GetMultiplication(context.Context, *OperationParameters) (*Answer, error)
	GetDivision(context.Context, *OperationParameters) (*Answer, error)
}

// UnimplementedArtihmeticServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArtihmeticServiceServer struct {
}

func (UnimplementedArtihmeticServiceServer) GetAddtion(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddtion not implemented")
}
func (UnimplementedArtihmeticServiceServer) GetSubstraction(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubstraction not implemented")
}
func (UnimplementedArtihmeticServiceServer) GetMultiplication(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiplication not implemented")
}
func (UnimplementedArtihmeticServiceServer) GetDivision(context.Context, *OperationParameters) (*Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivision not implemented")
}

// UnsafeArtihmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtihmeticServiceServer will
// result in compilation errors.
type UnsafeArtihmeticServiceServer interface {
	mustEmbedUnimplementedArtihmeticServiceServer()
}

func RegisterArtihmeticServiceServer(s grpc.ServiceRegistrar, srv ArtihmeticServiceServer) {
	s.RegisterService(&ArtihmeticService_ServiceDesc, srv)
}

func _ArtihmeticService_GetAddtion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtihmeticServiceServer).GetAddtion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArtihmeticService/GetAddtion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtihmeticServiceServer).GetAddtion(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtihmeticService_GetSubstraction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtihmeticServiceServer).GetSubstraction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArtihmeticService/GetSubstraction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtihmeticServiceServer).GetSubstraction(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtihmeticService_GetMultiplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtihmeticServiceServer).GetMultiplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArtihmeticService/GetMultiplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtihmeticServiceServer).GetMultiplication(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtihmeticService_GetDivision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtihmeticServiceServer).GetDivision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ArtihmeticService/GetDivision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtihmeticServiceServer).GetDivision(ctx, req.(*OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtihmeticService_ServiceDesc is the grpc.ServiceDesc for ArtihmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtihmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArtihmeticService",
	HandlerType: (*ArtihmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddtion",
			Handler:    _ArtihmeticService_GetAddtion_Handler,
		},
		{
			MethodName: "GetSubstraction",
			Handler:    _ArtihmeticService_GetSubstraction_Handler,
		},
		{
			MethodName: "GetMultiplication",
			Handler:    _ArtihmeticService_GetMultiplication_Handler,
		},
		{
			MethodName: "GetDivision",
			Handler:    _ArtihmeticService_GetDivision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arithmetic_svc.proto",
}
