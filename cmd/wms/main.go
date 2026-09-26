package main

import (
	"log"
	"net/http"

	"github.com/vitorlfaria/go-wms/internal/domain"
	"github.com/vitorlfaria/go-wms/internal/infra/database"
	"github.com/vitorlfaria/go-wms/internal/infra/web"
)

func main() {
	prp := database.NewMemoryProductRepository()
	prp.Products["123"] = &domain.Product{ID: "123", Name: "Macbook", Quantity: 10}

	mux := http.NewServeMux()

	web.RegisterProductRoutes(mux, prp)

	log.Println("Api running on port 8080!")
	http.ListenAndServe(":8080", mux)
}
