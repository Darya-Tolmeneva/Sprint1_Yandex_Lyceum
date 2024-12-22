package application

import (
	"bytes"
	"encoding/json"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/models"
	"net/http"
	"net/http/httptest"
	"strings"
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
		{"GET", "1+2", http.StatusMethodNotAllowed, "Method not allowed"},
		{"POST", "", http.StatusUnprocessableEntity, "Expression is not valid"},
		{"POST", `{"expression": 123}`, http.StatusUnprocessableEntity, "Expression is not valid"},
		{"POST", "1" + strings.Repeat("+1", 1000), http.StatusOK, "1001.00"},
		{"POST", "1+2a", http.StatusUnprocessableEntity, "Expression is not valid"},
	}
	for _, test := range tests {
		t.Run(test.expression, func(t *testing.T) {
			reqBody, _ := json.Marshal(models.Request{Expression: test.expression})
			req := httptest.NewRequest(test.method, "/api/v1/calculate", bytes.NewBuffer(reqBody))
			rr := httptest.NewRecorder()

			Handler(rr, req)

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
