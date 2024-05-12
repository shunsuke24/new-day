package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cont *controller) HealthCheck(c *gin.Context) {
	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "V2",
	})
}
