package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raygervais/dfw/pkg/db"
)

// Create a GIN server
func NewServer(port int) *gin.Engine {
	// Create a new GIN server
	r := gin.Default()

	// Return the server
	return r
}

// Create CRUD routes for DB
// func NewRoutes(r *gin.Engine, db *db.Database) {

// 	// Create a new route group
// 	v1 := r.Group("/v1")

// 	// Add a route to the server
// 	v1.POST("/rooms", func(c *gin.Context) {
// 		var room entities.Room
// 		if err := c.BindJSON(&room); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		if err := db.AddRoom(&room); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusCreated, room)
// 		return
// 	})

// 	// Add a route to the server
// 	v1.GET("/rooms/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		room := db.GetRoom(id)
// 		if room != nil {
// 			c.JSON(http.StatusOK, room)
// 			return
// 		} else {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"message": "Room not found",
// 			})
// 			return
// 		}
// 	})

// 	// Add a route to the server
// 	v1.DELETE("/rooms/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		db.RemoveRoom(id)
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Room deleted",
// 		})
// 	})

// 	v1.GET("/rooms", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, db.Rooms)
// 	})

// 	// Add a route to the server
// 	v1.POST("/rooms/:id/prompts", func(c *gin.Context) {
// 		id := c.Param("id")
// 		var prompt entities.Prompt
// 		if err := c.BindJSON(&prompt); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.AddPrompt(id, &prompt); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		c.JSON(http.StatusCreated, prompt)
// 	})

// 	v1.GET("/rooms/:id/prompts", func(c *gin.Context) {
// 		id := c.Param("id")
// 		prompts := db.GetPrompts(id)
// 		c.JSON(http.StatusOK, prompts)
// 	})

// 	v1.POST("/rooms/:id/comments", func(c *gin.Context) {
// 		id := c.Param("id")
// 		var comment entities.Comment
// 		if err := c.BindJSON(&comment); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.AddComment(id, &comment); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		c.JSON(http.StatusCreated, comment)
// 	})

// 	v1.GET("/rooms/:id/comments", func(c *gin.Context) {
// 		id := c.Param("id")
// 		comments := db.GetComments(id)
// 		c.JSON(http.StatusOK, comments)
// 	})

// }

// Create a new route group
func NewRoutes(r *gin.Engine, db *db.Database) {
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("static", "static")
	rootGroup := r.Group("/")

	// 	fs := http.FileServer(http.Dir("./static"))
	// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Send HTML
	rootGroup.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "RAY FIGURED IT OUT, which I BELIEVE IS AN AMAZING THING BEST",
		})
	})

	// rootGroup.GET("/rooms/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	room := db.GetRoom(id)

	// 	c.HTML(http.StatusOK, filepath.Join("pkg", "templates", "room.tmpl"), room)
	// 	return

	// 	// TODO: make 404 screen

	// })
}
