package requests

import "github.com/google/uuid"

type AddProductsToCartRequest struct {
	ProductsIDs []uuid.UUID `json:"products_ids"`
}
