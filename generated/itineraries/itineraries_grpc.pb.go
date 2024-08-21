// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: itineraries.proto

package itineraries

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ItinerariesService_CreateItinerary_FullMethodName = "/itineraries_service.ItinerariesService/CreateItinerary"
	ItinerariesService_UpdateItinerary_FullMethodName = "/itineraries_service.ItinerariesService/UpdateItinerary"
	ItinerariesService_DeleteItinerary_FullMethodName = "/itineraries_service.ItinerariesService/DeleteItinerary"
	ItinerariesService_ListItineraries_FullMethodName = "/itineraries_service.ItinerariesService/ListItineraries"
	ItinerariesService_GetItinerary_FullMethodName    = "/itineraries_service.ItinerariesService/GetItinerary"
	ItinerariesService_LeaveComment_FullMethodName    = "/itineraries_service.ItinerariesService/LeaveComment"
)

// ItinerariesServiceClient is the client API for ItinerariesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItinerariesServiceClient interface {
	CreateItinerary(ctx context.Context, in *CreateItineraryRequest, opts ...grpc.CallOption) (*CreateItineraryResponse, error)
	UpdateItinerary(ctx context.Context, in *UpdateItineraryRequest, opts ...grpc.CallOption) (*UpdateItineraryResponse, error)
	DeleteItinerary(ctx context.Context, in *DeleteItineraryRequest, opts ...grpc.CallOption) (*DeleteItineraryResponse, error)
	ListItineraries(ctx context.Context, in *ListItinerariesRequest, opts ...grpc.CallOption) (*ListItinerariesResponse, error)
	GetItinerary(ctx context.Context, in *GetItineraryRequest, opts ...grpc.CallOption) (*GetItineraryResponse, error)
	LeaveComment(ctx context.Context, in *LeaveCommentRequest, opts ...grpc.CallOption) (*LeaveCommentResponse, error)
}

type itinerariesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItinerariesServiceClient(cc grpc.ClientConnInterface) ItinerariesServiceClient {
	return &itinerariesServiceClient{cc}
}

func (c *itinerariesServiceClient) CreateItinerary(ctx context.Context, in *CreateItineraryRequest, opts ...grpc.CallOption) (*CreateItineraryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateItineraryResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_CreateItinerary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesServiceClient) UpdateItinerary(ctx context.Context, in *UpdateItineraryRequest, opts ...grpc.CallOption) (*UpdateItineraryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateItineraryResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_UpdateItinerary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesServiceClient) DeleteItinerary(ctx context.Context, in *DeleteItineraryRequest, opts ...grpc.CallOption) (*DeleteItineraryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteItineraryResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_DeleteItinerary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesServiceClient) ListItineraries(ctx context.Context, in *ListItinerariesRequest, opts ...grpc.CallOption) (*ListItinerariesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListItinerariesResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_ListItineraries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesServiceClient) GetItinerary(ctx context.Context, in *GetItineraryRequest, opts ...grpc.CallOption) (*GetItineraryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItineraryResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_GetItinerary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itinerariesServiceClient) LeaveComment(ctx context.Context, in *LeaveCommentRequest, opts ...grpc.CallOption) (*LeaveCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveCommentResponse)
	err := c.cc.Invoke(ctx, ItinerariesService_LeaveComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItinerariesServiceServer is the server API for ItinerariesService service.
// All implementations must embed UnimplementedItinerariesServiceServer
// for forward compatibility
type ItinerariesServiceServer interface {
	CreateItinerary(context.Context, *CreateItineraryRequest) (*CreateItineraryResponse, error)
	UpdateItinerary(context.Context, *UpdateItineraryRequest) (*UpdateItineraryResponse, error)
	DeleteItinerary(context.Context, *DeleteItineraryRequest) (*DeleteItineraryResponse, error)
	ListItineraries(context.Context, *ListItinerariesRequest) (*ListItinerariesResponse, error)
	GetItinerary(context.Context, *GetItineraryRequest) (*GetItineraryResponse, error)
	LeaveComment(context.Context, *LeaveCommentRequest) (*LeaveCommentResponse, error)
	mustEmbedUnimplementedItinerariesServiceServer()
}

// UnimplementedItinerariesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedItinerariesServiceServer struct {
}

func (UnimplementedItinerariesServiceServer) CreateItinerary(context.Context, *CreateItineraryRequest) (*CreateItineraryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItinerary not implemented")
}
func (UnimplementedItinerariesServiceServer) UpdateItinerary(context.Context, *UpdateItineraryRequest) (*UpdateItineraryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItinerary not implemented")
}
func (UnimplementedItinerariesServiceServer) DeleteItinerary(context.Context, *DeleteItineraryRequest) (*DeleteItineraryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItinerary not implemented")
}
func (UnimplementedItinerariesServiceServer) ListItineraries(context.Context, *ListItinerariesRequest) (*ListItinerariesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItineraries not implemented")
}
func (UnimplementedItinerariesServiceServer) GetItinerary(context.Context, *GetItineraryRequest) (*GetItineraryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItinerary not implemented")
}
func (UnimplementedItinerariesServiceServer) LeaveComment(context.Context, *LeaveCommentRequest) (*LeaveCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveComment not implemented")
}
func (UnimplementedItinerariesServiceServer) mustEmbedUnimplementedItinerariesServiceServer() {}

// UnsafeItinerariesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItinerariesServiceServer will
// result in compilation errors.
type UnsafeItinerariesServiceServer interface {
	mustEmbedUnimplementedItinerariesServiceServer()
}

func RegisterItinerariesServiceServer(s grpc.ServiceRegistrar, srv ItinerariesServiceServer) {
	s.RegisterService(&ItinerariesService_ServiceDesc, srv)
}

func _ItinerariesService_CreateItinerary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItineraryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).CreateItinerary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_CreateItinerary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).CreateItinerary(ctx, req.(*CreateItineraryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItinerariesService_UpdateItinerary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItineraryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).UpdateItinerary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_UpdateItinerary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).UpdateItinerary(ctx, req.(*UpdateItineraryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItinerariesService_DeleteItinerary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItineraryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).DeleteItinerary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_DeleteItinerary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).DeleteItinerary(ctx, req.(*DeleteItineraryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItinerariesService_ListItineraries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItinerariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).ListItineraries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_ListItineraries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).ListItineraries(ctx, req.(*ListItinerariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItinerariesService_GetItinerary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItineraryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).GetItinerary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_GetItinerary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).GetItinerary(ctx, req.(*GetItineraryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItinerariesService_LeaveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItinerariesServiceServer).LeaveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItinerariesService_LeaveComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItinerariesServiceServer).LeaveComment(ctx, req.(*LeaveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItinerariesService_ServiceDesc is the grpc.ServiceDesc for ItinerariesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItinerariesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itineraries_service.ItinerariesService",
	HandlerType: (*ItinerariesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItinerary",
			Handler:    _ItinerariesService_CreateItinerary_Handler,
		},
		{
			MethodName: "UpdateItinerary",
			Handler:    _ItinerariesService_UpdateItinerary_Handler,
		},
		{
			MethodName: "DeleteItinerary",
			Handler:    _ItinerariesService_DeleteItinerary_Handler,
		},
		{
			MethodName: "ListItineraries",
			Handler:    _ItinerariesService_ListItineraries_Handler,
		},
		{
			MethodName: "GetItinerary",
			Handler:    _ItinerariesService_GetItinerary_Handler,
		},
		{
			MethodName: "LeaveComment",
			Handler:    _ItinerariesService_LeaveComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "itineraries.proto",
}
