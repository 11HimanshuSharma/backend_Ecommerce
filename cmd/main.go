package main


import (
	"log"
	"net/http"
	"ecommerce/pkg/response"
)

func main(){
	//1 initialize go's stardard stream multiplexer 
	mux := http.NewServeMux()


	//2 . Define our Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		response.JSON(w, http.StatusOK, map[string]string{
			"Status": "OK",
			"message": "Stardard net/http server is running!"
		})
	})

	// start the server
	port := ":8080"
	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}