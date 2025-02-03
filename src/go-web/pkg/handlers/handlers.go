package handlers

import (
	"net/http"

	"go-web/pkg/render"
)

func Home(context string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(w, context)
	}
}
