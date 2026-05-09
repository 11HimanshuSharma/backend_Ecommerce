package response



import (
	"encoding/json"
	"net/http"
)


func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set(
	   "Content-Type", 
	   "application/json",
	)
	w.WriteHeader(status)


	// convert the go map/struct into jsonwrite it directly to the response stream
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, 
		"failed to encode response",
		http.StatusInternalServerError)
	}
}

// Error sends a standardized JSON error response

func Error(w http.ResponseWriter, status int, message string){
	JSON(w, status, map[string]string{"error": message})
}