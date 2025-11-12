package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	resp := APIResponse{
		Status:  "error",
		Message: message,
	}
	RespondJSON(w, status, resp)
}

func RespondSuccess(w http.ResponseWriter, message string, data interface{}) {
	resp := APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	RespondJSON(w, http.StatusOK, resp)
}
