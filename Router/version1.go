package router

import "github.com/gin-gonic/gin"
import (
	. "../Controller"
	. "../Function/Middleware"
)

func Version1(route *gin.Engine) {
	// Router Group V1
	V1 := route.Group("api/v1")
	// Use Middleware
	V1.Use(MiddleJwt)

	// Child Route V1
	V1.GET("/role", GetRoleUser)
	V1.GET("/view/history", GetBookingHistory)
}
