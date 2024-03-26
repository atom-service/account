package server

import (
	"net"

	"github.com/atom-service/account/internal/auth"
	"github.com/atom-service/account/package/proto"
	"github.com/atom-service/common/logger"
	"google.golang.org/grpc"
)

func StartServer(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	serverAuth := auth.NewServerAuthInterceptor()
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(serverAuth.ServerUnary))

	proto.RegisterAccountServiceServer(grpcServer, AccountServer)
	proto.RegisterPermissionServiceServer(grpcServer, PermissionServer)
	logger.Infof("start server at: %s", addr)
	return grpcServer.Serve(listen)
}
