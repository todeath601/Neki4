package handlers

import (
 "encoding/json"
 "net/http"
)

type Person struct {
 Name     string json:"name"
 Surname  string json:"surname"
 Age      int    json:"age"
}

func HandlePerson(w http.ResponseWriter, r *http.Request) {
 if r.Method == "POST" {
  var person Person
  err := json.NewDecoder(r.Body).Decode(&person)
  if err != nil {
   http.Error(w, err.Error(), http.StatusBadRequest)
   return
  }
  json.NewEncoder(w).Encode(person)
 } else {
  http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
 }
}