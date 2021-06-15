package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kranthy09/go-course/pkg/config"
	"github.com/kranthy09/go-course/pkg/handlers"
	"github.com/kranthy09/go-course/pkg/render"
)

var portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.Inproduction = false

	// create a session
	session = scs.New()

	// decides the lifetime of the session
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	// if we close the page or website, the session saves the cookies to samesite
	session.Cookie.SameSite = http.SameSiteLaxMode

	// Secure the stored Cookies
	session.Cookie.Secure = app.Inproduction

	// assign the session to application config
	app.Session = session

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
