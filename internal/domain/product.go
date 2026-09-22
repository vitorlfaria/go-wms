package domain

import (
	"context"
)

type Product struct {
	ID       string
	Name     string
	SKU      string
	Quantity int
}

func (p *Product) DecreaseStock(amount int) error {
	if p.Quantity-amount < 0 {
		return ErrProductStockNegative
	}

	p.Quantity -= amount
	return nil
}

type ProductRepository interface {
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, product *Product) error
}
