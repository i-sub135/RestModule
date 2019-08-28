package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(route *gin.Engine) {
	route.GET("/", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Data NotFound",
			"result": map[string]string{
				"pesan": "Micro Services",
			},
		})
	})
}
