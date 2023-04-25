package handlers

import (
	"github.com/gandra/helloworld-web-go-trevor-sawler/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
