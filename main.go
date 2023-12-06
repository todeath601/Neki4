package main

import (
	"log"
	"net/http"

	"github.com/example/handlers"
)

func main() {
	http.HandleFunc("/", handlers.PersonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
