package repositories

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"shop/models"
	"testing"
)

func TestCreate(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}
	createdProduct, err := repository.Create(testProduct)

	assert.NoError(t, err)
	assert.IsType(t, createdProduct.ID, uuid.UUID{})
	assert.Equal(t, testProduct.Name, createdProduct.Name)
	assert.Equal(t, testProduct.Price, createdProduct.Price)
	assert.Equal(t, testProduct.Description, createdProduct.Description)
}

func TestSuccessfulUpdate(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}
	createdProduct, _ := repository.Create(testProduct)

	updatingProduct := createdProduct
	updatingProduct.Name = "Updated name"
	updatingProduct.Price = float64(2)
	updatingProduct.Description = "Updated description"

	err := repository.Update(updatingProduct)

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
	testProduct := models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}

	err := repository.Update(testProduct)

	assert.Error(t, err)
}

func TestSuccessfulFindByID(t *testing.T) {
	repository := ProductRepositoryMock{}
	testProduct := models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}
	createdProduct, _ := repository.Create(testProduct)

	foundProduct, err := repository.FindByID(createdProduct.ID)
	assert.NoError(t, err)
	assert.Equal(t, createdProduct, foundProduct)
}

func TestErrorFindByID(t *testing.T) {
	repository := ProductRepositoryMock{}
	_, err := repository.FindByID(uuid.New())

	assert.Error(t, err)
}

func TestSuccessfulFindByIDs(t *testing.T) {
	ids := make([]uuid.UUID, 0)
	repository := ProductRepositoryMock{}
	testProduct := models.Product{
		Name:        "Test product",
		Price:       float64(1),
		Description: "Test description",
	}
	createdFirstProduct, _ := repository.Create(testProduct)
	createdSecondProduct, _ := repository.Create(testProduct)

	ids = append(ids, createdFirstProduct.ID, createdSecondProduct.ID)
	products := repository.FindByIDs(ids)

	assert.Contains(t, products, createdFirstProduct)
	assert.Contains(t, products, createdSecondProduct)
}

func TestEmptyResultFindByIDs(t *testing.T) {
	repository := ProductRepositoryMock{}

	ids := make([]uuid.UUID, 2)
	ids[0], ids[1] = uuid.New(), uuid.New()

	result := repository.FindByIDs(ids)

	assert.Equal(t, make([]models.Product, 0), result)
}

func TestFindByName(t *testing.T) {
	repository := ProductRepositoryMock{}
	neededProduct := models.Product{
		Name: "Needed product",
	}
	notNeededProduct := models.Product{
		Name: "Not needed product",
	}

	neededProduct, _ = repository.Create(neededProduct)
	_, _ = repository.Create(notNeededProduct)

	result := repository.FindByName(neededProduct.Name)

	assert.Len(t, result, 1)
	assert.Contains(t, result, neededProduct)
}

func TestGetAll(t *testing.T) {
	repository := ProductRepositoryMock{}
	firstProduct, _ := repository.Create(models.Product{
		Name:        "First product",
		Price:       float64(1),
		Description: "First description",
	})
	secondProduct, _ := repository.Create(models.Product{
		Name:        "Second product",
		Price:       float64(23),
		Description: "Second description",
	})

	result := repository.GetAll()

	assert.Contains(t, result, firstProduct)
	assert.Contains(t, result, secondProduct)
}

func TestSuccessfulDelete(t *testing.T) {
	repository := ProductRepositoryMock{}
	firstProduct, _ := repository.Create(models.Product{
		Name:        "First product",
		Price:       float64(1),
		Description: "First description",
	})

	err := repository.Delete(firstProduct.ID)

	assert.NoError(t, err)
	assert.NotContains(t, repository.products, firstProduct)
}

func TestErrorDelete(t *testing.T) {
	repository := ProductRepositoryMock{}

	err := repository.Delete(uuid.New())

	assert.Error(t, err)
}
