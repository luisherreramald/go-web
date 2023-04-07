package handlers

import (
	"api/api/internal/domain"
	"api/api/internal/products"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"api/api/pkg/web"
	"github.com/gin-gonic/gin"
)

func NewController (serviceProducts products.Service) *ControllerProducts {
	return &ControllerProducts{serviceProducts: serviceProducts}
}

type ControllerProducts struct {
	serviceProducts products.Service
}

func (controllerProduct *ControllerProducts) GetById() gin.HandlerFunc {
	return func (context *gin.Context) {

		var token = context.GetHeader("token")

		if token != os.Getenv("ACCESS_TOKEN") {
			web.SendErrorResponse("Unauthorized", http.StatusUnauthorized, "Invalid Token", context)
			return 
		} 

		var idProduct, err = strconv.Atoi(context.Param("id"))

		if err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
			return 
		}

		product, err := controllerProduct.serviceProducts.GetById(idProduct) 

		if err != nil {
			
			if errors.Is(err, products.ErrServiceNotFound) {
				web.SendErrorResponse("NotFound", http.StatusNotFound, err.Error(), context)
				return 
			}
			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return 
		}
		web.SendResponse(http.StatusOK, product, context)
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
		var token = context.GetHeader("token")

		if token != os.Getenv("ACCESS_TOKEN") {
			web.SendErrorResponse("Unauthorized", http.StatusUnauthorized, "Invalid Token", context)
			return 
		} 

		if err := context.ShouldBindJSON(&req); err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
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

			if errors.Is(err, products.ErrServiceNotUnique) {
				web.SendErrorResponse("Conflict", http.StatusConflict, err.Error(), context)
				return
			}

			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return
		}

		web.SendResponse(http.StatusCreated, product, context)
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
		var token = context.GetHeader("token")
		var req Request
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if token != os.Getenv("ACCESS_TOKEN") {
			web.SendErrorResponse("Unauthorized", http.StatusUnauthorized, "Invalid Token", context)
			return 
		} 

		if err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
			return 
		}

		if err := context.ShouldBindJSON(&req); err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
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
			fmt.Println(err)
			if errors.Is(err, products.ErrServiceNotFound) {
				web.SendErrorResponse("NotFound", http.StatusNotFound, err.Error(), context)
				return 
			}

			if errors.Is(err, products.ErrServiceNotUnique) {
				web.SendErrorResponse("Conflict", http.StatusConflict, err.Error(), context)
				return
			}	

			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return
		}

		web.SendResponse(http.StatusOK, product, context)
	}
}

func (controllerProducts *ControllerProducts) DeleteById() gin.HandlerFunc {
	
	return func(context *gin.Context) {
		var token = context.GetHeader("token")
		var idProduct, err = strconv.Atoi(context.Param("id"))

		if token != os.Getenv("ACCESS_TOKEN") {
			web.SendErrorResponse("Unauthorized", http.StatusUnauthorized, "Invalid Token", context)
			return 
		} 

		if err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
			return 
		}

		err = controllerProducts.serviceProducts.Delete(idProduct) 

		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				web.SendErrorResponse("NotFound", http.StatusNotFound, err.Error(), context)
				return
			}
			
			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return
		}

		web.SendResponse(http.StatusNoContent, nil, context)
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
		var token = context.GetHeader("token")

		if token != os.Getenv("ACCESS_TOKEN") {
			web.SendErrorResponse("Unauthorized", http.StatusUnauthorized, "Invalid Token", context)
			return 
		} 

		if err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
			return 
		}
		var product domain.Product 
		product, err = controllerProducts.serviceProducts.GetById(idProduct)

		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound){
				web.SendErrorResponse("NotFound", http.StatusNotFound, err.Error(), context)
				return	
			}

			if errors.Is(err, products.ErrServiceNotUnique) {
				web.SendErrorResponse("Conflict", http.StatusConflict, err.Error(), context)
				return
			}	

			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return
		}

		if err:= context.ShouldBindJSON(&product); err != nil {
			web.SendErrorResponse("BadRequest", http.StatusBadRequest, "Invalid Request", context)
			return 
		}

		product.Id = idProduct

		if err := controllerProducts.serviceProducts.Update(&product, idProduct); err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				web.SendErrorResponse("NotFound", http.StatusNotFound, err.Error(), context)
				return
			}

			if errors.Is(err, products.ErrServiceNotUnique) {
				web.SendErrorResponse("Conflict", http.StatusConflict, err.Error(), context)
				return
			}	
	
			web.SendErrorResponse("InternalServerError", http.StatusInternalServerError, "Internal Server Error", context)
			return

		}

		web.SendResponse(http.StatusOK, product, context)
	}
}