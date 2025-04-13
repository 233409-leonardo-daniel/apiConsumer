package controllers

import (
	"apiconsumer/src/order/domain/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrdersController(repo repositories.IOrder) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}
