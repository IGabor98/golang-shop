package requests

import "github.com/google/uuid"

type AddproductsToCartRequest struct {
	ProductsIDs []uuid.UUID `json:"products_ids"`
}
