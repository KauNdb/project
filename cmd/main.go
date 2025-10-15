package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/db"
	"project/internal/auth"
	"project/internal/product"
	"project/pkg/middleware"
)

func main() {
	router := http.NewServeMux()
	cfg := configs.LoadConfig()
	db := db.NewDb(cfg)
	productRepository := product.NewProductRepository(db)

	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config: cfg,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Loggin(router),
	}

	fmt.Println("Server is listening on port 8081...")
	server.ListenAndServe()
}
