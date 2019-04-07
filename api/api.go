package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hueyjj/buildmeaplaylist-api/handlers"
)

// Serve creates a server handling routes to create, update, and
// remove database records
func Serve(ip, port string) {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")

	ipport := fmt.Sprintf("%s:%s", ip, port)
	server := &http.Server{
		Handler:      router,
		Addr:         ipport,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Unable to start server: %v\n", err)
		}
	}()

	log.Printf("Server running at %s\n", ipport)

	// Accept Ctrl+C as signal to quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Wait a deadline to end
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("Server gracefully shutting down")
	os.Exit(0)
}
