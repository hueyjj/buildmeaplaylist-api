package main

import (
	"log"
	"os"

	"github.com/hueyjj/buildmeaplaylist-api/api"
)

func main() {
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")

	if ip != "" && port != "" {
		log.Printf("Environment variable set: IP=%s PORT=%s\n", ip, port)

	} else {
		log.Printf("IP=%s and/or PORT=%s unset\n", ip, port)
		os.Exit(1)
	}

	api.Serve(ip, port)
}
