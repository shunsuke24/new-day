package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/new-day/app"
)

func main() {
	code := goMain()
	os.Exit(code)
}

func goMain() (code int) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found!")

	}
	app.Run()

	return 0
}
