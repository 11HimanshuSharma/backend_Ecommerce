package main

import (
	"ecommerce/internal/handler"
	"ecommerce/internal/repository"
	"ecommerce/internal/service"
	"ecommerce/pkg/response"
	"log"
	"net/http"
)

func main() {
	//1 initialize go's stardard stream multiplexer

	// data
	repo := repository.NewInMemoryProductRepo()

	//service
	svc := service.NewProductService(repo)

	// handler
	h := handler.NewProductHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("GET /products", h.ListProducts)
	mux.HandleFunc("GET /products/{id}", h.GetProduct)

	//2 . Define our Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"Status":  "OK",
			"message": "Stardard net/http server is running!",
		})
	})

	// start the server
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
