package main

import (
	"fmt"
	"github.com/arl/statsviz"
	example "github.com/arl/statsviz/_example"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	//ExampleNetHttp()
	//ExampleGinWeb()
	ExampleEcho()
}
func ExampleNetHttp() {
	statsviz.RegisterDefault()
	log.Print(http.ListenAndServe(":6060", nil))
}

func ExampleGinWeb() {
	go example.Work()

	fmt.Println("Point your browser to http://localhost:8085/debug/statsviz/\\n\\n")

	router := gin.New()
	router.GET("/debug/statsviz/*filepath", func(context *gin.Context) {
		if context.Param("filepath") == "/ws" {
			statsviz.Ws(context.Writer, context.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(context.Writer, context.Request)
	})
	router.Run(":8085")
}

func ExampleEcho() {
	go example.Work()

	ech := echo.New()

	mux := http.NewServeMux()

	_ = statsviz.Register(mux)

	ech.GET("/debug/statsviz/", echo.WrapHandler(mux))

	ech.GET("/debug/statsviz/*", echo.WrapHandler(mux))

	fmt.Println("Point your browser to http://localhost:8082/debug/statsviz/")
	ech.Logger.Fatal(ech.Start(":8082"))
}
