// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: protobufs/apiserver.proto

package protobufs

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

// PrismcloudApiserverClient is the client API for PrismcloudApiserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrismcloudApiserverClient interface {
	Version(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ApiVersion, error)
	GetNamespaces(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Namespaces, error)
	CreateNamespace(ctx context.Context, in *NamespaceCreateRequest, opts ...grpc.CallOption) (*Namespace, error)
	DeleteNamespace(ctx context.Context, in *NamespaceDeleteRequest, opts ...grpc.CallOption) (*Void, error)
	CreatePod(ctx context.Context, in *PodCreateRequest, opts ...grpc.CallOption) (*Void, error)
	DeletePod(ctx context.Context, in *PodDeleteRequest, opts ...grpc.CallOption) (*Void, error)
	GetPod(ctx context.Context, in *ServiceFilter, opts ...grpc.CallOption) (*Pod, error)
	CreateLBIngress(ctx context.Context, in *LBIngressCreateRequest, opts ...grpc.CallOption) (*LBIngress, error)
	DeleteLBIngress(ctx context.Context, in *LBIngressDeleteRequest, opts ...grpc.CallOption) (*Void, error)
	GetLBIngress(ctx context.Context, in *ServiceFilter, opts ...grpc.CallOption) (*LBIngress, error)
}

type prismcloudApiserverClient struct {
	cc grpc.ClientConnInterface
}

func NewPrismcloudApiserverClient(cc grpc.ClientConnInterface) PrismcloudApiserverClient {
	return &prismcloudApiserverClient{cc}
}

