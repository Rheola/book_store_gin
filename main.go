package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"text/template"
)

// This function's name is a must.
// App Engine uses it to drive the requests properly.
func main() {
	// Starts a new Gin instance
	r := gin.New()
	r.Static("/css", "./css")
	r.Static("/font-awesome", "./font-awesome")
	r.Static("/js", "./js")

	r.GET("/admin", adminIndex)

	// Handle all requests using net/http
	http.Handle("/", r)
	r.Run(":8080")
}

func adminIndex(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/index.html")
	if err != nil {
		c.String(500, "Внутренняя ошибка сервера")
	}

	obj := gin.H{"title": "Главная"}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Внутренняя ошибка сервера")
	}
}
