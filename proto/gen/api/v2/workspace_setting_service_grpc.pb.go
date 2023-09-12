// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v2/workspace_setting_service.proto

package apiv2

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

const (
	WorkspaceSettingService_GetWorkspaceSetting_FullMethodName    = "/slash.api.v2.WorkspaceSettingService/GetWorkspaceSetting"
	WorkspaceSettingService_UpdateWorkspaceSetting_FullMethodName = "/slash.api.v2.WorkspaceSettingService/UpdateWorkspaceSetting"
)

// WorkspaceSettingServiceClient is the client API for WorkspaceSettingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceSettingServiceClient interface {
	GetWorkspaceSetting(ctx context.Context, in *GetWorkspaceSettingRequest, opts ...grpc.CallOption) (*GetWorkspaceSettingResponse, error)
	UpdateWorkspaceSetting(ctx context.Context, in *UpdateWorkspaceSettingRequest, opts ...grpc.CallOption) (*UpdateWorkspaceSettingResponse, error)
}

type workspaceSettingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceSettingServiceClient(cc grpc.ClientConnInterface) WorkspaceSettingServiceClient {
	return &workspaceSettingServiceClient{cc}
}

func (c *workspaceSettingServiceClient) GetWorkspaceSetting(ctx context.Context, in *GetWorkspaceSettingRequest, opts ...grpc.CallOption) (*GetWorkspaceSettingResponse, error) {
	out := new(GetWorkspaceSettingResponse)
	err := c.cc.Invoke(ctx, WorkspaceSettingService_GetWorkspaceSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceSettingServiceClient) UpdateWorkspaceSetting(ctx context.Context, in *UpdateWorkspaceSettingRequest, opts ...grpc.CallOption) (*UpdateWorkspaceSettingResponse, error) {
	out := new(UpdateWorkspaceSettingResponse)
	err := c.cc.Invoke(ctx, WorkspaceSettingService_UpdateWorkspaceSetting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceSettingServiceServer is the server API for WorkspaceSettingService service.
// All implementations must embed UnimplementedWorkspaceSettingServiceServer
// for forward compatibility
type WorkspaceSettingServiceServer interface {
	GetWorkspaceSetting(context.Context, *GetWorkspaceSettingRequest) (*GetWorkspaceSettingResponse, error)
	UpdateWorkspaceSetting(context.Context, *UpdateWorkspaceSettingRequest) (*UpdateWorkspaceSettingResponse, error)
	mustEmbedUnimplementedWorkspaceSettingServiceServer()
}

// UnimplementedWorkspaceSettingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceSettingServiceServer struct {
}

func (UnimplementedWorkspaceSettingServiceServer) GetWorkspaceSetting(context.Context, *GetWorkspaceSettingRequest) (*GetWorkspaceSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceSetting not implemented")
}
func (UnimplementedWorkspaceSettingServiceServer) UpdateWorkspaceSetting(context.Context, *UpdateWorkspaceSettingRequest) (*UpdateWorkspaceSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspaceSetting not implemented")
}
func (UnimplementedWorkspaceSettingServiceServer) mustEmbedUnimplementedWorkspaceSettingServiceServer() {
}

// UnsafeWorkspaceSettingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceSettingServiceServer will
// result in compilation errors.
type UnsafeWorkspaceSettingServiceServer interface {
	mustEmbedUnimplementedWorkspaceSettingServiceServer()
}

func RegisterWorkspaceSettingServiceServer(s grpc.ServiceRegistrar, srv WorkspaceSettingServiceServer) {
	s.RegisterService(&WorkspaceSettingService_ServiceDesc, srv)
}

func _WorkspaceSettingService_GetWorkspaceSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceSettingServiceServer).GetWorkspaceSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceSettingService_GetWorkspaceSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceSettingServiceServer).GetWorkspaceSetting(ctx, req.(*GetWorkspaceSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceSettingService_UpdateWorkspaceSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceSettingServiceServer).UpdateWorkspaceSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceSettingService_UpdateWorkspaceSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceSettingServiceServer).UpdateWorkspaceSetting(ctx, req.(*UpdateWorkspaceSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceSettingService_ServiceDesc is the grpc.ServiceDesc for WorkspaceSettingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceSettingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slash.api.v2.WorkspaceSettingService",
	HandlerType: (*WorkspaceSettingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspaceSetting",
			Handler:    _WorkspaceSettingService_GetWorkspaceSetting_Handler,
		},
		{
			MethodName: "UpdateWorkspaceSetting",
			Handler:    _WorkspaceSettingService_UpdateWorkspaceSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/workspace_setting_service.proto",
}
