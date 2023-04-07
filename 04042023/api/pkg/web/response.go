package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Code string  `json:"code"`
	Message string  `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func SendResponse(status int, data interface{}, ctx *gin.Context) {
	var response = response{
		Data: data,
	}

	ctx.JSON(status, response)
}	

func SendErrorResponse(status string, code int, message string, ctx *gin.Context) {
	var errorResponse = ErrorResponse{
		Status: status,
		Code: fmt.Sprint(code),
		Message: message,
	}

	ctx.JSON(code, errorResponse)
}