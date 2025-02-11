package handlers

import (
	"net/http"

	"go-web/pkg/render"
)

func Home(context string) http.HandlerFunc {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello web!"

	return func(w http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(w, context, &render.TemplateData{
			StringMap: stringMap,
		})
	}
}
