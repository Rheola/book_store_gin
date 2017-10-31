package main

import (
	"github.com/gin-gonic/gin"
	"text/template"
	"github.com/rheola/book_store_gin/models"
	"strconv"
)

func adminCategoryIndex(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/category/index.html")
	if err != nil {
		c.String(500, "Внутренняя ошибка сервера (парсинг шаблона)")
	}

	var categories []models.Category

	db.Find(&categories)

	obj := gin.H{"title": "Категории", "categories": categories}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Внутренняя ошибка сервера (выполнение шаблона)")
	}
}

func categoryForm(c *gin.Context) {
	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/category/create.html")
	if err != nil {
		c.String(500, "Внутренняя ошибка сервера (парсинг шаблона)")
		return
	}

	obj := gin.H{"title": "Добавить категорию"}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Внутренняя ошибка сервера (выполнение шаблона)")
		return
	}
}

func categoryView(c *gin.Context) {
	id := c.Params.ByName("id")

	tmpl, err := template.ParseFiles("templates/layouts/admin.html", "templates/admin/category/view.html")
	if err != nil {
		c.String(500, "Internal Server Error")
	}

	var category models.Category

	db.First(&category, id)

	obj := gin.H{"title": "Просмотр категории", "category": category}

	if err := tmpl.Execute(c.Writer, obj); err != nil {
		c.String(500, "Internal Server Error")
	}
}

func categorySave(c *gin.Context) {
	var form models.Category

	c.Bind(&form)

	var category models.Category

	category.Name = form.Name

	db.Save(&category)

	c.Redirect(301, "/admin/category/view/"+strconv.FormatInt(category.Id, 10))
}
func toDoCategory(c *gin.Context) {
	c.String(500, "Не реализовано")
}


