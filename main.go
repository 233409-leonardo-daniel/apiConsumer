package main

import (
	"apiconsumer/src/core"
	"apiconsumer/src/order/infrastructure/adapters"
	"apiconsumer/src/order/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectToDataBase()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}

	orderRepo := adapters.NewMySQLRepository(db)
	rabbitRepo := adapters.NewRabbitMQRepository()

	router := gin.Default()

	routes.SetupOrderRoutes(router, orderRepo, rabbitRepo)

	log.Println("Iniciando el Servidor en el puerto 8082...")
	if err := router.Run(":8082"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
