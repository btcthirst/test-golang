package models

import "time"

type Comment struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Body      string
	Email     string
	PostID    uint64
}

func NewComment(cm *Comment) {
	var c Comment

	db := GetDB()
	c = *cm
	db.Create(&c)
}
