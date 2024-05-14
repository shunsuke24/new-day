package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/new-day/usecase"
)

type Controller interface {
	HealthCheck(c *gin.Context)
	Recommend(c *gin.Context)
	Send(c *gin.Context)
}

type controller struct {
	usecase usecase.Usecase
}

type Params struct {
	Usecase usecase.Usecase
}

func NewController(p *Params) Controller {
	return &controller{
		usecase: p.Usecase,
	}
}
