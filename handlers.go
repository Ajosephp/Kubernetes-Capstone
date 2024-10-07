package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// homeHandler serves the "/" route with a simple hello message.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, welcome to my API! \nCOMP4016 - Andrew")
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
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
		// Decode the JSON body into the `data` struct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		// Ensure the name is present in the JSON body
		if data.Name == "" {
			http.Error(w, "Name field is required", http.StatusBadRequest)
			return
		}
		response := fmt.Sprintf("Hello %s!", data.Name)
		w.Write([]byte(response))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func killHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutting down..."))
	go func() {
		time.Sleep(1 * time.Second)
		stop <- os.Interrupt
	}()
}

// Assign02 Below:
// configValueHandler retrieves the value from a ConfigMap
func configValueHandler(w http.ResponseWriter, r *http.Request) {
	value := os.Getenv("CONFIG_VALUE")
	if value == "" {
		http.Error(w, "configValue not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, value)
}

func secretValueHandler(w http.ResponseWriter, r *http.Request) {
	value := os.Getenv("SECRET_VALUE")
	if value == "" {
		http.Error(w, "secretValue not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, value)
}

func envValueHandler(w http.ResponseWriter, r *http.Request) {
	value := os.Getenv("ENV_VALUE")
	if value == "" {
		http.Error(w, "envValue not found", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, value)
}
