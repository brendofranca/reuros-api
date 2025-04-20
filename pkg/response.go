package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
