package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/db"
	"project/internal/auth"
	"project/internal/order"
	"project/internal/product"
)

func main() {
	router := http.NewServeMux()
	cfg := configs.LoadConfig()
	db := db.NewDb(cfg)
	// Repository
	productRepository := product.NewProductRepository(db)
	authRepository := auth.NewAuthRepository(db)
	orderRepositopry := order.NewOrderRepository(db)

	// Handlers
	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config:         cfg,
		AuthRepository: authRepository,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepository,
		Config:            cfg,
	})
	order.NewOrderHandler(router, order.OrderHandlerDeps{
		OrderRepository: orderRepositopry,
		Config:          cfg,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081...")
	server.ListenAndServe()
}
