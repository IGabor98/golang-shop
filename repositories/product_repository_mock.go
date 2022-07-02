package repositories

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"shop/models"
)

type ProductRepositoryMock struct {
	products []models.Product
}

func (r *ProductRepositoryMock) Create(product models.Product) (models.Product, error) {
	product.ID = uuid.New()

	r.products = append(r.products, product)

	return r.FindByID(product.ID)
}

func (r *ProductRepositoryMock) Update(product models.Product) error {
	for k, v := range r.products {
		if v.ID == product.ID {
			r.products[k] = product

			return nil
		}
	}

	return errors.New("a product with this ID doesn't exist")
}

func (r *ProductRepositoryMock) FindByID(ID uuid.UUID) (models.Product, error) {
	for _, v := range r.products {
		if v.ID == ID {
			return v, nil
		}
	}
	return models.Product{}, errors.New("a product with this ID doesn't exist")
}

func (r *ProductRepositoryMock) FindByIDs(IDs []uuid.UUID) []models.Product {
	result := make([]models.Product, 0)

	for _, v := range r.products {
		if slices.Contains(IDs, v.ID) {
			result = append(result, v)
		}
	}
	return result
}

func (r *ProductRepositoryMock) FindByName(name string) []models.Product {
	result := make([]models.Product, 0)

	for _, v := range r.products {
		if v.Name == name {
			result = append(result, v)
		}
	}
	return result
}

func (r *ProductRepositoryMock) GetAll() []models.Product {
	return r.products
}

func (r *ProductRepositoryMock) Delete(ID uuid.UUID) error {
	index, err := r.findIndexByID(ID)

	if err != nil {
		return err
	}

	r.products = append(r.products[:index], r.products[index+1:]...)

	return nil
}

func (r *ProductRepositoryMock) findIndexByID(ID uuid.UUID) (int, error) {
	for k, v := range r.products {
		if v.ID == ID {
			return k, nil
		}
	}
	return 0, errors.New("a product with this ID doesn't exist")
}
