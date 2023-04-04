package handlers

import (
	"api/api/internal/domain"
	"api/api/internal/products"
	"errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"fmt"
)

func NewController (serviceProducts products.Service) *ControllerProducts {
	return &ControllerProducts{serviceProducts: serviceProducts}
}

type ControllerProducts struct {
	serviceProducts products.Service
}

func (controllerProduct *ControllerProducts) GetById() gin.HandlerFunc {
	return func (context *gin.Context) {
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}

		product, err := controllerProduct.serviceProducts.GetById(idProduct) 

		if err != nil {
			
			if errors.Is(err, products.ErrServiceNotFound) {

				context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil})
			
			return 
			}
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data": nil})
			
			return 
		}

		context.JSON(http.StatusOK, gin.H{"message":"Success", "data": product})
	}
}

func (controllerProducts *ControllerProducts) Create() gin.HandlerFunc {
	type Request struct {
		Name string `json:"name" validate:"required"`
		Quantity int `json:"quantity" validate:"number, required"`
		CodeValue string `json:"code_value" validate:"required"`
		IsPublished bool `json:"is_published"`
		Expiration string `json:"expiration" validate:"required"`
		Price float64 `json:"price" validate:"number,required"`
	}

	return func(context *gin.Context) {
		var req Request
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}

		product := &domain.Product{
			Name: req.Name,
			Quantity: req.Quantity,
			CodeValue: req.CodeValue,
			IsPublished: req.IsPublished,
			Expiration: req.Expiration,
			Price: req.Price,
		}

		err := controllerProducts.serviceProducts.Create(product) 

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data":nil})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"message": "Success", "data": product})
	}
}

func (controllerProducts *ControllerProducts) UpdateById() gin.HandlerFunc {
	type Request struct {
		Name string `json:"name" validate:"required"`
		Quantity int `json:"quantity" validate:"number, required"`
		CodeValue string `json:"code_value" validate:"required"`
		IsPublished bool `json:"is_published"`
		Expiration string `json:"expiration" validate:"required"`
		Price float64 `json:"price" validate:"number,required"`
	}

	return func(context *gin.Context) {
		var req Request
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}

		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}	

		product := &domain.Product{
			Id: idProduct,
			Name: req.Name,
			Quantity: req.Quantity,
			CodeValue: req.CodeValue,
			IsPublished: req.IsPublished,
			Expiration: req.Expiration,
			Price: req.Price,
		}

		err = controllerProducts.serviceProducts.Update(product, idProduct) 

		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				context.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data":nil})
			}

			context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data":nil})

			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Success", "data": product})
	}
}

func (controllerProducts *ControllerProducts) DeleteById() gin.HandlerFunc {
	
	return func(context *gin.Context) {
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}

		err = controllerProducts.serviceProducts.Delete(idProduct) 

		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				context.JSON(http.StatusNotFound, gin.H{"message":err.Error()})
			}
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Product was remove"})
	}
}

func (controllerProducts *ControllerProducts) UpdatePartial() gin.HandlerFunc {
	type Request struct {
		Name string `json:"name" validate:"required"`
		Quantity int `json:"quantity" validate:"number, required"`
		CodeValue string `json:"code_value" validate:"required"`
		IsPublished bool `json:"is_published"`
		Expiration string `json:"expiration" validate:"required"`
		Price float64 `json:"price" validate:"number,required"`
	}

	return func(context *gin.Context) {
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
			return 
		}
		var product *domain.Product 
		product, err = controllerProducts.serviceProducts.GetById(idProduct)

		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound){
			context.JSON(http.StatusNotFound, gin.H{"message":err.Error()})
			
			return
			}	

			context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}

		if err:= context.ShouldBindJSON(&product); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request error"})
			return 
		}

		product.Id = idProduct

		if err := controllerProducts.serviceProducts.Update(product, idProduct); err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			}
			context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return

		}

		context.JSON(http.StatusOK, gin.H{"message": "Succes", "data": product})
	}
}