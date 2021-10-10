package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Connect() {
	// Please define your user name and password for my sql.
	d, err := gorm.Open("postgres", "postgres://postgres:12345@localhost/godemo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
