package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vitorlfaria/go-wms/internal/domain"
	"github.com/vitorlfaria/go-wms/internal/usecase"
)

type ReserveRequest struct {
	ProductID string `json:"product_id"`
	Amount    int    `json:"amount"`
}

type ProductHandler struct {
	ReserveStockUC *usecase.ReserveStockUseCase
}

func NewProductHandler(uc *usecase.ReserveStockUseCase) *ProductHandler {
	return &ProductHandler{
		ReserveStockUC: uc,
	}
}

func (h *ProductHandler) ReserveStock(w http.ResponseWriter, r *http.Request) {
	var req ReserveRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.ReserveStockUC.Execute(r.Context(), req.ProductID, req.Amount); err != nil {
		if errors.Is(err, domain.ErrProductNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if errors.Is(err, domain.ErrProductStockNegative) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
