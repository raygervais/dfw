package main

import (
	"github.com/raygervais/dfw/pkg/db"
	"github.com/raygervais/dfw/pkg/http"
	email "github.com/raygervais/dfw/pkg/smtp"
)

func main() {
	// Create database
	db := db.Database{}

	// Create server
	r := http.NewServer(8080)

	// Add Routes
	http.NewRoutes(r, &db)

	// Create a go thread to listen for requests
	go r.Run()

	// Create a go thread to send emails
	email.Thread(&db)

}
