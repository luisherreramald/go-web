package main

import (
	"api/api/cmd/api/handlers"
	"api/api/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"api/api/pkg/store"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	err = store.LoadProducts()

	if err != nil {
		panic(err)
	}
	
	repositoryLocal := products.NewRepositoryLocal()
	serviceProducts := products.NewService(repositoryLocal)
	controllerProducts := handlers.NewController(serviceProducts)

	server := gin.Default()

	server.GET("/ping", func (context *gin.Context)  {
		context.String(200, "Pong")
	})

	products := server.Group("/products")
	{
		products.GET("/:id", controllerProducts.GetById())
		products.POST("", controllerProducts.Create())
		products.PUT("/:id", controllerProducts.UpdateById())
		products.DELETE("/:id", controllerProducts.DeleteById())
		products.PATCH("/:id", controllerProducts.UpdatePartial())
	}

	server.Run()
}