package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	router.GET("/ping", func(context *gin.Context) {
		message := "pong"
		context.JSON(http.StatusOK, message)

	})

	server := &http.Server{
		Addr: ":8080",
		/*	Handler:        routes,*/
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
