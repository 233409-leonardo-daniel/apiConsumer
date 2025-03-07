package adapters

import (
	"apiconsumer/src/order/domain/entities"
	"apiconsumer/src/order/domain/repositories"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repositories.IOrder {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Save(idProduct int32, quantity int32, totalPrice float64, status string) error {
	order := entities.NewOrder(idProduct, quantity, totalPrice, status)
	result := repo.db.Create(order)
	return result.Error
}

func (repo *MySQLRepository) GetAll() ([]entities.Order, error) {
	var orders []entities.Order
	result := repo.db.Find(&orders)
	return orders, result.Error
}

func (repo *MySQLRepository) GetById(id int32) (*entities.Order, error) {
	var order entities.Order
	result := repo.db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *MySQLRepository) Update(id int32, idProduct int32, quantity int32, totalPrice float64, status string) error {
	var order entities.Order
	result := repo.db.First(&order, id)
	if result.Error != nil {
		return result.Error
	}

	order.IdProduct = idProduct
	order.Quantity = quantity
	order.TotalPrice = totalPrice
	order.Status = status

	result = repo.db.Save(&order)
	return result.Error
}

func (repo *MySQLRepository) Delete(id int32) error {
	var order entities.Order
	result := repo.db.First(&order, id)
	if result.Error != nil {
		return result.Error
	}
	result = repo.db.Delete(&order)
	return result.Error
}
