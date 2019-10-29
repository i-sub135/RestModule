package functionMiddleware

import (
	"net/http"

	. "../Jwt"

	"github.com/gin-gonic/gin"
)

func MiddleJwt(res *gin.Context) {
	token := res.GetHeader("x-access-token")

	if token == "" {
		res.JSON(http.StatusNotAcceptable, gin.H{
			"code":   http.StatusNotAcceptable,
			"error":  "Token can not be empty",
			"result": map[string]string{},
		})
		res.Abort()
		return
	}
	_, err := ValidJwt(token)
	if err != nil {
		res.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"error":  "Token Not Valid",
			"result": map[string]string{},
		})
		res.Abort()
	}
	res.Next()
}
