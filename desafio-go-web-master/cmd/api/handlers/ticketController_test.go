package handler

import (
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/bootcamp-go/desafio-go-web/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateServerTicketsHandler(ticketsController *Controller) *gin.Engine {
	gin.SetMode(gin.TestMode)
	server := gin.New()
	
	server.GET("/ticket/getByCountry/:dest", ticketsController.GetTicketsByCountry())
	server.GET("/ticket/getAverage/:dest", ticketsController.AverageDestination())

	return server
}

func TestGetNumbersOfTicketsByDestination(t *testing.T){
	list, err := utils.LoadTicketsFromFile("../../../tickets.csv")
			
	if err != nil {
		panic(err)
	}

	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	
	controller := NewController(service)
	t.Run("Get numbers of tickets by destination", func(t *testing.T){
			countryName := "China"
			expectedNumbers := "178"
			codeStatusExpected := 200

			serverTest := CreateServerTicketsHandler(controller)

			request := httptest.NewRequest(http.MethodGet, "/ticket/getByCountry/"+countryName, nil)

			response := httptest.NewRecorder()

			serverTest.ServeHTTP(response, request)

			assert.Equal(t, codeStatusExpected, response.Result().StatusCode)
			assert.Equal(t, expectedNumbers, response.Body.String())
	})

	t.Run("Get numbers of tickets by destination not found", func(t *testing.T){
		countryName := "FarFarAway"
		expectedNumbers := "0"
		codeStatusExpected := 200

		serverTest := CreateServerTicketsHandler(controller)

		request := httptest.NewRequest(http.MethodGet, "/ticket/getByCountry/"+countryName, nil)

		response := httptest.NewRecorder()

		serverTest.ServeHTTP(response, request)

		assert.Equal(t, codeStatusExpected, response.Result().StatusCode)
		assert.Equal(t, expectedNumbers, response.Body.String())
})
}

func TestGetAveragefTicketsByDestination(t *testing.T){
	list, err := utils.LoadTicketsFromFile("../../../tickets.csv")
			
	if err != nil {
		panic(err)
	}

	repository := tickets.NewRepository(list)
	service := tickets.NewService(repository)
	
	controller := NewController(service)
	
	t.Run("Get average of tickets by destination", func(t *testing.T){
			countryName := "Portugal"
			expectedAverage := "0.029"
			codeStatusExpected := 200


			serverTest := CreateServerTicketsHandler(controller)

			request := httptest.NewRequest(http.MethodGet, "/ticket/getAverage/"+countryName, nil)

			response := httptest.NewRecorder()

			serverTest.ServeHTTP(response, request)

			assert.Equal(t, codeStatusExpected, response.Result().StatusCode)
			assert.Equal(t, expectedAverage, response.Body.String())
	})

	t.Run("Get average of tickets by destination not found", func(t *testing.T){
		countryName := "FarFarAway"
		expectedAverage := "0"
		codeStatusExpected := 200

		serverTest := CreateServerTicketsHandler(controller)

		request := httptest.NewRequest(http.MethodGet, "/ticket/getAverage/"+countryName, nil)

		response := httptest.NewRecorder()

		serverTest.ServeHTTP(response, request)

		assert.Equal(t, codeStatusExpected, response.Result().StatusCode)
		assert.Equal(t, expectedAverage, response.Body.String())
})
}
