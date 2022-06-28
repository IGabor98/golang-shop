package repositories

import (
	"errors"
	"github.com/google/uuid"
	"shop/models"
)

type CartRepositoryMock struct {
	Carts             []*models.Cart
	ProductRepository ProductRepository
}

func (r *CartRepositoryMock) CreateCart(productsIDs []uuid.UUID) *models.Cart {
	cart := &models.Cart{
		ID: uuid.New(),
	}

	products := r.ProductRepository.FindByIDs(productsIDs)

	cart.AddProducts(products)

	r.Carts = append(r.Carts, cart)

	return cart
}

func (r *CartRepositoryMock) FindCartByID(ID uuid.UUID) (*models.Cart, error) {
	for _, cart := range r.Carts {
		if cart.ID == ID {
			return cart, nil
		}
	}

	return &models.Cart{}, errors.New("A cart with this ID doesn't exist")
}
