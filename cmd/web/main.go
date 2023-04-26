package main

import (
	"fmt"
	"github.com/gandra/helloworld-web-go-trevor-sawler/config"
	"github.com/gandra/helloworld-web-go-trevor-sawler/pkg/handlers"
	"github.com/gandra/helloworld-web-go-trevor-sawler/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache:", err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
