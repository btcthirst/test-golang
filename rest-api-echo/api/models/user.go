package models

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type UserDB struct {
	gorm.Model
	User
}

func NewUser(u *User) {
	db := Connect()
	var us UserDB
	us.User = *u

	db.Create(&us)
}
