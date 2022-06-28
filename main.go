package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/products", server.createProduct)
	r.Get("/products/all", server.GetAll)

	r.Post("/carts/create", server.CreateCart)
	r.Get("/carts/{cartID}", server.GetCart)

	http.ListenAndServe(":8090", r)
}

func (s *Server) createProduct(w http.ResponseWriter, req *http.Request) {
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

func (s *Server) GetAll(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.productRepository.GetAll())
}

func (s *Server) CreateCart(w http.ResponseWriter, req *http.Request) {
	createCartRequest := &requests.CreateCartRequest{}

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(createCartRequest); err != nil {
		panic(err)
	}

	cart := s.cartRepository.CreateCart(createCartRequest.ProductsIDs)

	json.NewEncoder(w).Encode(cart)

}

func (s *Server) GetCart(w http.ResponseWriter, req *http.Request) {
	cartID, err := uuid.Parse(chi.URLParam(req, "cartID"))

	if err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	cart, err := s.cartRepository.FindCartByID(cartID)

	if err != nil {
		json.NewEncoder(w).Encode(err)

		return
	}

	json.NewEncoder(w).Encode(cart)
}
