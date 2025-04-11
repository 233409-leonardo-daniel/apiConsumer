package controllers

import (
	usecases "apiconsumer/src/order/application/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	useCase *usecases.CreateOrder
}

func NewCreateOrderController(useCase *usecases.CreateOrder) *CreateOrderController {
	return &CreateOrderController{useCase: useCase}
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

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}
