package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func NewUser(u User) {
	db := GetDB()
	db.Create(&u)
}
