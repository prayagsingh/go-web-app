package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate for rendering the template using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {

	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing teplate", err)
		return
	}
}
