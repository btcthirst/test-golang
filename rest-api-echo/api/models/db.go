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
	database = "rest_api_echo" /*db must be created before calling*/
)

//Connect () *gorm.DB
func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
