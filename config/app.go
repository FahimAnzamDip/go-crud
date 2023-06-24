package config

import (
	"github.com/fahimanzamdip/go-crud/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_crud?charset=utf8&parseTime=True&loc=Local")
	utils.CheckNilError(err)
	db = d
}

func GetDB() *gorm.DB {
	return db
}
