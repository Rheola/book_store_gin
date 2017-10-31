package models

type Author struct {
	Id   int64  `sql:"primary_key:yes;"`
	Name string `sql:"type:text;" json:"content" form:"name"`
}

func (Author) TableName() string {
	return "author"
}
