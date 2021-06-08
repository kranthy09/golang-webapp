package main

import (
	"fmt"
	"net/http"

	"github.com/kranthy09/go-course/pkg/handlers"
)

var portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/", handlers.About)
	fmt.Printf("Starting server at portNumber: %v", portNumber)
	http.ListenAndServe(portNumber, nil)
}
