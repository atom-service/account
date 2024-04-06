package server

import (
	"log/slog"
	"net"

	"github.com/atom-service/account/internal/auth"
	"github.com/atom-service/account/package/proto"
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
	slog.Info("start server at: %s", addr)
	return grpcServer.Serve(listen)
}
