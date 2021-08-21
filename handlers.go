package main

import (
	"net/http"
)

// Home is about home page hadler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.html")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.html")
}
