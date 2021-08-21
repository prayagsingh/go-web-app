package handlers

import (
	"net/http"

	"github.com/prayagsingh/go-web-app/pkg/render"
)

// Home is about home page hadler
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.html")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.html")
}
