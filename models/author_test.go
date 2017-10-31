package models

import (
	"testing"
	"strings"
)

func getValidAuthor() Author {
	a := Author{
		Name: "Александр Сергеевич Пушкин",
	}
	return a
}

func TestValidFail(t *testing.T) {
	a := getValidAuthor()
	a.Name = ""

	if a.Validate() == nil {
		t.Error("Expected Empty  Error, got nil ")
	}

	a.Name = strings.Repeat("z", 256)
	if a.Validate() == nil {
		t.Error("Expected Lenght  Error, got nil ")
	}
}

func TestValidTableName(t *testing.T) {

	a := getValidAuthor()
	if a.TableName() != "author" {
		t.Error("Expected author, got  ", a.TableName())
	}
}

func TestValidFailOk(t *testing.T) {
	a := getValidAuthor()
	err := a.Validate()
	if err != nil {
		t.Error("Fail Validate. Error  ", err)
	}
}
