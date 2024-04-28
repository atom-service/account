package main

import (
	"context"
	"time"

	"github.com/atom-service/account/internal/model"
	"github.com/atom-service/account/internal/server"
	"github.com/yinxulai/goconf"
)

func init() {
	goconf.Declare("port", "8080", true, "Service listening port")
}

func main() {
	// 声明&初始化配置
	goconf.MustLoad()

	context, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	// 初始化 model
	if err := model.Init(context); err != nil {
		panic(err)
	}

	listenAddress := ":" + goconf.MustGet("port")
	panic(server.StartServer(listenAddress))
}
