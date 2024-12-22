package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/calculator"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/pkg/models"
	"net/http"
	"os"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
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

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", Handler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"Not Found"}`, http.StatusNotFound)
	})

	return http.ListenAndServe(":"+a.config.Addr, nil)
}
