package main

import (
	"net/http"
	"persona/handlers"
)

func main() {
	http.HandleFunc("/person", handlers.HandlePerson)
	http.ListenAndServe(":8080", nil)
}
