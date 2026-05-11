package main

import (
	"ecommerce/internal/repository"
	"ecommerce/pkg/response"
	"ecommerce/internal/service"
	"ecommerce/internal/handler"
	"ecommerce/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	//1 initialize go's stardard stream multiplexer

	// updating cred to match local postgres
	dbCfg := database.Config{
		Host: "localhost",
		Port: 5432,
		User: "postgres",
		Password: "password",
		DBName: "ecommerce",
		SSLMode: "disable",
	}

	db,err := database.Connect(dbCfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close() // make sure we close connection when main exits

	// data
	// repo := repository.NewInMemoryProductRepo()
	repo := repository.NewPostgresProductRepo(db)

	//service
	svc := service.NewProductservice(repo)

	// handler
	h := handler.NewProductHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /products", h.CreateProduct)
	mux.HandleFunc("GET /products", h.ListProducts)
	mux.HandleFunc("GET /products/{id}", h.GetProduct)

	//2 . Define our Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"Status":  "OK",
			"message": "Stardard net/http server is running!",
		})
	})

	// apply middleware
	//ordres matters! outermost middleware runs first on the way IN.
	// Request comes in: RequestID -> recovery -> logging -> handler
	// response goes out : handler-> logging -> recovery -> requestId

	var handler http.Handler = mux
	handler = middleware.Logging(handler)
	handler = middleware.Recovery(handler)
	handler = middleware.RequestID(handler)
	// start the server
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
