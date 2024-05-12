package app

import (
	"github.com/new-day/env"
)

func Run() {
	// router := gin.Default()
	conf := env.NewConfig()
	router := initializeServer(conf)
	router.CORS()

	router.Route()
	router.Engine.Run()
}
