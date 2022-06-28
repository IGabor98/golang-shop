package controllers

import (
	"encoding/json"
	"net/http"
	"shop/models"
	"shop/repositories"
)

type ProductController struct {
	ProductRepository repositories.ProductRepository
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, req *http.Request) {
	product := &models.Product{}

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(product); err != nil {
		panic(err)
	}

	createdProduct, err := c.ProductRepository.Create(*product)

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(createdProduct)
}

func (c *ProductController) GetAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(c.ProductRepository.GetAll())
}
