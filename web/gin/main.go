package main

import (
	"github.com/gin-gonic/gin"
	"github.com/megoo/middleware"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
	})

	router.GET("/path", middleware.AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "OK"})
	})
	router.Run(":8080")
}
