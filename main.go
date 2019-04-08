package main

import (
	"log"
	"os"

	"github.com/hueyjj/buildmeaplaylist-api/api"
	"github.com/hueyjj/buildmeaplaylist-api/handlers"
	"github.com/hueyjj/buildmeaplaylist-api/session"
	_ "github.com/lib/pq"
)

func main() {
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	connStr := os.Getenv("DB")
	sessionKey := os.Getenv("SESSION_KEY")

	if ip != "" && port != "" && connStr != "" && sessionKey != "" {
		log.Printf("Environment variable set: IP=%s PORT=%s DB=%s SESSION_KEY=%s\n", ip, port, connStr, sessionKey)

	} else {
		log.Printf("Some environmental variable(s) unset IP=%s PORT=%s DB=%s SESSION_KEY=%s\n", ip, port, connStr, sessionKey)
		os.Exit(1)
	}

	db := handlers.NewConnection(connStr)
	log.Println("Established connection to postgresql database.")

	store := session.NewCookieStore([]byte(sessionKey))
	log.Println("Session store created")

	api.Serve(db, store, ip, port)
}
