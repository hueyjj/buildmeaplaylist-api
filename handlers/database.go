package handlers

import (
	"database/sql"
	"log"
	"os"
)

//Postgres holds database driver
type Postgres struct {
	db *sql.DB
}

//NewConnection configures and connects to a postgresql database
func NewConnection(connStr string) *Postgres {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return &Postgres{
		db: db,
	}
}
