package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate for rendering the template using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {

	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		return
	}

	// ok is used to check if the template exists or not
	// if template found, ok wil return true else false
	t, ok := templateCache[tmpl]

	//fmt.Println("t is ", t.DefinedTemplates())

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		fmt.Println("Error writing parsed template to buffer", err)
		return
	}
	_, err = buf.WriteTo(rw)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// store the template found during parsing
	myCache := map[string]*template.Template{}

	// find all the pages in template dir which ends with page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Printf("Page is currently %s and name is %s \n", page, name)
		// New allocates the new template with a given name
		// Funcs adds the elements of the argument map to the template's function map. It must be
		// 	 called before the template is parsed.
		// ParseFiles parses the named files and associates the resulting templates with t. If an
		//   error occurs, parsing stops and the returned template is nil otherwise it is t
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// Adding templates to cache
		myCache[name] = ts
		//fmt.Println("Defined templates stored in myCache is ", myCache[name].DefinedTemplates())
		//fmt.Println("template stored in myCache is ", myCache[name].Tree.Name)
	}

	return myCache, nil
}
