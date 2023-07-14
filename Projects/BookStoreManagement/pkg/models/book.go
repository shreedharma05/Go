package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shreedharma05/Go/BookStoreManagement/pkg/config"
)

var db *gorm.DB

type Book struct {
	Name        string
	Author      string
	Publication string
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
