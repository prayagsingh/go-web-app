package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is about home page hadler
func Home(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "home.page.html")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "about.page.html")
}

// renderTemplate for rendering the template
func renderTemplate(rw http.ResponseWriter, tmpl string) {

	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing teplate", err)
		return
	}

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
