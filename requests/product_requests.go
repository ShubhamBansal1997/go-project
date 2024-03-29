package requests

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ProductFilter struct {
	Category string  `form:"category"`
	Name     string  `form:"name"`
	SKU      string  `form:"sku"`
	PriceGte float64 `form:"price__gte"`
	PriceLte float64 `form:"price__lte"`
	OrderBy  string  `form:"order_by"`
}

func (pf ProductFilter) Validate() error {
	return validation.ValidateStruct(&pf,
		validation.Field(&pf.Category, validation.When(pf.Category != "", validation.Required).Else(validation.Length(1, 100))),
		validation.Field(&pf.Name, validation.When(pf.Name != "", validation.Required).Else(validation.Length(1, 100))),
		validation.Field(&pf.SKU, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+$"))),
		validation.Field(&pf.PriceGte),
		validation.Field(&pf.PriceLte),
		validation.Field(&pf.OrderBy, validation.By(orderByRule)),
	)
}

func orderByRule(value interface{}) error {
	orderBy, _ := value.(string) // Type assertion; safe because we know it's a string
	if orderBy != "" && !strings.HasSuffix(orderBy, " desc") && !strings.HasSuffix(orderBy, " asc") {
		return validation.NewError("validation_order_by", "OrderBy must end with 'desc' or 'asc'")
	}
	return nil
}

type FilterProductRequest struct {
	ProductFilter
}
