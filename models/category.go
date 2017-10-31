package models

type Category struct {
	Id   int64  `sql:"primary_key:yes;"`
	Name string `sql:"type:text;" json:"content" form:"name"`
}

func (Category) TableName() string {
	return "category"
}
