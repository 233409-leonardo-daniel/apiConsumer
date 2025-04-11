package usecases

import (
	"apiconsumer/src/order/domain/repositories"
)

type IOrder interface {
	Execute(idProduct int32, quantity int32, totalPrice float64, status string) error
}

type CreateOrder struct {
	db     repositories.IOrder
	rabbit repositories.IRabbitMQ
}

func NewCreateOrder(db repositories.IOrder, rabbit repositories.IRabbitMQ) *CreateOrder {
	return &CreateOrder{db: db, rabbit: rabbit}
}

func (co *CreateOrder) Execute(idProduct int32, quantity int32, totalPrice float64, status string) error {
	err := co.rabbit.Publish(idProduct, quantity, totalPrice, status)
	if err != nil {
		return err
	}
	return co.db.Save(idProduct, quantity, totalPrice, status)
}
