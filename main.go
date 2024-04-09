package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"task/app/router"
	"task/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
	init := config.Init()
	app := router.Init(init)
	_ = app.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
