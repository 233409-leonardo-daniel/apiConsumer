package usecases

import (
	"apiconsumer/src/order/domain/entities"
	"apiconsumer/src/order/domain/repositories"
)

type ViewOrder struct {
	db repositories.IOrder
}

func NewViewOrder(db repositories.IOrder) *ViewOrder {
	return &ViewOrder{db: db}
}

func (vo *ViewOrder) Execute() ([]entities.Order, error) {
	return vo.db.GetAll()
}
