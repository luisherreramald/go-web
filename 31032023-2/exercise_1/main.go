package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var productos = []Producto{}

func LoadProducts(fileName string) (bool, error) {
	fileReader, err := os.Open(fileName)

	if err != nil {
		return false, fmt.Errorf("File %s does not exist", fileName)
	}

	products, err := ioutil.ReadAll(fileReader)

	if err != nil {
		return false, fmt.Errorf("File %s can not read", fileName)
	}

	json.Unmarshal(products, &productos)	

	return true, nil 
}

func ping(ctx *gin.Context){
	ctx.String(http.StatusOK, "Pong")
}

func getAllProducts(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK, productos)
}

func handleError(code int, message string) map[string]any {
	return map[string]any {
		"code": code,
		"message": message,
	}
}

func getProductById(ctx *gin.Context) {
	var id = ctx.Param("id")
	for _, product := range productos {
		idInt, err := strconv.Atoi(id)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, handleError(500, "Internal Server Error"))
			return
		}
		if idInt == product.Id {
			ctx.IndentedJSON(http.StatusOK, product)
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, handleError(404, "Product Not Found"))
}

func getProductsByPrice(ctx *gin.Context) {
	var price = ctx.Query("priceGt")
	var productsFilter = []Producto{}

	for _, product := range productos {
		priceFloat, err := strconv.ParseFloat(price, 64)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, handleError(500, "Internal Server Error"))
			return
		}
		if product.Price > priceFloat {
			productsFilter = append(productsFilter, product)
		}
	}
	ctx.IndentedJSON(http.StatusOK, productsFilter)
}

func InitServer() {
	router := gin.Default()
	
	router.GET("/ping", ping)
	router.GET("/productos", getAllProducts)
	router.GET("/productos/:id", getProductById)
	router.GET("/productos/search", getProductsByPrice)

	router.Run()
}

func main () {

	defer func(){
		err := recover() 

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Server error")
	}()

	var fileName string = "products.json"
	_, err :=LoadProducts(fileName)

	if err != nil {
		panic(err)
	}

	InitServer()
}