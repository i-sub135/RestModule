package controller

import (
	"net/http"

	Model "../Model"
	"github.com/gin-gonic/gin"
)

// GetRoleUser -- Get Data User
func GetRoleUser(res *gin.Context) {
	data := Model.GetData()
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

// GetBookingHistory -- Get data History Biiking
func GetBookingHistory(res *gin.Context) {
	data := Model.GetDataHistory()
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
