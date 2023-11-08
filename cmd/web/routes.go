package main

import (
	"github.com/bmizerany/pat"
	"github.com/cjwood032/go-course/pkg/config"
	"github.com/cjwood032/go-course/pkg/handlers"
	"net/http"
)

func routes(app *config.ApplConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
