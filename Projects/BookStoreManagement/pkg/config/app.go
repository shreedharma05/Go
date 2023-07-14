package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var ConnectDB = func() {
	d, err := gorm.Open("mysql", "shreedharma:Shree@05@/BookStore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

var GetDB = func() *gorm.DB {
	return db
}
