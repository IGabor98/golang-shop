package repositories

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"shop/models"
	"testing"
)

func TestCreateCart(t *testing.T) {
	repository := CartRepositoryMock{
		ProductRepository: &ProductRepositoryMock{},
	}
	productsIDs := createProducts(repository.ProductRepository)

	cart := repository.CreateCart(productsIDs)

	assert.Len(t, cart.Products, 2)
}

func TestCartRepositoryMock_FindCartByID(t *testing.T) {
	repository := CartRepositoryMock{
		ProductRepository: &ProductRepositoryMock{},
	}

	cart := &models.Cart{
		ID: uuid.New(),
	}
	repository.carts = append(repository.carts, cart)

	result, err := repository.FindCartByID(cart.ID)

	assert.NoError(t, err)
	assert.Equal(t, cart, result)
}

func TestCartRepositoryMock_FindCartByID_Error(t *testing.T) {
	repository := CartRepositoryMock{
		ProductRepository: &ProductRepositoryMock{},
	}

	result, err := repository.FindCartByID(uuid.New())

	assert.Error(t, err)
	assert.Equal(t, &models.Cart{}, result)
}

func TestCartRepositoryMock_AddProducts(t *testing.T) {
	repository := CartRepositoryMock{
		ProductRepository: &ProductRepositoryMock{},
	}

	cart := &models.Cart{
		ID: uuid.New(),
	}
	repository.carts = append(repository.carts, cart)

	productsIDs := createProducts(repository.ProductRepository)

	err := repository.AddProducts(cart.ID, productsIDs)

	assert.NoError(t, err)
	assert.Len(t, cart.Products, 2)
	assert.Equal(t, cart.Products[0].ID, productsIDs[0])
	assert.Equal(t, cart.Products[1].ID, productsIDs[1])
}

func TestCartRepositoryMock_AddProducts_Error(t *testing.T) {
	repository := CartRepositoryMock{
		ProductRepository: &ProductRepositoryMock{},
	}

	cartID := uuid.New()
	productIDs := make([]uuid.UUID, 1)
	productIDs[0] = uuid.New()
	err := repository.AddProducts(cartID, productIDs)

	assert.Error(t, err)
}

func createProducts(productRepository ProductRepository) []uuid.UUID {
	result := make([]uuid.UUID, 2)

	firstProduct, _ := productRepository.Create(models.Product{
		Name:        "First product",
		Price:       float64(1),
		Description: "First description",
	})
	secondProduct, _ := productRepository.Create(models.Product{
		Name:        "Second product",
		Price:       float64(23),
		Description: "Second description",
	})
	result[0] = firstProduct.ID
	result[1] = secondProduct.ID

	return result
}
