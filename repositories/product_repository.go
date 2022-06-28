package repositories

import (
	"github.com/google/uuid"
	"shop/models"
)

type ProductRepository interface {
	Create(product models.Product) (*models.Product, error)
	Update(product models.Product) (*models.Product, error)
	Delete(ID uuid.UUID) error
	FindByName(name string) []*models.Product
	GetAll() *[]models.Product
	FindByIDs(IDs []uuid.UUID) []*models.Product
}
