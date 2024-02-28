package server

import "github.com/atom-service/account/package/protos"

type PermissionServer struct {
	protos.UnimplementedPermissionServer
}

func NewPermissionServer() *PermissionServer {
	return &PermissionServer{}
}
