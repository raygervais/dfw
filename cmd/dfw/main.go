package main

import (
	"github.com/raygervais/dfw/pkg/db"
	"github.com/raygervais/dfw/pkg/http"
)

func main() {
	// Create database
	db := db.Database{}

	// Create server
	r := http.NewServer(8080)

	// Add Routes
	http.NewRoutes(r, &db)

	// Start server
	r.Run()
}
