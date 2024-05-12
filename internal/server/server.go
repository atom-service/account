package server

import (
	"context"
	"log/slog"
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
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(serverAuth.ServerUnary, UnaryLogInterceptor))

	proto.RegisterAccountServiceServer(grpcServer, AccountServer)
	proto.RegisterPermissionServiceServer(grpcServer, PermissionServer)
	slog.Info("start server at", "addr", addr)
	reflection.Register(grpcServer)
	return grpcServer.Serve(listen)
}

func UnaryLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	slog.DebugContext(ctx, "Received request", slog.String("method", info.FullMethod), slog.Any("request", req))

	resp, err := handler(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "Error occurred during request", slog.Any("error", err))
	} else {
		slog.DebugContext(ctx, "Sent response", slog.Any("response", resp))
	}

	return resp, err
}
