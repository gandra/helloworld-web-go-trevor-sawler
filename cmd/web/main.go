package main

import (
	"fmt"
	"github.com/gandra/helloworld-web-go-trevor-sawler/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
