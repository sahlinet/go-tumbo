package project

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Project struct {
	gorm.Model
	Name string
}

var db gorm.DB

func NewProject() (*Project, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return &Project{}, fmt.Errorf("error in gorm.Open: %s", err.Error())
	}
	p := Project{}
	db.Create(&p)
	return &p, nil
}

