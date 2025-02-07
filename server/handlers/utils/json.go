package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending JSON response, %v", err), http.StatusInternalServerError)
		return
	}
}
