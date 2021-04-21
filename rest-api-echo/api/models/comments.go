package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body   string
	Email  string
	PostID uint64
}

func NewComment(cm *Comment) {
	var c Comment

	db := Connect()
	c = *cm
	db.Create(&c)
}
