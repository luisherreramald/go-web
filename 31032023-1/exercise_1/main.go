package main

import (
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	router.GET("/ping", func (ctx *gin.Context){
		ctx.String(200, "Pong")
	})

	router.Run()
}