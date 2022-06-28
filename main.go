package main

import (
	"encoding/json"
	"net/http"
	"shop/models"
	"shop/repositories"
	"shop/requests"
)

type Server struct {
	productRepository repositories.ProductRepository
	cartRepository    repositories.CartRepository
}

func main() {
	productRepository := &repositories.ProductRepositoryMock{}
	cartRepository := &repositories.CartRepositoryMock{
		ProductRepository: productRepository,
	}
	server := Server{
		productRepository: productRepository,
		cartRepository:    cartRepository,
	}

	http.HandleFunc("/products", server.createProduct)
	http.HandleFunc("/products/all", server.GetAll)

	http.HandleFunc("/carts/create", server.CreateCart)

	http.ListenAndServe(":8090", nil)
}

func (s *Server) createProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		product := &models.Product{}

		defer req.Body.Close()
		if err := json.NewDecoder(req.Body).Decode(product); err != nil {
			panic(err)
		}

		createdProduct, err := s.productRepository.Create(*product)

		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(createdProduct)
	}
}

func (s *Server) GetAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.productRepository.GetAll())
}

func (s *Server) CreateCart(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		createCartRequest := &requests.CreateCartRequest{}

		defer req.Body.Close()
		if err := json.NewDecoder(req.Body).Decode(createCartRequest); err != nil {
			panic(err)
		}

		cart := s.cartRepository.CreateCart(createCartRequest.ProductsIDs)

		json.NewEncoder(w).Encode(cart)
	}
}
