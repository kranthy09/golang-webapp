package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Nosurf adds CSRF protection to all POST requests
func Nosurf(next http.Handler) http.Handler {
	csrfToken := nosurf.New(next)

	csrfToken.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",              //entire website
		Secure:   app.Inproduction, // set to production bool value
		SameSite: http.SameSiteLaxMode,
	})

	return csrfToken
}

// 	SessionLoad loads and saves session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
