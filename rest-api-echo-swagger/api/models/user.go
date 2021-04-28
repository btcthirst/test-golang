package models

import (
	"time"
)

type User struct {
	Name      string
	Email     string
	Password  string
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(u *User) {
	var us User
	db := GetDB()

	us = *u

	db.Create(&us)
}
