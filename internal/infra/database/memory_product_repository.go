package database

import (
	"context"
	"sync"

	"github.com/vitorlfaria/go-wms/internal/domain"
)

type MemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		mu:       sync.RWMutex{},
		products: make(map[string]*domain.Product),
	}
}

func (pr *MemoryProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	pr.mu.RLock()
	defer pr.mu.RUnlock()

	product := pr.products[id]
	if product == nil {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (pr *MemoryProductRepository) Update(ctx context.Context, product *domain.Product) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	p := pr.products[product.ID]
	if p == nil {
		return domain.ErrProductNotFound
	}

	pr.products[product.ID] = product
	return nil
}
