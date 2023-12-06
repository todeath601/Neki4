package main

import (
	"net/http"

	"github.com/username/project/handlers"
)

func main() {
	http.HandleFunc("/person", handlers.HandlePerson)
	http.ListenAndServe(":8080", nil)
}
