package middlewares

import (
	"api/api/pkg/web"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func CheckToken(ctx *gin.Context) {

	var token = ctx.GetHeader("token")

	if token != os.Getenv("ACCESS_TOKEN") {
		web.SendErrorMiddleware("Unauthorized", http.StatusUnauthorized, "Invalid Token", ctx)
		
		return 
	} 

	ctx.Next()

}
