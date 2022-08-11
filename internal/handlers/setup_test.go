package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Fishyva/bookings/internal/config"
	"github.com/Fishyva/bookings/internal/models"
	"github.com/Fishyva/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func GetRoutes() http.Handler {

	//What will store in the session
	gob.Register(models.Reservation{})

	//App Config inproduction mode
	app.InProduction = false

	// creates a session cookie
	session = scs.New()
	// how long it lasts
	session.Lifetime = 24 * time.Hour
	// if you want it to last even after they close their browser
	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode
	// to see if its HTTPS
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplate(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/availability", Repo.Availability)
	mux.Post("/availability", Repo.PostAvailability)
	mux.Post("/availability-json", Repo.AvailabilityJSON)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservations)
	mux.Post("/make-reservation", Repo.PostReservations)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	// for go to allow html to find static things
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
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

func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	//step 1 need to get the files at a location
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	// getting the indivisual pages from templates folder
	for _, page := range pages {
		name := filepath.Base(page)

		hs, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// see if your layout matches
		matches, err := filepath.Glob(fmt.Sprintf("%s/*layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			hs, err = hs.ParseGlob(fmt.Sprintf("%s/*layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = hs

	}
	return myCache, nil
}
