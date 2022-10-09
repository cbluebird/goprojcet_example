package model

import (
	"log"
	"resume/db"
)

func InitModel() {
	ok := db.DB.AutoMigrate(&User{}, &Resume{})
	if ok != nil {
		log.Panicln("Database Error: ", ok)
	}
}
