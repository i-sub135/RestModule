package controller

import (
	"net/http"

	. "../Model"
	"github.com/gin-gonic/gin"
)

func GetRoleUser(res *gin.Context) {
	data := GetData()
	if len(data) > 0 {
		res.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Ok",
			"result":  data,
		})

	} else {
		res.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Data NotFound",
			"result":  map[string]string{},
		})
		res.Abort()
	}
}

func GetBookingHistory(res *gin.Context) {
	data := GetDataHistory()
	if len(data) > 0 {
		res.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Ok",
			"result":  data,
		})

	} else {
		res.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Data NotFound",
			"message": map[string]string{},
		})
		res.Abort()
	}
}
