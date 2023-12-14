package core

import (
	"github.com/gin-gonic/gin"
	"github.com/protect-we-network/server/internal/core/routers"
	"github.com/protect-we-network/server/internal/packages/config"
)

func App() error {
	engine := gin.Default()
	routers.Init(engine)
	return engine.Run(config.Port)
}
