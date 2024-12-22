package main

import (
	"fmt"
	"github.com/Darya-Tolmeneva/Sprint1_Yandex_Lyceum/internal/application"
)

func main() {
	app := application.New()
	fmt.Println("Run Server")
	err := app.RunServer()
	if err != nil {
		return
	}
}
