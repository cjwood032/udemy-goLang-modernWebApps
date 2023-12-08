package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/cjwood032/go-course/pkg/config"
	"github.com/cjwood032/go-course/pkg/handlers"
	"github.com/cjwood032/go-course/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":3000"

var app config.ApplConfig
var session *scs.SessionManager

func main() {

	app.UseCache = true
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
