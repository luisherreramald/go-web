package main

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/api/router"
	"github.com/bootcamp-go/desafio-go-web/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := utils.LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	router := router.NewRouter(r, list)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}