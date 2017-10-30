package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Handle all requests using net/http
	http.Handle("/", r)

	r.Run(":8080")
}