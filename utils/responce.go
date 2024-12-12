package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	success := false
	if status < 400 {
		success = true
	}
	res := map[string]interface{}{
		"status":  status,
		"success": success,
		"data":    payload,
	}
	json.NewEncoder(w).Encode(res)
}
