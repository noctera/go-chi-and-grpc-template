package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"my-app/internal/pkg/models"
	"my-app/internal/pkg/services"
	"net/http"
)

type CreateTestRequest struct {
	Name string `json:"name"`
}

func CreateTestHandler(w http.ResponseWriter, r *http.Request) {

	var req CreateTestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	// You can add more validation for JSON fields if necessary

	test := models.Test{
		Name: req.Name,
	}

	id, err := services.CreateTest(test)
	if err != nil {
		http.Error(w, "Failed to create test", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"id": id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func DeleteTestByIdHandler(w http.ResponseWriter, r *http.Request) {
	testId := chi.URLParam(r, "id")

	err := services.DeleteTestById(testId)
	if err != nil {
		http.Error(w, "Failed delete test", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
