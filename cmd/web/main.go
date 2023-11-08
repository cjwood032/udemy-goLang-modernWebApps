package main

import (
	"fmt"
	"github.com/cjwood032/go-course/pkg/config"
	"github.com/cjwood032/go-course/pkg/handlers"
	"github.com/cjwood032/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":3000"

func main() {
	var app config.ApplConfig
	app.UseCache = true
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
