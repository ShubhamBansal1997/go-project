package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BasicAddress struct {
	Value string `json:"value" validate:"required"`
}

func (ba BasicAddress) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Value, validation.Required),
	)
}

type CreateAddressRequest struct {
	BasicAddress
}

type UpdateAddressRequest struct {
	BasicAddress
}
