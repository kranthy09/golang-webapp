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

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf("Starting server at portNumber: %v", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	srv.ListenAndServe()
}
