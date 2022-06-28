package requests

import "github.com/google/uuid"

type CreateCartRequest struct {
	ProductsIDs []uuid.UUID `json:"products_ids"`
}
