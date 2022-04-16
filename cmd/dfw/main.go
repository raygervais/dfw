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

// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"path/filepath"
// 	"strings"
// )

// func main() {
// 	fs := http.FileServer(http.Dir("./static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	http.HandleFunc("/", serveTemplate)

// 	log.Println("Listening on :3000...")
// 	err := http.ListenAndServe(":3000", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func serveTemplate(w http.ResponseWriter, r *http.Request) {
// 	lp := filepath.Join("pkg", "templates", "base.html")
// 	fp := filepath.Join("pkg", "templates", filepath.Clean(strings.Split(r.URL.Path, ".")[0]+".html"))

// 	fmt.Println(fp)

// 	tmpl, _ := template.ParseFiles(lp, fp)
// 	tmpl.ExecuteTemplate(w, "layout", nil)
// }
