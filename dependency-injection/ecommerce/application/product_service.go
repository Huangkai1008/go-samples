package application

import "go-samples/dependency-injection/ecommerce/domain"

type IProductService interface {
	GetFeaturedProducts() []*domain.DiscountedProduct
}

type ProductService struct {
	repository  domain.ProductRepository
	userContext domain.UserContext
}

func (s *ProductService) GetFeaturedProducts() []*domain.DiscountedProduct {
	var discountProducts []*domain.DiscountedProduct
	for _, p := range s.repository.GetFeaturedProducts() {
		discountProducts = append(discountProducts, p.ApplyDiscountFor(s.userContext))
	}
	return discountProducts
}
