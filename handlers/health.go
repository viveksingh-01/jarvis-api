package handlers

import (
	"encoding/json"
	"net/http"
)

func CheckServerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Server's health is OK.",
	})
}
