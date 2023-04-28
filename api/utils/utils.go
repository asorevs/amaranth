package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err RestErr) {
	RespondJson(w, err.Status, err)
}
