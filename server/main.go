package main

import (
	"net/http"

	"github.com/todeath601/Neki4/server/handlers/"
)

func main() {
	http.HandleFunc("/person", handlers.HandlePerson)
	http.ListenAndServe(":8080", nil)
}
