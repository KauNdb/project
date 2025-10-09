package main

import (
	"fmt"
	"net/http"
	"project/configs"
	"project/db"
	"project/internal/auth"
)

func main() {
	router := http.NewServeMux()
	cfg := configs.LoadConfig()
	_ = db.NewDb(cfg)
	auth.NewHandler(router, auth.AuthHandlerDeps{
		Config: cfg,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081...")
	server.ListenAndServe()
}
