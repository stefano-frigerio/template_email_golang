package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type Template struct {
	Name       string
	Language   string
	Parameters map[string]interface{}
	EmailTxt   string
}

func generateTemplate() {
	param := map[string]interface{}{
		"name":  "Stefano",
		"promo": "freecoupon",
	}

	db.Create(&Template{Name: "Promo", Language: "Italian", EmailTxt: "Ciao", Parameters: param})
	db.Create(&Template{Name: "Promo", Language: "English", EmailTxt: "Hello", Parameters: param})
}

func retrieveTemplate(Name string, Language string) *Template {

	t := Template{Name: Name, Language: Language}

	db.AutoMigrate(&Template{})

	var template Template
	db.Where(&Template{Name: Name, Language: Language}).Find(&template)
	if template.Language == "" {
		db.First(&template, "Name = ?", Name)
	}
	fmt.Println(template)
	return &t
}

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("template.db"), &gorm.Config{})
	if err != nil {
		panic("Connection failed")
	}
	//generateTemplate()
	retrieveTemplate("Promo", "English")
}
