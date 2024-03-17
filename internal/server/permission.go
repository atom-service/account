package server

import (
	"context"

	"github.com/atom-service/account/package/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var PermissionServer = &permissionServer{}

type permissionServer struct {
	protos.UnimplementedPermissionServiceServer
}

func (s *permissionServer) CreateRole(ctx context.Context, request *protos.CreateRoleRequest) (response *protos.CreateRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}

func (s *permissionServer) QueryRoles(ctx context.Context, request *protos.QueryRolesRequest) (response *protos.QueryRolesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRoles not implemented")
}

func (s *permissionServer) UpdateRole(ctx context.Context, request *protos.UpdateRoleRequest) (response *protos.UpdateRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}

func (s *permissionServer) DeleteRole(ctx context.Context, request *protos.DeleteRoleRequest) (response *protos.DeleteRoleResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func (s *permissionServer) CreateResource(ctx context.Context, request *protos.CreateResourceRequest) (response *protos.CreateResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}

func (s *permissionServer) QueryResources(ctx context.Context, request *protos.QueryResourcesRequest) (response *protos.QueryResourcesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryResources not implemented")
}

func (s *permissionServer) DeleteResource(ctx context.Context, request *protos.DeleteResourceRequest) (response *protos.DeleteResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}

func (s *permissionServer) UpdateResource(ctx context.Context, request *protos.UpdateResourceRequest) (response *protos.UpdateResourceResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}

func (s *permissionServer) SummaryForUser(ctx context.Context, request *protos.SummaryForUserRequest) (response *protos.SummaryForUserResponse, err error) {
	response = &protos.SummaryForUserResponse{}
	response.Data = []*protos.UserResourceSummary{}
	response.State = protos.State_SUCCESS
	return nil, nil
}

func (s *permissionServer) ApplyRoleForUser(ctx context.Context, request *protos.ApplyRoleForUserRequest) (response *protos.ApplyRoleForUserResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRoleForUser not implemented")
}

func (s *permissionServer) RemoveRoleForUser(ctx context.Context, request *protos.RemoveRoleForUserRequest) (response *protos.RemoveRoleForUserResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoleForUser not implemented")
}
