package models

import "github.com/google/uuid"

type Cart struct {
	ID       uuid.UUID
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

func (c *Cart) AddProducts(products []Product) {
	c.Products = append(c.Products, products...)
	c.Total = c.calculateTotal()
}

func (c *Cart) calculateTotal() float64 {
	result := float64(0)

	for _, product := range c.Products {
		result += product.Price
	}

	return result
}
