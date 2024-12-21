package api

import (
	"Sprint1/internal/calculator"
	"Sprint1/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var request models.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusUnprocessableEntity)
		return
	}
	result, err := calculator.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, models.ErrInvalidSymbols) || errors.Is(err, models.ErrInvalidParenthesis) || errors.Is(err, models.ErrInvalidOperations) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(models.Response{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.Response{Error: "Internal server error"})
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{Result: fmt.Sprintf("%.2f", result)})
}
