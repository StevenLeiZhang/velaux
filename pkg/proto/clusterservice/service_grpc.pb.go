// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clusterservice

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

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	PutCluster(ctx context.Context, in *PutClusterRequest, opts ...grpc.CallOption) (*PutClusterResponse, error)
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	DelCluster(ctx context.Context, in *DelClusterRequest, opts ...grpc.CallOption) (*DelClusterResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) PutCluster(ctx context.Context, in *PutClusterRequest, opts ...grpc.CallOption) (*PutClusterResponse, error) {
	out := new(PutClusterResponse)
	err := c.cc.Invoke(ctx, "/vela.api.service.clusterservice.ClusterService/PutCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*GetClusterResponse, error) {
	out := new(GetClusterResponse)
	err := c.cc.Invoke(ctx, "/vela.api.service.clusterservice.ClusterService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/vela.api.service.clusterservice.ClusterService/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) DelCluster(ctx context.Context, in *DelClusterRequest, opts ...grpc.CallOption) (*DelClusterResponse, error) {
	out := new(DelClusterResponse)
	err := c.cc.Invoke(ctx, "/vela.api.service.clusterservice.ClusterService/DelCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	PutCluster(context.Context, *PutClusterRequest) (*PutClusterResponse, error)
	GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error)
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	DelCluster(context.Context, *DelClusterRequest) (*DelClusterResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) PutCluster(context.Context, *PutClusterRequest) (*PutClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCluster not implemented")
}
func (UnimplementedClusterServiceServer) GetCluster(context.Context, *GetClusterRequest) (*GetClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedClusterServiceServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedClusterServiceServer) DelCluster(context.Context, *DelClusterRequest) (*DelClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelCluster not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_PutCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).PutCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vela.api.service.clusterservice.ClusterService/PutCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).PutCluster(ctx, req.(*PutClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vela.api.service.clusterservice.ClusterService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vela.api.service.clusterservice.ClusterService/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_DelCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).DelCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vela.api.service.clusterservice.ClusterService/DelCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).DelCluster(ctx, req.(*DelClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vela.api.service.clusterservice.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutCluster",
			Handler:    _ClusterService_PutCluster_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _ClusterService_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _ClusterService_ListClusters_Handler,
		},
		{
			MethodName: "DelCluster",
			Handler:    _ClusterService_DelCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/clusterservice/service.proto",
}