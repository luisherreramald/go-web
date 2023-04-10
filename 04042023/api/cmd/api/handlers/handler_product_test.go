package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"api/api/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerProductHandler (productController ControllerProducts) *gin.Engine {
	gin.SetMode(gin.TestMode)

	server := gin.New()

	server.GET("/products", productController.GetAllProducts())
	server.GET("/:id", productController.GetById())
	server.POST("/", productController.Create())
	server.DELETE("/:id", productController.DeleteById())

	return server
}

func TestProductController_GetAllProducts(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		var(
			expectedCodeStatus = 200
		)	

		repositoryLocal := products.NewRepositoryLocal()
		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}

func TestProductController_GetProductById(t *testing.T) {
	t.Run("Should return product by id", func(t *testing.T) {
		var(
			expectedCodeStatus = 200
			idProduct = "1"
			expectedProduct =  `{
				"data": {
        					"id": 1,
        					"name": "Carne de perro alta en proteina",
        					"quantity": 8,
        					"code_value": "8728",
        					"is_published": false,
        					"expiration": "03/01/2025",
        					"price": 50.4
        				}
					}`			
			)

		repositoryLocal := products.NewRepositoryLocal()
		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodGet, "/"+idProduct, nil)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
		assert.JSONEq(t, expectedProduct, response.Body.String())
	})
}

func TestProductController_CreateProductById(t *testing.T) {
	t.Run("Should create new product", func(t *testing.T) {
		var(
			expectedCodeStatus = 201
			newProduct =  `{
						"name": "Producto Nuevo",
        				"quantity": 5,
        				"code_value": "92437sd31",
        				"is_published": true,
        				"expiration": "03/01/2026",
        				"price": 40
        			}
				}`				
		)
		readerProducts := bytes.NewBuffer([]byte(newProduct))
		repositoryLocal := products.NewRepositoryLocal()

		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodPost, "/", readerProducts)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}

func TestProductController_DeleteById(t *testing.T) {
	t.Run("Should delete a product", func(t *testing.T) {
		var(
			expectedCodeStatus = 204		
			idProduct = "1"
		)

		repositoryLocal := products.NewRepositoryLocal()

		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodDelete, "/"+idProduct, nil)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}

func TestProductController_DeleteByIdNotFound(t *testing.T) {
	t.Run("Should delete a product", func(t *testing.T) {
		var(
			expectedCodeStatus = 404		
			idProduct = "8"
		)

		repositoryLocal := products.NewRepositoryLocal()

		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodDelete, "/"+idProduct, nil)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}

func TestProductController_CreateProductByIdBadRequest(t *testing.T) {
	t.Run("Should create new product", func(t *testing.T) {
		var(
			expectedCodeStatus = 400
			newProduct =  `{
						"name": "Producto Nuevo",
        				"quantity": asdasd,
        				"code_value": "924371",
        				"is_published": "sdsdf",
        				"expiration": "03/01/2026",
        				"price": 40
        			}
				}`				
		)
		readerProducts := bytes.NewBuffer([]byte(newProduct))
		repositoryLocal := products.NewRepositoryLocal()

		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodPost, "/", readerProducts)
		response := httptest.NewRecorder()

		server := createServerProductHandler(controllerProducts)

		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}

func TestProductController_CreateProductUnauthorized(t *testing.T) {
	t.Run("Should create new product", func(t *testing.T) {
		var(
			expectedCodeStatus = 401
			newProduct =  `{
						"name": "Producto Nuevo",
        				"quantity": asdasd,
        				"code_value": "924371",
        				"is_published": "sdsdf",
        				"expiration": "03/01/2026",
        				"price": 40
        			}
				}`				
		)
		readerProducts := bytes.NewBuffer([]byte(newProduct))
		repositoryLocal := products.NewRepositoryLocal()

		serviceProducts := products.NewService(repositoryLocal)
		controllerProducts := *NewController(serviceProducts)	
		
		request := httptest.NewRequest(http.MethodPost, "/", readerProducts)
		response := httptest.NewRecorder()

		request.Header.Set("token", "asdasdasd")
		server := createServerProductHandler(controllerProducts)
		
		server.ServeHTTP(response, request)

		assert.Equal(t, expectedCodeStatus, response.Result().StatusCode)
	})
}