package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"shop/repositories"
	"shop/requests"
)

type CartController struct {
	CartRepository repositories.CartRepository
}

func (c *CartController) CreateCart(w http.ResponseWriter, req *http.Request) {
	createCartRequest := &requests.CreateCartRequest{}

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(createCartRequest); err != nil {
		panic(err)
	}

	cart := c.CartRepository.CreateCart(createCartRequest.ProductsIDs)

	json.NewEncoder(w).Encode(cart)
}

func (c *CartController) GetCart(w http.ResponseWriter, req *http.Request) {
	cartID, err := uuid.Parse(chi.URLParam(req, "cartID"))

	if err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	cart, err := c.CartRepository.FindCartByID(cartID)

	if err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	json.NewEncoder(w).Encode(cart)
}

func (c *CartController) AddProductsToCart(w http.ResponseWriter, req *http.Request) {
	cartID, err := uuid.Parse(chi.URLParam(req, "cartID"))

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	addProductsToCartRequest := &requests.AddproductsToCartRequest{}

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(addProductsToCartRequest); err != nil {
		panic(err)
	}

	err = c.CartRepository.AddProducts(cartID, addProductsToCartRequest.ProductsIDs)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	json.NewEncoder(w).Encode("Products added")
}
