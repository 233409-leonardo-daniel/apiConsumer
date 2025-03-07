package routes

import (
	usecases "apiconsumer/src/order/application/use_cases"
	"apiconsumer/src/order/domain/repositories"
	"apiconsumer/src/order/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, repo repositories.IOrder) {
	createOrderUseCase := usecases.NewCreateOrder(repo)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", createOrderController.Run)

	}
}
