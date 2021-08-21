package handlers

import (
	"net/http"

	"github.com/prayagsingh/go-web-app/pkg/render"
)

// Home is the handler for the home page
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.html")
}

// About is the handler for the about page
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.html")
}
