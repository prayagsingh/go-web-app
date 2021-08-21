package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is about home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "This is the home page")
}

// About is "about us" page handler
func About(rw http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprint(rw, fmt.Sprintf("This is the about page and sum is %d", sum))
}

func Divide(rw http.ResponseWriter, r *http.Request) {
	divide, err := divideValues(10.0, 5.0)
	if err != nil {
		fmt.Fprintf(rw, "can't divide by zero")
		return
	}
	fmt.Fprintf(rw, fmt.Sprintf("This is the divide func and value is %f", divide))
}

func addValues(x, y int) int {

	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("y should be greator than zero")
		return 0, err
	}
	return x / y, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
