package response

import (
	"encoding/json"
	"net/http"
	"product-api-postgres/internal/models"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string         `json:"message"`
	Product models.Product `json:"product"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{
		Error: message,
	}
	WriteJSON(w, status, errorResponse)

}

func WriteSucces(w http.ResponseWriter, status int, message string, product models.Product) {
	successResponse := SuccessResponse{
		Message: message,
		Product: product,
	}
	WriteJSON(w, status, successResponse)
}
