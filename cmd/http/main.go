package main

import (
	"context"

	"github.com/protect-we-network/server/internal/core"
	"github.com/protect-we-network/server/internal/packages/config"
	"github.com/protect-we-network/server/internal/packages/logger"
	"github.com/protect-we-network/server/internal/packages/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	logger.Init()
	config.Init()
	context := context.Background()
	clientOptions := options.Client().ApplyURI(config.Mongo.URI)
	client, err := mongo.Connect(context, clientOptions)
	PanicError(err)
	database := client.Database("default")

	// 初始化
	user.Init(database)

	PanicError(core.App())
}
