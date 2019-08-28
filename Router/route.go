package router

import (
	"log"

	. "../Controller"
	. "../Function/Middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Apps() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}
	route := gin.Default()

	// Router Group V1
	V1 := route.Group("api/v1")
	// Use Middleware
	V1.Use(MiddleJwt)

	// Child Route V1
	V1.GET("/role", GetRoleUser)
	V1.GET("/view/history", GetBookingHistory)

	route.Run()
}
