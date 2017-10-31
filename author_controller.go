package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
	"github.com/rheola/book_store_gin/models"
)

func authorForm(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/author/create.html")
	if err != nil {
		c.String(500, "Внутренняя ошибка сервера (парсинг шаблона)")
		return
	}

	obj := gin.H{"title": "Добавить автора"}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Внутренняя ошибка сервера (выполнение шаблона)")
		return
	}
}

func authorSave(c *gin.Context) {
	var form models.Author

	c.Bind(&form)

	var author models.Author

	author.Name = form.Name

	db.Save(&author)

	c.Redirect(301, "/admin/author/view/"+strconv.FormatInt(author.Id, 10))
}

func adminAuthorIndex(c *gin.Context) {
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

func authorView(c *gin.Context) {
	id := c.Params.ByName("id")

	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/author/view.html")
	if err != nil {
		c.String(500, "Internal Server Error")
	}

	var author models.Author

	db.First(&author, id)

	obj := gin.H{"title": "Просмотр автора", "author": author}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Internal Server Error")
	}
}

func toDoAuthor(c *gin.Context) {
	c.String(500, "Не реализовано")
}