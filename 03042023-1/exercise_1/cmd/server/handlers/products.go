package handlers

import (
	"03042023-1/exercise_1/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewControllerProducts(serviceProduct *internal.ServiceProducts) *ControllerProducts {
	return &ControllerProducts{serviceProduct: serviceProduct}
}

type ControllerProducts struct {
	serviceProduct *internal.ServiceProducts
}

func (controllerProduct *ControllerProducts) CreateProduct() gin.HandlerFunc {

	return func (context *gin.Context) {
		var req internal.Product

		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Data not valid", "data": nil})
			return 
		}

		product, err := controllerProduct.serviceProduct.Save(&req) 
		
		if err != nil {

			if errors.Is(err, internal.ErrProductExist) {
				context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "The product has already been registered", "data": nil})
				return
			}

			context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data": nil})
			
			return 
		}

		context.JSON(http.StatusCreated, gin.H{"message": "succes", "data": product})
	}
}

func (controllerProduct *ControllerProducts) GetProduct() gin.HandlerFunc {

	return func (context *gin.Context) {
		var idProduct, err = strconv.Atoi(context.Param("id"))


		product, err := controllerProduct.serviceProduct.GetProductById(idProduct) 

		if err != nil {
			
			if errors.Is(err, internal.ErrProductNotFound) {

				context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
			
			return 
			}

			context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data": nil})
			
			return 
		}

		context.JSON(http.StatusOK, gin.H{"data": product})
	}
}