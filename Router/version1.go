package router

import "github.com/gin-gonic/gin"
import (
	Ctrl "../Controller"
	Midd "../Function/Middleware"
)

// Version1 --> Router Version Api api/v1
func Version1(route *gin.Engine) {
	// Router Group V1
	V1 := route.Group("api/v1")
	// Use Middleware
	V1.Use(Midd.MiddleJwt)

	// Child Route V1
	V1.GET("/role", Ctrl.GetRoleUser)
	V1.GET("/view/history", Ctrl.GetBookingHistory)

}
