package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-day/domain/model"
)

type UserInformation struct {
	ID           int    `json:"id"`
	IconImageURL string `json:"icon_image_url"`
	Name         string `json:"name"`
}

func (cont *controller) CreateUser(c *gin.Context) {
	var userInfo UserInformation
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()

	usecaseUserInfo := &model.User{
		ID:           userInfo.ID,
		IconImageURL: userInfo.IconImageURL,
		Name:         userInfo.Name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	createdUser, err := cont.usecase.CreateUser(usecaseUserInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, createdUser)
}
