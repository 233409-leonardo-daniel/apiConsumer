package repositories

import "apiconsumer/src/order/domain/entities"

type IOrder interface {
	Save(idProduct int32, quantity int32, totalPrice float64, status string) error
	GetAll() ([]entities.Order, error)
	GetById(id int32) (*entities.Order, error)
	Delete(id int32) error
}
