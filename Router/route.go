package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Apps() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	route := gin.Default()

	// Module route
	Version1(route)
	Home(route)

	route.Run()
}
