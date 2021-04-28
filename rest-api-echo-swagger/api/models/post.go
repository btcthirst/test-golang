package models

import "time"

type Post struct {
	ID         uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Title      string
	Body       string
	ImageURL   string
	UserID     uint64
	CommentsID uint64
}

func NewPost(pp *Post) {
	var p Post
	db := GetDB()
	p = *pp

	db.Create(&p)
}
