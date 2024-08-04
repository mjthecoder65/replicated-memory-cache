package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mjthecoder65/replicated-memory-cache/cache"
	"github.com/mjthecoder65/replicated-memory-cache/health_check"
	"github.com/mjthecoder65/replicated-memory-cache/utils"
)

func main() {
	config := utils.LoadConfig()

	if config.Environment.Value == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/get/:key", cache.GetKeyHandler)
	router.POST("/set", cache.SetKeyHandler)
	router.POST("/sync", cache.SyncDataHandler)
	router.GET("/health", health_check.HealthCheckHander)

	router.Run(config.Server.Address)
}
