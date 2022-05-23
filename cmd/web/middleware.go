package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return crsfHandler
}
// SessionLoad loads and saves the session on every request 
func SessionLoad(next http.Handler) http.Handler {
	//LoadAndSave provides middleware which automatically loads and saves session data 
	//for the current request, and communicates the session token to and from the client in a cookie.
	return session.LoadAndSave(next)
}