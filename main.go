package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		// fmt.Print writes to console while Fprint writes to browser
		n, err := fmt.Fprintf(rw, "Hello World !!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))

	})

	_ = http.ListenAndServe(portNumber, nil)
}
