// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: parcels.proto

package parcels

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ParcelsClient is the client API for Parcels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParcelsClient interface {
	GetParcels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ParcelsResponse, error)
	GetParcelByTrackingCode(ctx context.Context, in *ParcelReq, opts ...grpc.CallOption) (*ParcelsResponse, error)
}

type parcelsClient struct {
	cc grpc.ClientConnInterface
}

func NewParcelsClient(cc grpc.ClientConnInterface) ParcelsClient {
	return &parcelsClient{cc}
}

func (c *parcelsClient) GetParcels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ParcelsResponse, error) {
	out := new(ParcelsResponse)
	err := c.cc.Invoke(ctx, "/grpc.parcels.Parcels/GetParcels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelsClient) GetParcelByTrackingCode(ctx context.Context, in *ParcelReq, opts ...grpc.CallOption) (*ParcelsResponse, error) {
	out := new(ParcelsResponse)
	err := c.cc.Invoke(ctx, "/grpc.parcels.Parcels/GetParcelByTrackingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParcelsServer is the server API for Parcels service.
// All implementations must embed UnimplementedParcelsServer
// for forward compatibility
type ParcelsServer interface {
	GetParcels(context.Context, *emptypb.Empty) (*ParcelsResponse, error)
	GetParcelByTrackingCode(context.Context, *ParcelReq) (*ParcelsResponse, error)
	mustEmbedUnimplementedParcelsServer()
}

// UnimplementedParcelsServer must be embedded to have forward compatible implementations.
type UnimplementedParcelsServer struct {
}

func (UnimplementedParcelsServer) GetParcels(context.Context, *emptypb.Empty) (*ParcelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParcels not implemented")
}
func (UnimplementedParcelsServer) GetParcelByTrackingCode(context.Context, *ParcelReq) (*ParcelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParcelByTrackingCode not implemented")
}
func (UnimplementedParcelsServer) mustEmbedUnimplementedParcelsServer() {}

// UnsafeParcelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParcelsServer will
// result in compilation errors.
type UnsafeParcelsServer interface {
	mustEmbedUnimplementedParcelsServer()
}

func RegisterParcelsServer(s grpc.ServiceRegistrar, srv ParcelsServer) {
	s.RegisterService(&Parcels_ServiceDesc, srv)
}

func _Parcels_GetParcels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsServer).GetParcels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.parcels.Parcels/GetParcels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsServer).GetParcels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Parcels_GetParcelByTrackingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParcelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsServer).GetParcelByTrackingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.parcels.Parcels/GetParcelByTrackingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsServer).GetParcelByTrackingCode(ctx, req.(*ParcelReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Parcels_ServiceDesc is the grpc.ServiceDesc for Parcels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Parcels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.parcels.Parcels",
	HandlerType: (*ParcelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParcels",
			Handler:    _Parcels_GetParcels_Handler,
		},
		{
			MethodName: "GetParcelByTrackingCode",
			Handler:    _Parcels_GetParcelByTrackingCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parcels.proto",
}