package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var stop = make(chan os.Signal, 1)

func main() {
	// Main application mux and server on port 8080
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/kill", killHandler)
	mux.HandleFunc("/configValue", configValueHandler)
	mux.HandleFunc("/secretValue", secretValueHandler)
	mux.HandleFunc("/envValue", envValueHandler)
	mux.HandleFunc("/saveString", saveStringHandler)
	mux.HandleFunc("/getString", getStringHandler)
	mux.HandleFunc("/busywait", busyWaitHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the main server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	log.Println("Starting server on :8080")

	// Start a separate server for /isAlive on port 30010
	isAliveMux := http.NewServeMux()
	isAliveMux.HandleFunc("/isAlive", isAliveHandler)

	isAliveSrv := &http.Server{
		Addr:    ":30010",
		Handler: isAliveMux,
	}

	// Start the isAlive server in a goroutine
	go func() {
		if err := isAliveSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("isAlive ListenAndServe(): %v", err)
		}
	}()

	log.Println("Starting isAlive server on :30010")

	// Listen for interrupt signals for graceful shutdown
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Graceful shutdown logic for both servers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	if err := isAliveSrv.Shutdown(ctx); err != nil {
		log.Fatalf("isAlive Server Shutdown Failed: %+v", err)
	}
	log.Println("Servers Exited Properly")
}
