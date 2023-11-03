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
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}
	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting Application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
