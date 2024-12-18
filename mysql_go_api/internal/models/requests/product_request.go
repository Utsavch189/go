package requests

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          string    `json:"id"`
	ProductName string    `json:"product_name" validate:"required,alphanum,min=3,max=20"`
	Price       float32   `json:"price" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	IsActive    bool      `json:"is_active"`
}

// Constructor for Product
func NewProduct(product_name string, price float32) *Product {
	uid := uuid.New()
	return &Product{
		Id:          uid.String(),
		ProductName: product_name,
		Price:       price,
		CreatedAt:   time.Now(),
		IsActive:    true,
	}
}
