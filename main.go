package main

import (
	"context"
	"time"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/internal/server"
	"github.com/atom-service/common/config"
	"github.com/atom-service/common/logger"
)

func init() {
	config.Declare("port", "8080", true, "服务监听的端口")
}

func main() {
	// 声明&初始化配置
	config.MustLoad()
	logger.SetLevel(logger.InfoLevel)

	context, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	// 初始化 model
	if err := model.Init(context); err != nil {
		panic(err)
	}

	listenAddress := ":" + config.MustGet("port")
	panic(server.StartServer(listenAddress))
}
