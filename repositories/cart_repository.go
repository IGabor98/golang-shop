package repositories

import (
	"github.com/google/uuid"
	"shop/models"
)

type CartRepository interface {
	CreateCart(productsIDs []uuid.UUID) *models.Cart
	FindCartByID(ID uuid.UUID) (*models.Cart, error)
}
