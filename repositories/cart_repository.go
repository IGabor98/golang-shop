package repositories

import (
	"github.com/google/uuid"
	"shop/models"
)

type CartRepository interface {
	CreateCart(products []*models.Product) *models.Cart
	findCartByID(ID uuid.UUID) (*models.Cart, error)
}
