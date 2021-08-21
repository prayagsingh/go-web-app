package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate for rendering the template using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(rw)
	if err != nil {
		fmt.Println("Error getting template Cache", err)
		return
	}

	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parseTemplate.Execute(rw, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}

// RenderTemplateTest for template caching
func RenderTemplateTest(rw http.ResponseWriter) (map[string]*template.Template, error) {

	// store the template found during parsing
	mycache := map[string]*template.Template{}

	// find all the pages in template dir which ends with page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently ", page)
		// New allocates the new template with a given name
		// Funcs adds the elements of the argument map to the template's function map. It must be
		// 	 called before the template is parsed.
		// ParseFiles parses the named files and associates the resulting templates with t. If an
		//   error occurs, parsing stops and the returned template is nil; otherwise it is t
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return mycache, err
		}

		if len(matches) > 0 {
			ts, err = template.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return mycache, err
			}
		}

		// Adding templates to cache
		mycache[name] = ts
	}

	return mycache, nil
}
