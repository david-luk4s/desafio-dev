package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/david-luk4s/desafio-dev/application"
)

type PayResponse struct {
	Message string `json:"message"`
}

func upload(w http.ResponseWriter, r *http.Request) {
	payload := PayResponse{}
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		payload.Message = "not allowed method"
		json.NewEncoder(w).Encode(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myfile")
	if err != nil {
		payload.Message = err.Error()
		json.NewEncoder(w).Encode(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Process full transaction
	if err := application.ProcessTransaction(ctx, file); err != nil {
		payload.Message = err.Error()
		json.NewEncoder(w).Encode(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload.Message = "file processed"
	json.NewEncoder(w).Encode(payload)
	w.WriteHeader(http.StatusOK)
}
func list(w http.ResponseWriter, r *http.Request) {
	payload := PayResponse{}
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		payload.Message = "not allowed method"
		json.NewEncoder(w).Encode(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get full transaction
	items, err := application.ListTransaction(ctx)
	if err != nil {
		payload.Message = err.Error()
		json.NewEncoder(w).Encode(payload)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(items)
	w.WriteHeader(http.StatusOK)
}
