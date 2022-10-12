package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		authorized := check(token)
		if authorized {
			context.Next()
			return
		}
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		context.Abort()
		return
	}
}

func check(token string) bool {
	return true
}
