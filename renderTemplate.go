package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate for rendering the template
func renderTemplate(rw http.ResponseWriter, tmpl string) {

	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing teplate", err)
		return
	}
}
