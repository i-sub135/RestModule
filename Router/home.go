package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home --- root Router
func Home(route *gin.Engine) {
	route.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Home Api",
			"result": map[string]string{
				"pesan": "Micro Services",
			},
		})
	})
}
