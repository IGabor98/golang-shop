package repositories

import (
	"errors"
	"github.com/google/uuid"
	"shop/models"
)

type CartRepositoryMock struct {
	ProductRepository ProductRepository
	carts             []*models.Cart
}

func (r *CartRepositoryMock) CreateCart(productsIDs []uuid.UUID) *models.Cart {
	cart := &models.Cart{
		ID: uuid.New(),
	}

	products := r.ProductRepository.FindByIDs(productsIDs)

	cart.AddProducts(products)

	r.carts = append(r.carts, cart)

	return cart
}

func (r *CartRepositoryMock) FindCartByID(ID uuid.UUID) (*models.Cart, error) {
	for _, cart := range r.carts {
		if cart.ID == ID {
			return cart, nil
		}
	}

	return &models.Cart{}, errors.New("A cart with this ID doesn't exist")
}

func (r *CartRepositoryMock) AddProducts(cartID uuid.UUID, productsIDs []uuid.UUID) error {
	cart, err := r.FindCartByID(cartID)
	if err != nil {
		return errors.New("cart doesn't exist")
	}

	products := r.ProductRepository.FindByIDs(productsIDs)

	cart.AddProducts(products)

	return nil
}
