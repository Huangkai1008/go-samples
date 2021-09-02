package domain

type Product struct {
	Name       string
	UnitPrice  float64
	IsFeatured bool
}

func (p *Product) ApplyDiscountFor(user UserContext) *DiscountedProduct {
	discount := 1.0
	if user.IsInRole(PreferredCustomer) {
		discount = 0.95
	}
	return &DiscountedProduct{
		name:      p.Name,
		unitPrice: p.UnitPrice * discount,
	}
}

type DiscountedProduct struct {
	name      string
	unitPrice float64
}
