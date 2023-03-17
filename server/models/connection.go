package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	dsn := "host=localhost user=postgres password=99831745gusta dbname=ownboss port=5432"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Company{}, &User{}, &Employee{}, &Service{}, &Event{})
	if err != nil {
		log.Print(err)
	}
}
