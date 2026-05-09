package response



import (
	"encoding/json"
	"net/http"
	"ecommerce/pkg/apperror"
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


func HandleError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*apperror.AppError); ok {
		JSON(w, appErr.HTTPStatus(), map[string] string{
			"code": appErr.Code,
			"message": appErr.Message,
		})
		return
	}
	JSON(w, http.StatusInternalServerError, map[string]string{
		"code": "INTERNAL_ERROR",
		"message": "An unexpected error occurred",
	})
}
