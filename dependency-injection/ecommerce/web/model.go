package web

import (
	"fmt"

	"go-samples/dependency-injection/ecommerce/domain"
)

type ProductViewModel struct {
	Name        string
	UnitPrice   float64
	SummaryText string
}

func NewProductViewModel(product *domain.DiscountedProduct) *ProductViewModel {
	summaryText := fmt.Sprintf("%s (%.2f)", product.Name, product.UnitPrice)
	return &ProductViewModel{
		Name:        product.Name,
		UnitPrice:   product.UnitPrice,
		SummaryText: summaryText,
	}
}
