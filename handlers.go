package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"golang.org/x/exp/rand"
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

// Assign03 Below:
func saveStringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var payload struct {
			Data string `json:"data"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.Data == "" {
			http.Error(w, "Invalid or missing 'data' field", http.StatusBadRequest)
			return
		}
		// Write the data to a file in the mounted volume
		err = os.WriteFile("/data/saved_string.txt", []byte(payload.Data), 0644)
		if err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getStringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := os.ReadFile("/data/saved_string.txt")
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "Data not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to read data", http.StatusInternalServerError)
			}
			return
		}
		w.Write(data)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func busyWaitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Starting CPU-intensive task..."))
		go func() {
			end := time.Now().Add(3 * time.Minute)
			for time.Now().Before(end) {
				// Perform some CPU-intensive calculations
				_ = math.Sqrt(float64(rand.Intn(1000000)))
			}
		}()
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func isAliveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("true"))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
