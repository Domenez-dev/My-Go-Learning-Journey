package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PORT_NUMBER = ":8080"

func Home(context string) http.HandlerFunc {
	_ = context
	return func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "home.page.tmpl")
	}
}

func RenderTemplate(w http.ResponseWriter, tmp string) {
	ParsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := ParsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", Home("Hello World"))
	http.HandleFunc("/about", Home("About"))

	fmt.Println("Listening on Port ", PORT_NUMBER)
	err := http.ListenAndServe(PORT_NUMBER, nil)
	if err != nil {
		fmt.Println(err)
	}
}
