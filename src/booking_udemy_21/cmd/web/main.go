package main

import (
	"fmt"
	"net/http"

	"github.com/bny64/go-course/pkg/handlers"
)

const portNumber = ":8080"

// main is the application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
