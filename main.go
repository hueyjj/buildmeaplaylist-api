package main

import (
	"log"
	"os"

	"github.com/hueyjj/buildmeaplaylist-api/api"
	"github.com/hueyjj/buildmeaplaylist-api/handlers"
	_ "github.com/lib/pq"
)

func main() {
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	connStr := os.Getenv("DB")

	if ip != "" && port != "" && connStr != "" {
		log.Printf("Environment variable set: IP=%s PORT=%s DB=%s\n", ip, port, connStr)

	} else {
		log.Printf("Some environmental variable(s) unset IP=%s PORT=%s DB=%s\n", ip, port, connStr)
		os.Exit(1)
	}

	db := handlers.NewConnection(connStr)
	log.Println("Established connection to postgresql database.")

	api.Serve(db, ip, port)
}
