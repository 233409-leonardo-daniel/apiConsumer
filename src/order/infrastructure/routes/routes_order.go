package routes

import (
	usecases "apiconsumer/src/order/application/use_cases"
	"apiconsumer/src/order/domain/repositories"
	"apiconsumer/src/order/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, repo repositories.IOrder, rabbit repositories.IRabbitMQ, wsRepo repositories.IWebSocket) {
	createOrderUseCase := usecases.NewCreateOrder(repo, rabbit)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase, wsRepo)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", createOrderController.Run)
	}
}
