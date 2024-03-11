package server

import "github.com/atom-service/account/package/protos"

type PermissionServer struct {
	protos.UnimplementedPermissionServiceServer
}

func NewPermissionServer() *PermissionServer {
	return &PermissionServer{}
}
