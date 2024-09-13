package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler serves the "/" route with a simple hello message.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, welcome to my API! \nCOMP4016 - A01325136 - Andrew")
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "bar")
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data struct {
			Name string `json:"name"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		response := fmt.Sprintf("Hello %s!", data.Name)
		w.Write([]byte(response))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
