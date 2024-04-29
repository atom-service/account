package server

import (
	"log"
	"net"

	"github.com/atom-service/account/internal/auth"
	"github.com/atom-service/account/package/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	log.Printf("start server at: %s", addr)
	reflection.Register(grpcServer)
	return grpcServer.Serve(listen)
}
