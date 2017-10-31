package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"text/template"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var db = InitDB()

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/book_store?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("db.open() error: %v", err))
	}

	return db
}

// This function's name is a must.
// App Engine uses it to drive the requests properly.
func main() {

	// Starts a new Gin instance
	r := gin.New()
	r.Static("/css", "./css")
	r.Static("/font-awesome", "./font-awesome")
	r.Static("/js", "./js")

	r.GET("/admin", adminIndex)

	//авторы
	r.GET("/admin/author", adminAuthorIndex)
	r.GET("/admin/author/view/:id", authorView)
	r.GET("/admin/author/create", authorForm)
	r.POST("/admin/author/save", authorSave)

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
