package tests

import (
	"Sprint1/internal/models"
	"Sprint1/pkg/api"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		method     string
		expression string
		statusCode int
		result     string
	}{
		{"POST", "1+2", http.StatusOK, "3.00"},
		{"POST", "1+2*", http.StatusUnprocessableEntity, "Expression is not valid"},
		{"POST", "1/0", http.StatusInternalServerError, "Internal server error"},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			reqBody, _ := json.Marshal(models.Request{Expression: test.expression})
			req := httptest.NewRequest(test.method, "/api/v1/calculate", bytes.NewBuffer(reqBody))
			rr := httptest.NewRecorder()

			api.Handler(rr, req)

			res := rr.Result()
			defer res.Body.Close()

			if res.StatusCode != test.statusCode {
				t.Errorf("expected status %v, got %v", test.statusCode, res.StatusCode)
			}

			var result models.Response
			json.NewDecoder(res.Body).Decode(&result)

			if test.statusCode == http.StatusOK && result.Result != test.result {
				t.Errorf("expected result %v, got %v", test.result, result.Result)
			}
			if test.statusCode != http.StatusOK && result.Error != test.result {
				t.Errorf("expected error %v, got %v", test.result, result.Error)
			}

		})
	}
}
