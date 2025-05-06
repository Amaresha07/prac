package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "John Doe"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)
}
