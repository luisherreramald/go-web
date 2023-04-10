package router

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/api/handlers"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	db []domain.Ticket
}

func NewRouter(routerInstance *gin.Engine, db []domain.Ticket) Router {
	return Router{
		router:  routerInstance,
		db: db,
	}
}

func (routes Router) MapRoutes() {
	repository := tickets.NewRepository(routes.db)
	service := tickets.NewService(repository)
	controller := handler.NewController(service)

	routes.router.GET("/ticket/getByCountry/:dest", controller.GetTicketsByCountry())
	routes.router.GET("/ticket/getAverage/:dest", controller.AverageDestination())
}


