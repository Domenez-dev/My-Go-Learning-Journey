package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"go-web/pkg/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home("home.html")))

	return mux
}
