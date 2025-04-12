package controllers

import (
	usecases "apiconsumer/src/order/application/use_cases"
	"bytes"
	"encoding/json"
	"net/http"

	"apiconsumer/src/order/domain/repositories"

	"log"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase       *usecases.CreateOrder
	webSocketRepo repositories.IWebSocket
}

func NewCreateOrderController(useCase *usecases.CreateOrder, wsRepo repositories.IWebSocket) *CreateOrderController {
	return &CreateOrderController{useCase: useCase, webSocketRepo: wsRepo}
}

func (co *CreateOrderController) Run(c *gin.Context) {
	var input struct {
		IdProduct  int32   `json:"idProduct"`
		Quantity   int32   `json:"quantity"`
		TotalPrice float64 `json:"totalPrice"`
		Status     string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := co.useCase.Execute(input.IdProduct, input.Quantity, input.TotalPrice, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = sendOrderStatusToWebSocket(input.Status)
	if err != nil {
		log.Printf("Error al enviar estado de la orden al WebSocket: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func sendOrderStatusToWebSocket(status string) error {
	url := "http://localhost:8083/ws" // Cambia esto a la URL de tu servidor WebSocket

	message := map[string]string{"status": status}
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	return err
}
