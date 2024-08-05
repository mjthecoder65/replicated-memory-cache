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

	v1 := router.Group("/v1")

	v1.GET("/get/:key", cache.GetKeyHandler)
	v1.POST("/set", cache.SetKeyHandler)
	v1.POST("/sync", cache.SyncDataHandler)
	v1.GET("/health", health_check.HealthCheckHander)

	router.Run(config.Server.Address)
}
