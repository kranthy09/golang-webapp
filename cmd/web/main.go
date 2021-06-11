package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kranthy09/go-course/pkg/config"
	"github.com/kranthy09/go-course/pkg/handlers"
	"github.com/kranthy09/go-course/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot cache template in main.go", err)
	}

	fmt.Println("tc:", tc)

	app.TemplateCache = tc

	fmt.Println("app:", app)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting server at portNumber: %v", portNumber)
	http.ListenAndServe(portNumber, nil)
}