func (c *prismcloudApiserverClient) Version(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ApiVersion, error) {
	out := new(ApiVersion)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) GetNamespaces(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Namespaces, error) {
	out := new(Namespaces)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/GetNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) CreateNamespace(ctx context.Context, in *NamespaceCreateRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) DeleteNamespace(ctx context.Context, in *NamespaceDeleteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) CreatePod(ctx context.Context, in *PodCreateRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/CreatePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) DeletePod(ctx context.Context, in *PodDeleteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) GetPod(ctx context.Context, in *ServiceFilter, opts ...grpc.CallOption) (*Pod, error) {
	out := new(Pod)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/GetPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) CreateLBIngress(ctx context.Context, in *LBIngressCreateRequest, opts ...grpc.CallOption) (*LBIngress, error) {
	out := new(LBIngress)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/CreateLBIngress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) DeleteLBIngress(ctx context.Context, in *LBIngressDeleteRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/DeleteLBIngress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prismcloudApiserverClient) GetLBIngress(ctx context.Context, in *ServiceFilter, opts ...grpc.CallOption) (*LBIngress, error) {
	out := new(LBIngress)
	err := c.cc.Invoke(ctx, "/prismcloud.PrismcloudApiserver/GetLBIngress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrismcloudApiserverServer is the server API for PrismcloudApiserver service.
// All implementations must embed UnimplementedPrismcloudApiserverServer
// for forward compatibility
type PrismcloudApiserverServer interface {
	Version(context.Context, *Void) (*ApiVersion, error)
	GetNamespaces(context.Context, *Void) (*Namespaces, error)
	CreateNamespace(context.Context, *NamespaceCreateRequest) (*Namespace, error)
	DeleteNamespace(context.Context, *NamespaceDeleteRequest) (*Void, error)
	CreatePod(context.Context, *PodCreateRequest) (*Void, error)
	DeletePod(context.Context, *PodDeleteRequest) (*Void, error)
	GetPod(context.Context, *ServiceFilter) (*Pod, error)
	CreateLBIngress(context.Context, *LBIngressCreateRequest) (*LBIngress, error)
	DeleteLBIngress(context.Context, *LBIngressDeleteRequest) (*Void, error)
	GetLBIngress(context.Context, *ServiceFilter) (*LBIngress, error)
	mustEmbedUnimplementedPrismcloudApiserverServer()
}

// UnimplementedPrismcloudApiserverServer must be embedded to have forward compatible implementations.
type UnimplementedPrismcloudApiserverServer struct {
}

func (UnimplementedPrismcloudApiserverServer) Version(context.Context, *Void) (*ApiVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedPrismcloudApiserverServer) GetNamespaces(context.Context, *Void) (*Namespaces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaces not implemented")
}
func (UnimplementedPrismcloudApiserverServer) CreateNamespace(context.Context, *NamespaceCreateRequest) (*Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedPrismcloudApiserverServer) DeleteNamespace(context.Context, *NamespaceDeleteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedPrismcloudApiserverServer) CreatePod(context.Context, *PodCreateRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePod not implemented")
}
func (UnimplementedPrismcloudApiserverServer) DeletePod(context.Context, *PodDeleteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (UnimplementedPrismcloudApiserverServer) GetPod(context.Context, *ServiceFilter) (*Pod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (UnimplementedPrismcloudApiserverServer) CreateLBIngress(context.Context, *LBIngressCreateRequest) (*LBIngress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLBIngress not implemented")
}
func (UnimplementedPrismcloudApiserverServer) DeleteLBIngress(context.Context, *LBIngressDeleteRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLBIngress not implemented")
}
func (UnimplementedPrismcloudApiserverServer) GetLBIngress(context.Context, *ServiceFilter) (*LBIngress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLBIngress not implemented")
}
func (UnimplementedPrismcloudApiserverServer) mustEmbedUnimplementedPrismcloudApiserverServer() {}

// UnsafePrismcloudApiserverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrismcloudApiserverServer will
// result in compilation errors.
type UnsafePrismcloudApiserverServer interface {
	mustEmbedUnimplementedPrismcloudApiserverServer()
}

func RegisterPrismcloudApiserverServer(s grpc.ServiceRegistrar, srv PrismcloudApiserverServer) {
	s.RegisterService(&PrismcloudApiserver_ServiceDesc, srv)
}

func _PrismcloudApiserver_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).Version(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).GetNamespaces(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).CreateNamespace(ctx, req.(*NamespaceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).DeleteNamespace(ctx, req.(*NamespaceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_CreatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).CreatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/CreatePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).CreatePod(ctx, req.(*PodCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).DeletePod(ctx, req.(*PodDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/GetPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).GetPod(ctx, req.(*ServiceFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_CreateLBIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LBIngressCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).CreateLBIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/CreateLBIngress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).CreateLBIngress(ctx, req.(*LBIngressCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_DeleteLBIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LBIngressDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).DeleteLBIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/DeleteLBIngress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).DeleteLBIngress(ctx, req.(*LBIngressDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrismcloudApiserver_GetLBIngress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrismcloudApiserverServer).GetLBIngress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prismcloud.PrismcloudApiserver/GetLBIngress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrismcloudApiserverServer).GetLBIngress(ctx, req.(*ServiceFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// PrismcloudApiserver_ServiceDesc is the grpc.ServiceDesc for PrismcloudApiserver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrismcloudApiserver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prismcloud.PrismcloudApiserver",
	HandlerType: (*PrismcloudApiserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _PrismcloudApiserver_Version_Handler,
		},
		{
			MethodName: "GetNamespaces",
			Handler:    _PrismcloudApiserver_GetNamespaces_Handler,
		},
		{
			MethodName: "CreateNamespace",
			Handler:    _PrismcloudApiserver_CreateNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _PrismcloudApiserver_DeleteNamespace_Handler,
		},
		{
			MethodName: "CreatePod",
			Handler:    _PrismcloudApiserver_CreatePod_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _PrismcloudApiserver_DeletePod_Handler,
		},
		{
			MethodName: "GetPod",
			Handler:    _PrismcloudApiserver_GetPod_Handler,
		},
		{
			MethodName: "CreateLBIngress",
			Handler:    _PrismcloudApiserver_CreateLBIngress_Handler,
		},
		{
			MethodName: "DeleteLBIngress",
			Handler:    _PrismcloudApiserver_DeleteLBIngress_Handler,
		},
		{
			MethodName: "GetLBIngress",
			Handler:    _PrismcloudApiserver_GetLBIngress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/apiserver.proto",
}
