package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title      string
	Body       string
	ImageURL   string
	UserID     uint64
	CommentsID uint64 //??нужны ли двойные ссылки коменты к постам и посты к коментам
}

func NewPost(p Post) {
	db := GetDB()
	db.Create(&p)
}
