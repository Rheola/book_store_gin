package models

import (
	"strings"
	"unicode/utf8"
	"fmt"
)

type Category struct {
	Id   int64  `sql:"primary_key:yes;"`
	Name string `sql:"type:text;" json:"content" form:"name"`
}

func (Category) TableName() string {
	return "category"
}

func (c Category) Validate() error {
	c.Name = strings.TrimSpace(c.Name)

	if utf8.RuneCountInString(c.Name) > 50 {
		return fmt.Errorf("Невалидное поле Name: %s (длина должна быть до 50 символов)", c.Name)
	}

	if c.Name == "" {
		return fmt.Errorf("Пустое поле Name: %s", c.Name)
	}

	return nil
}
