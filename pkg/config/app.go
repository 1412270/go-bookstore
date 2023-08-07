package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost user=postgres password=Abc12345 dbname=bookstore sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
