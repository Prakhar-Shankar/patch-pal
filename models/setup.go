package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlite.Open("patchpal.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect with Database")
	}

	err = database.AutoMigrate(&Bug{})
	if err != nil{
		panic("Failed to migrate bug model: "+ err.Error())
	}
	DB = database
}