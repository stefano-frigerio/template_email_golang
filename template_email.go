package main

import (
	"fmt"

	"github.com/alexkappa/mustache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var param map[string]interface{}

type Template struct {
	Name     string
	Subject  string
	Body     string
	Language string
}

func generateTemplate() {
	db.Create(&Template{Name: "Promo", Subject: "Oggetto", Body: "<div>Ciao {{first_name}} {{last_name}} </div>", Language: "Italian"})
	db.Create(&Template{Name: "Promo", Subject: "Subject", Body: "<div>Hello {{first_name}} {{last_name}} </div>", Language: "English"})
}

func retrieveTemplate(Name string, Language string, param map[string]interface{}) *Template {

	db.AutoMigrate(&Template{})

	var template Template
	db.Where(&Template{Name: Name, Language: Language}).Find(&template)
	if template.Language == "" {
		db.First(&template, "Name = ?", Name)
	}

	t := mustache.New()
	err := t.ParseString(template.Body)
	if err != nil {
		// handle error
	}
	tbody, _ := t.RenderString(param)
	if err != nil {
		// handle error
	}
	template.Body = tbody

	return &template
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("template.db"), &gorm.Config{})
	if err != nil {
		panic("Connection failed")
	}
	//generateTemplate()

	param = make(map[string]interface{})
	param["first_name"] = "Stefano"
	param["last_name"] = "Frigerio"

	tempprint := retrieveTemplate("Promo", "English", param)

	fmt.Println(tempprint)
}
