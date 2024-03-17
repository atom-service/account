package main

import (
	"context"
	"net"
	"time"

	"github.com/atom-service/account/internal/database"
	"github.com/atom-service/account/internal/helper"
	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/internal/server"
	"github.com/atom-service/account/package/auth"
	"github.com/atom-service/account/package/protos"
	"github.com/atom-service/common/config"
	"google.golang.org/grpc"
)

func init() {
	config.Declare("port", "8080", true, "服务监听的端口")
}

func main() {
	// 声明&初始化配置
	config.MustLoad()

	listenAddress := ":" + config.MustGet("port")
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		panic(err)
	}

	context, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	// 初始化数据库
	if err := database.Init(context); err != nil {
		panic(err)
	}

	// 初始化 model
	if err := model.Init(context); err != nil {
		panic(err)
	}

	// 初始化管理员账号
	if err := server.AccountServer.InitAdminUser(context); err != nil {
		panic(err)
	}

	serverAuth := auth.NewServerAuthInterceptor(listenAddress, helper.GodSecretKey, helper.GodSecretValue)
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(serverAuth.ServerUnary))

	protos.RegisterAccountServiceServer(grpcServer, server.AccountServer)
	protos.RegisterPermissionServiceServer(grpcServer, server.PermissionServer)
	panic(grpcServer.Serve(listen))
}
