// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: services/brokerage/api/src/market.proto

package api

import (
	context "context"
	api "github.com/mrNobody95/Gate/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MarketServiceClient is the client API for MarketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketServiceClient interface {
	Set(ctx context.Context, in *Market, opts ...grpc.CallOption) (*Market, error)
	Get(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*Market, error)
	List(ctx context.Context, in *MarketListRequest, opts ...grpc.CallOption) (*Markets, error)
	ChangeStatus(ctx context.Context, in *api.StatusChangeRequest, opts ...grpc.CallOption) (*api.Status, error)
}

type marketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketServiceClient(cc grpc.ClientConnInterface) MarketServiceClient {
	return &marketServiceClient{cc}
}

func (c *marketServiceClient) Set(ctx context.Context, in *Market, opts ...grpc.CallOption) (*Market, error) {
	out := new(Market)
	err := c.cc.Invoke(ctx, "/brokerageApi.MarketService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketServiceClient) Get(ctx context.Context, in *MarketRequest, opts ...grpc.CallOption) (*Market, error) {
	out := new(Market)
	err := c.cc.Invoke(ctx, "/brokerageApi.MarketService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketServiceClient) List(ctx context.Context, in *MarketListRequest, opts ...grpc.CallOption) (*Markets, error) {
	out := new(Markets)
	err := c.cc.Invoke(ctx, "/brokerageApi.MarketService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketServiceClient) ChangeStatus(ctx context.Context, in *api.StatusChangeRequest, opts ...grpc.CallOption) (*api.Status, error) {
	out := new(api.Status)
	err := c.cc.Invoke(ctx, "/brokerageApi.MarketService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServiceServer is the server API for MarketService service.
// All implementations should embed UnimplementedMarketServiceServer
// for forward compatibility
type MarketServiceServer interface {
	Set(context.Context, *Market) (*Market, error)
	Get(context.Context, *MarketRequest) (*Market, error)
	List(context.Context, *MarketListRequest) (*Markets, error)
	ChangeStatus(context.Context, *api.StatusChangeRequest) (*api.Status, error)
}

// UnimplementedMarketServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMarketServiceServer struct {
}

func (UnimplementedMarketServiceServer) Set(context.Context, *Market) (*Market, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedMarketServiceServer) Get(context.Context, *MarketRequest) (*Market, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMarketServiceServer) List(context.Context, *MarketListRequest) (*Markets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMarketServiceServer) ChangeStatus(context.Context, *api.StatusChangeRequest) (*api.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}

// UnsafeMarketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServiceServer will
// result in compilation errors.
type UnsafeMarketServiceServer interface {
	mustEmbedUnimplementedMarketServiceServer()
}

func RegisterMarketServiceServer(s grpc.ServiceRegistrar, srv MarketServiceServer) {
	s.RegisterService(&MarketService_ServiceDesc, srv)
}

func _MarketService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Market)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.MarketService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).Set(ctx, req.(*Market))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.MarketService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).Get(ctx, req.(*MarketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.MarketService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).List(ctx, req.(*MarketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.StatusChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.MarketService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServiceServer).ChangeStatus(ctx, req.(*api.StatusChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketService_ServiceDesc is the grpc.ServiceDesc for MarketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brokerageApi.MarketService",
	HandlerType: (*MarketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _MarketService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MarketService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MarketService_List_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _MarketService_ChangeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/brokerage/api/src/market.proto",
}
