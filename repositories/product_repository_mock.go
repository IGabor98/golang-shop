package repositories

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"shop/models"
)

type ProductRepositoryMock struct {
	Products []models.Product
}

func (r *ProductRepositoryMock) Create(product models.Product) (*models.Product, error) {
	product.ID = uuid.New()

	r.Products = append(r.Products, product)

	return &product, nil
}

func (r *ProductRepositoryMock) Update(product models.Product) (*models.Product, error) {
	p, _, err := r.FindByID(product.ID)
	if err != nil {
		return p, err
	}

	return &product, nil
}

func (r *ProductRepositoryMock) FindByID(ID uuid.UUID) (*models.Product, int, error) {
	for k, v := range r.Products {
		if v.ID == ID {
			return &v, k, nil
		}
	}
	return &models.Product{}, 0, nil
}

func (r *ProductRepositoryMock) FindByIDs(IDs []uuid.UUID) []*models.Product {
	result := make([]*models.Product, 0)

	for _, v := range r.Products {
		if slices.Contains(IDs, v.ID) {
			result = append(result, &v)
		}
	}
	return result
}

func (r *ProductRepositoryMock) FindByName(name string) []*models.Product {
	result := make([]*models.Product, 0)

	for _, v := range r.Products {
		if v.Name == name {
			result = append(result, &v)
		}
	}
	return result
}

func (r *ProductRepositoryMock) GetAll() *[]models.Product {

	return &r.Products
}

func (r *ProductRepositoryMock) Delete(ID uuid.UUID) error {
	_, index, err := r.FindByID(ID)

	if err != nil {
		return err
	}

	r.Products = remove(r.Products, index)

	return nil
}
func remove(products []models.Product, s int) []models.Product {
	return append(products[:s], products[s+1:]...)
}
