package web

import (
	"net/http"

	"github.com/vitorlfaria/go-wms/internal/domain"
	"github.com/vitorlfaria/go-wms/internal/usecase"
)

func RegisterProductRoutes(mux *http.ServeMux, repo domain.ProductRepository) {
	reserveUC := usecase.NewReserveStockUseCase(repo)
	handler := NewProductHandler(reserveUC)

	mux.HandleFunc("POST api/v1/products/reserve", handler.ReserveStock)
}
