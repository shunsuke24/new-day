package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cont *controller) Recommend(c *gin.Context) {
	// Return
	result, err := cont.usecase.Recommend()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, result)
}
