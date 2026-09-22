package usecase

import (
	"context"

	"github.com/vitorlfaria/go-wms/internal/domain"
)

type ReserveStockUseCase struct {
	productRepository domain.ProductRepository
}

func NewReserveStockUseCase(productRepo domain.ProductRepository) *ReserveStockUseCase {
	return &ReserveStockUseCase{productRepository: productRepo}
}

func (r *ReserveStockUseCase) Execute(ctx context.Context, productID string, amount int) error {
	product, err := r.productRepository.GetByID(ctx, productID)
	if err != nil {
		return err
	}

	if err := product.DecreaseStock(amount); err != nil {
		return err
	}

	if err := r.productRepository.Update(ctx, product); err != nil {
		return err
	}

	return nil
}
