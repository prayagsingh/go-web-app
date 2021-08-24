package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf add CSRF protection to all the POST requests
func NoSurf(next http.Handler) http.Handler {

	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		// set it to true when in production
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
