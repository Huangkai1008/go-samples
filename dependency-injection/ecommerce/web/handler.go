package web

import (
	"go-samples/dependency-injection/ecommerce/application"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	*application.ProductService
}

func (h *HomeHandler) GetFeaturedProducts(c echo.Context) error {
	products, err := h.ProductService.GetFeaturedProducts()
	if err != nil {
		return err
	}

	var viewModels []*ProductViewModel
	for _, p := range products {
		viewModels = append(viewModels, NewProductViewModel(p))
	}

	return c.JSON(200, viewModels)
}
