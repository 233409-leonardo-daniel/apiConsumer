package entities

import "time"

type Order struct {
	IdOrder    int32     `json:"id" gorm:"column:idOrder;primaryKey;autoIncrement"`
	IdProduct  int32     `json:"idProduct" gorm:"column:idProduct"`
	Quantity   int32     `json:"quantity"`
	TotalPrice float64   `json:"totalPrice" gorm:"column:total_price;type:decimal(10,2)"`
	Status     string    `json:"status" gorm:"type:enum('Pending','Shipped','Completed','Cancelled');default:'Pending'"`
	OrderDate  time.Time `json:"orderDate" gorm:"column:orderDate;default:CURRENT_TIMESTAMP"`
}

func NewOrder(idProduct int32, quantity int32, totalPrice float64, status string) *Order {
	return &Order{
		IdProduct:  idProduct,
		Quantity:   quantity,
		TotalPrice: totalPrice,
		Status:     status,
	}
}
