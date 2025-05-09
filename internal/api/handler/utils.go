package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// encode wraps the response creation process
func encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Error encoding response: %w", err)
	}
	return nil
}

// decode wraps the data retrieval process from the request
func decode[T any](r *http.Request) (*T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return &v, fmt.Errorf("Error decoding request: %w", err)
	}
	return &v, nil
}
