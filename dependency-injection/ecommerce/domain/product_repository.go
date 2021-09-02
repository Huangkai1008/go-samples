package domain

type ProductRepository interface {
	GetFeaturedProducts() []*Product
}
