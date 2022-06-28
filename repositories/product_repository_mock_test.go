package repositories

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"shop/models"
	"testing"
)

func TestCreate(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := getTestProduct()
	createdProduct, err := repository.Create(testProduct)

	assert.NoError(t, err)
	assert.IsType(t, createdProduct.ID, uuid.UUID{})
	assert.Equal(t, testProduct.Name, createdProduct.Name)
	assert.Equal(t, testProduct.Price, createdProduct.Price)
	assert.Equal(t, testProduct.Description, createdProduct.Description)
}

func TestSuccessfulUpdate(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := getTestProduct()
	createdProduct, _ := repository.Create(testProduct)

	updatingProduct := createdProduct
	updatingProduct.Name = "Updated name"
	updatingProduct.Price = float64(2)
	updatingProduct.Description = "Updated description"

	err := repository.Update(*updatingProduct)

	assert.NoError(t, err)

	updatedProduct, err := repository.FindByID(createdProduct.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdProduct.ID, updatedProduct.ID)
	assert.Equal(t, updatingProduct.Name, updatedProduct.Name)
	assert.Equal(t, updatingProduct.Price, updatedProduct.Price)
	assert.Equal(t, updatingProduct.Description, updatedProduct.Description)
}

func TestErrorUpdate(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := getTestProduct()

	err := repository.Update(testProduct)

	assert.Error(t, err)
}

func getTestProduct() models.Product {
	return models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}
}
