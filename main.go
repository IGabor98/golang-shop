package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shop/controllers"
	"shop/repositories"
)

func main() {
	productRepository := &repositories.ProductRepositoryMock{}
	cartRepository := &repositories.CartRepositoryMock{
		ProductRepository: productRepository,
	}

	productController := &controllers.ProductController{
		ProductRepository: productRepository,
	}
	cartController := &controllers.CartController{
		CartRepository: cartRepository,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Post("/products", productController.CreateProduct)
	r.Get("/products/all", productController.GetAll)

	r.Post("/carts/create", cartController.CreateCart)
	r.Get("/carts/{cartID}", cartController.GetCart)
	r.Put("/carts/{cartID}/add", cartController.AddProductsToCart)

	http.ListenAndServe(":8090", r)
}
