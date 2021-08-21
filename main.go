package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is about home page hadler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "This is the home page")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprint(rw, fmt.Sprintf("This is the about page and sum is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
