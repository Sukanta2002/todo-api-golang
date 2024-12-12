package utils

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
)

func ApiError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res := map[string]interface{}{
		"status":   status,
		"success":  false,
		"data":     nil,
		"messsage": message,
		"stack":    string(debug.Stack()),
	}
	json.NewEncoder(w).Encode(res)
}
