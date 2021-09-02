package domain

type IProductService interface {
	GetFeaturedProducts() []*DiscountedProduct
}

type ProductService struct {
	repository  ProductRepository
	userContext UserContext
}

func (s *ProductService) GetFeaturedProducts() []*DiscountedProduct {
	var discountProducts []*DiscountedProduct
	for _, p := range s.repository.GetFeaturedProducts() {
		discountProducts = append(discountProducts, p.ApplyDiscountFor(s.userContext))
	}
	return discountProducts
}
