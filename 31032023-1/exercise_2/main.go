package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Name struct {
	Name string `json:"nombre"`
	LastName string `json:"apellido"`
}

func main () {
	router := gin.Default()
	router.POST("/saludo", func (ctx *gin.Context){
		name := Name{}
		if err := ctx.BindJSON(&name); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		var saludo = fmt.Sprintf("Hola %s %s",name.Name, name.LastName)
		ctx.String(200, saludo)
	})

	router.Run()
}