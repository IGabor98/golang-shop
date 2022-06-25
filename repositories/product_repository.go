package repositories

import "shop/models"

type ProductRepository interface {
	Create(product models.Product) (*models.Product, error)
	Update(product models.Product) (*models.Product, error)
	Delete(ID string) error
	FindByName(name string) []*models.Product
	GetAll() *[]models.Product
}
