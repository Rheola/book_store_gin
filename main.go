package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"text/template"
	"github.com/rheola/book_store_gin/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
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
	r.GET("/admin/author", adminAuthorIndex)

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

func adminAuthorIndex(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:root@/book_store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("db.open() error: %v", err))
	}

	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/author/index.html")
	if err != nil {
		c.String(500, "Внутренняя ошибка сервера (парсинг шаблона)")
	}

	var authors []models.Author

	db.Find(&authors)

	obj := gin.H{"title": "Авторы", "authors": authors}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Внутренняя ошибка сервера (выполнение шаблона)")
	}
}
