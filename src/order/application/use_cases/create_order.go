package usecases

import (
	"apiconsumer/src/order/domain/repositories"
	"apiconsumer/src/order/infraestructure/adapters"
)

type IOrder interface {
	Execute(idProduct int32, quantity int32, totalPrice float64, status string) error
}

type CreateOrder struct {
	db repositories.IOrder
}

func NewCreateOrder(db repositories.IOrder) *CreateOrder {
	return &CreateOrder{db: db}
}

func (co *CreateOrder) Execute(idProduct int32, quantity int32, totalPrice float64, status string) error {
	adapters.Execute(
		idProduct,
		quantity,
		totalPrice,
		status,
	)
	return co.db.Save(idProduct, quantity, totalPrice, status)
}
