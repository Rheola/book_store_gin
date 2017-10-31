package models

import (
	"strings"
	"unicode/utf8"
	"fmt"
)

type Author struct {
	Id   int64  `sql:"primary_key:yes;"`
	Name string `sql:"type:text;" json:"content" form:"name"`
}

func (Author) TableName() string {
	return "author"
}

func (a Author) Validate() error {
	a.Name = strings.TrimSpace(a.Name)

	if utf8.RuneCountInString(a.Name) > 255 {
		return fmt.Errorf("Невалидное поле Name: %s (длиннее 255)", a.Name)
	}

	if a.Name == "" {
		return fmt.Errorf("Пустое поле Name: %s", a.Name)
	}

	return nil
}
