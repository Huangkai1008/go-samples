package repository

import (
	"go-samples/dependency-injection/ecommerce/domain"

	"github.com/jmoiron/sqlx"
)

type SQLProductRepository struct {
	db *sqlx.DB
}

func (r *SQLProductRepository) GetFeaturedProducts() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Select(&products, "SELECT * FROM products WHERE is_featured = true")
	if err != nil {
		return nil, err
	}
	return products, nil
}
