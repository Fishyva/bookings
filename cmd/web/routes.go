package main

import (
	"net/http"

	"github.com/Fishyva/bookings/pkg/config"
	"github.com/Fishyva/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
//Routes is basically the paths for your site 
func routes(app *config.AppConfig) http.Handler{

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	mux.Get("/generals-quarters",handlers.Repo.Generals)
	mux.Get("/majors-suite",handlers.Repo.Majors)

	mux.Get("/availability",handlers.Repo.Availability)
	mux.Post("/availability", handlers.Repo.PostAvailability)
	mux.Post("/availability-json", handlers.Repo.AvailabilityJSON)
	


	mux.Get("/contact",handlers.Repo.Contact)

	mux.Get("/make-reservation",handlers.Repo.Reservations)
	
	// for go to allow html to find static things
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))

	return mux
}