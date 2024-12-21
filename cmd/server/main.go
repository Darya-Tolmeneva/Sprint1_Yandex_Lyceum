package main

import (
	"Sprint1/pkg/api"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", api.Handler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
