package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/atom-service/account/internal/auth"
	publicAuth "github.com/atom-service/account/package/auth"
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
	loggerPrefix := "AtomService/Account"
	authData := publicAuth.ResolveAuth(ctx)

	if authData != nil {
		if authData.User != nil && authData.User.ID != nil {
			loggerPrefix = fmt.Sprintf("%s %s:%d", loggerPrefix, "userID", *authData.User.ID)
		}

		if authData.Secret != nil && authData.Secret.Key != nil {
			loggerPrefix = fmt.Sprintf("%s %s:%s", loggerPrefix, "secretKey", *authData.Secret.Key)
		}
	}

	// 处理请求之前的逻辑
	slog.DebugContext(ctx, "Received request %s: %s %v", loggerPrefix, info.FullMethod, req)

	// 调用实际的 gRPC 处理程序处理请求
	resp, err := handler(ctx, req)

	// 处理请求之后的逻辑
	if err != nil {
		slog.ErrorContext(ctx, "Error occurred during request %s: %v", loggerPrefix, err)
	} else {
		slog.DebugContext(ctx, "Sent response %s: %v", loggerPrefix, resp)
	}

	return resp, err
}
