package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzu/gorm"
	_ "github.com/jinzu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// gorm.open("{name of database", "nameOfUser:password/{nameoftable}?")
	d, err := gorm.Open("bookstoredb", "bookstore:Hh44974497/book?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
