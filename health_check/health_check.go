package health_check

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHander(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Ok"})
}
