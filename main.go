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
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/kill", killHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	log.Println("Starting server on :8080")

	// Listen for interrupt signals for graceful shutdown
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Graceful shutdown logic
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Println("Server Exited Properly")
}
