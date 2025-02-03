package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmp string) {
	ParsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := ParsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
