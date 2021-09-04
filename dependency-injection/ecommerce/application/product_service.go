package application

import "go-samples/dependency-injection/ecommerce/domain"

type IProductService interface {
	GetFeaturedProducts() ([]*domain.DiscountedProduct, error)
}

type ProductService struct {
	repository  domain.ProductRepository
	userContext domain.UserContext
}

func (s *ProductService) GetFeaturedProducts() ([]*domain.DiscountedProduct, error) {
	var discountProducts []*domain.DiscountedProduct
	featuredProducts, err := s.repository.GetFeaturedProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range featuredProducts {
		discountProducts = append(discountProducts, p.ApplyDiscountFor(s.userContext))
	}
	return discountProducts, nil
}
