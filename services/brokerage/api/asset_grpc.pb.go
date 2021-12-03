// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: services/brokerage/api/src/asset.proto

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

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetServiceClient interface {
	Set(ctx context.Context, in *api.Asset, opts ...grpc.CallOption) (*api.Void, error)
	GetByName(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*api.Asset, error)
	List(ctx context.Context, in *GetAssetListRequest, opts ...grpc.CallOption) (*api.Assets, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) Set(ctx context.Context, in *api.Asset, opts ...grpc.CallOption) (*api.Void, error) {
	out := new(api.Void)
	err := c.cc.Invoke(ctx, "/brokerageApi.AssetService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) GetByName(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*api.Asset, error) {
	out := new(api.Asset)
	err := c.cc.Invoke(ctx, "/brokerageApi.AssetService/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) List(ctx context.Context, in *GetAssetListRequest, opts ...grpc.CallOption) (*api.Assets, error) {
	out := new(api.Assets)
	err := c.cc.Invoke(ctx, "/brokerageApi.AssetService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
// All implementations should embed UnimplementedAssetServiceServer
// for forward compatibility
type AssetServiceServer interface {
	Set(context.Context, *api.Asset) (*api.Void, error)
	GetByName(context.Context, *GetAssetRequest) (*api.Asset, error)
	List(context.Context, *GetAssetListRequest) (*api.Assets, error)
}

// UnimplementedAssetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (UnimplementedAssetServiceServer) Set(context.Context, *api.Asset) (*api.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedAssetServiceServer) GetByName(context.Context, *GetAssetRequest) (*api.Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedAssetServiceServer) List(context.Context, *GetAssetListRequest) (*api.Assets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServiceServer will
// result in compilation errors.
type UnsafeAssetServiceServer interface {
	mustEmbedUnimplementedAssetServiceServer()
}

func RegisterAssetServiceServer(s grpc.ServiceRegistrar, srv AssetServiceServer) {
	s.RegisterService(&AssetService_ServiceDesc, srv)
}

func _AssetService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.Asset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.AssetService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).Set(ctx, req.(*api.Asset))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.AssetService/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).GetByName(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerageApi.AssetService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).List(ctx, req.(*GetAssetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetService_ServiceDesc is the grpc.ServiceDesc for AssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "brokerageApi.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _AssetService_Set_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _AssetService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AssetService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/brokerage/api/src/asset.proto",
}
