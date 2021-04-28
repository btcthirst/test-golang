package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "max_dev"
	password = "qwe123RTY123#p"
	host     = "127.0.0.1"
	port     = "3306"
	database = "rest_api_echo_swagger" /*db must be created before calling*/
)

var DB *gorm.DB

//NewDB() *gorm.DB
func NewDB() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
