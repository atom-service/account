package server

import (
	"context"

	"github.com/atom-service/account/package/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PermissionServer struct {
	protos.UnimplementedPermissionServiceServer
}

func NewPermissionServer() *PermissionServer {
	return &PermissionServer{}
}


func (s *PermissionServer) CreateRole(ctx context.Context,request *protos.CreateRoleRequest) (result *protos.CreateRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}

func (s *PermissionServer) QueryRoles(ctx context.Context,request *protos.QueryRolesRequest) (result *protos.QueryRolesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRoles not implemented")
}

func (s *PermissionServer) UpdateRole(ctx context.Context,request *protos.UpdateRoleRequest) (result *protos.UpdateRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}

func (s *PermissionServer) DeleteRole(ctx context.Context,request *protos.DeleteRoleRequest) (result *protos.DeleteRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func (s *PermissionServer) CreateResource(ctx context.Context,request *protos.CreateResourceRequest) (result *protos.CreateResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}

func (s *PermissionServer) QueryResources(ctx context.Context,request *protos.QueryResourcesRequest) (result *protos.QueryResourcesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResources not implemented")
}

func (s *PermissionServer) DeleteResource(ctx context.Context,request *protos.DeleteResourceRequest) (result *protos.DeleteResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}

func (s *PermissionServer) UpdateResource(ctx context.Context,request *protos.UpdateResourceRequest) (result *protos.UpdateResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}

func (s *PermissionServer) AddRoleForUser(ctx context.Context,request *protos.AddRoleForUserRequest) (result *protos.AddRoleForUserResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoleForUser not implemented")
}

func (s *PermissionServer) RemoveRoleForUser(ctx context.Context,request *protos.RemoveRoleForUserRequest) (result *protos.RemoveRoleForUserResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoleForUser not implemented")
}

func (s *PermissionServer) Check(ctx context.Context,request *protos.CheckRequest) (result *protos.CheckResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
