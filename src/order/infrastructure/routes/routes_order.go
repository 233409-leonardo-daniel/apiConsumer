package routes

import (
	usecases "apiconsumer/src/order/application/use_cases"
	"apiconsumer/src/order/domain/repositories"
	"apiconsumer/src/order/infrastructure/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, repo repositories.IOrder, rabbit repositories.IRabbitMQ) {
	// Configura CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Cambia esto al origen que necesites
	router.Use(cors.New(config))

	createOrderUseCase := usecases.NewCreateOrder(repo, rabbit)
	createOrderController := controllers.NewCreateOrderController(createOrderUseCase)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", createOrderController.Run)
		orderGroup.GET("", controllers.GetAllOrdersController(repo))
	}
}
