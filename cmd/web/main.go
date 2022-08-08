package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Fishyva/bookings/internal/config"
	"github.com/Fishyva/bookings/internal/handlers"
	"github.com/Fishyva/bookings/internal/models"
	"github.com/Fishyva/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// Home is the homepage handler
var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on port:%s", portNumber))

	//the underscore heremeans you dont care about the error
	//  _ = http.ListenAndServe(portNumber,nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}

func run() error {
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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// this is listening to the web to send something through its port

	// http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	return nil
}
