package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	if data.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", data.Name)
}
