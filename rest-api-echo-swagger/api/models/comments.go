package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string
	Email  string
	PostID uint64
}

func NewComment(c Comment) {
	db := GetDB()
	db.Create(&c)
}
