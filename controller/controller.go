package controller

import (
	"encoding/json"
	m "mini-apis/models"
	s "mini-apis/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CardValidator(w http.ResponseWriter, r *http.Request) {
	var number m.Card
	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&number)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the user request using go-playground/validator
	err = validate.Struct(number)
	if err != nil {
		http.Error(w, "Validation failed for card no", http.StatusBadRequest)
		return
	}

	result := s.ValidateCard(number.Number)

	// Prepare the response
	response := map[string]bool{"valid": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
