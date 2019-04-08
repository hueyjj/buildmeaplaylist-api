package handlers

import (
	"database/sql"
	"log"
	"os"
)

//Controller holds database driver
type Controller struct {
	db *sql.DB
}

//Member represent in the members table
type Member struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

//NewConnection configures and connects to a postgresql database
func NewConnection(connStr string) *Controller {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return &Controller{
		db: db,
	}
}
