package middleware

import (
	"net/http"
	"res_api/utiles"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Please provide a valid token"})
		return
	}

	userId, err := utiles.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized. Please provide a valid token"})
	}

	context.Set("userId", userId)

	context.Next()

}
