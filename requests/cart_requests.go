package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CartItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"min=0"`
}

func (bp CartItem) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.ProductID, validation.Required),
		validation.Field(&bp.Quantity),
	)
}

type UpdateCartItemRequest struct {
	CartItem
}
