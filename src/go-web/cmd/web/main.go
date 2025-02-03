package main

import (
	"fmt"
	"net/http"

	"go-web/pkg/handlers"
)

const PORT_NUMBER = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home("home.html"))
	// http.HandleFunc("/about", Home("About"))

	fmt.Println("Listening on Port ", PORT_NUMBER)
	err := http.ListenAndServe(PORT_NUMBER, nil)
	if err != nil {
		fmt.Println(err)
	}
}
