package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prayagsingh/go-web-app/pkg/config"
	"github.com/prayagsingh/go-web-app/pkg/handlers"
	"github.com/prayagsingh/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}

	app.TemplateCache = tc
	// set to false if you want to load template from disk in development mode
	// if it is set to true then any changes made to templates won't reflect dynamically because
	// it is reading from template cache instead of disk. it is faster in loading when comparing it
	// to reading it from disk
	app.UseCache = false

	// This allow Handler functions to have access to appConfig via repository
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
