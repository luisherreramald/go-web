package main

import (
	"03042023-1/exercise_1/cmd/server/handlers"
	"03042023-1/exercise_1/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	db := []*internal.Product{}

	serviceProduct := internal.NewServiceProduct(db, len(db))
	controllerProduct := handlers.NewControllerProducts(serviceProduct)

	server := gin.Default()

	server.GET("/ping", func(context *gin.Context){
		context.String(http.StatusOK, "pong")
	})

	productRouter := server.Group("/products") 
	{
		productRouter.POST("", controllerProduct.CreateProduct())
		productRouter.GET(":id", controllerProduct.GetProduct())
	}

	server.Run()
}