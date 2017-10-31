package models

import (
	"testing"
	"strings"
)

func getValidCategory() Category {
	a := Category{
		Name: "Классическая литература",
	}
	return a
}

func TestCategoryValidFail(t *testing.T) {
	c := getValidCategory()
	c.Name = ""

	if c.Validate() == nil {
		t.Error("Expected Empty Error, got nil ")
	}

	c.Name = strings.Repeat("z", 51)

	if c.Validate() == nil {
		t.Error("Expected length Error, got nil ")
	}
}

func TestValidCategoryTableName(t *testing.T) {

	c := getValidCategory()
	if c.TableName() != "category" {
		t.Error("Expected category, got  ", c.TableName())
	}
}

func TestValidCategoryFailOk(t *testing.T) {
	c := getValidCategory()
	err := c.Validate()
	if err != nil {
		t.Error("Fail Validate. Error  ", err)
	}
}
