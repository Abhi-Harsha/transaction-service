package server

import (
	"fmt"
	"net/http"

	cfg "github.com/Abhi-Harsha/transaction-service/app/config"
	"github.com/Abhi-Harsha/transaction-service/app/logger"
	"github.com/gin-gonic/gin"
)

func InitializeApp(config string) {
	conf := cfg.NewAppConfig(config)
	logger.InitLogger(conf.Logger)

	logger.GetLogger().Info(fmt.Sprintf("config: %v", conf))

	engine := setupApp(conf.Server)

	port := fmt.Sprintf("%d", conf.Server.Port)

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.Run(":" + port)
}

func setupApp(serverConfig cfg.ServerConfig) *gin.Engine {
	gin.SetMode(serverConfig.GinMode)
	return gin.New()
}
