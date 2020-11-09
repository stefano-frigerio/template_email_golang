package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Template struct {
	Name     string
	Language string
}

func retrieveTemplate(Name string) *Template {
	t := Template{Name: Name}

	db.AutoMigrate(&Template{})

	var template Template
	db.First(&template, "Name = ?", Name)
	fmt.Println(template)
	return &t
}

func main() {
	db, err := gorm.Open(sqlite.Open("template.db"), &gorm.Config{})
	if err != nil {
		panic("Connection failed")
	}
	db.Create(&Template{Name: "IT", Language: "Italian"})
	retrieveTemplate("IT")
}
