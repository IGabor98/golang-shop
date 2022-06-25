package main

import (
	"encoding/json"
	"net/http"
	"shop/models"
	"shop/repositories"
)

type Server struct {
	productRepository repositories.ProductRepository
}

func main() {
	productRepository := repositories.ProductRepositoryMock{}
	server := Server{
		productRepository: &productRepository,
	}

	http.HandleFunc("/products", server.createProduct)
	http.HandleFunc("/products/all", server.GetAll)

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
