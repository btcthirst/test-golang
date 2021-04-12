package models

import "time"

//Post is an...
type Post struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	ImageURL string `gorm:"type:varchar(255)" json:"image_url"`
	Body     string `gorm:"type:text" json:"body"`
	UserID   uint64 `gorm:"not null" json:"user_id"`

	CreatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"updated_at"`
}

func NewPost(post Post) error {

	db := Connect()

	err := db.Create(&post).Error

	return err
}

func CheckPost(post Post) bool {
	//point to update func
	db := Connect()
	var posts []Post
	db.Find(&posts)
	for _, p := range posts {
		if p.Title == post.Title || p.Body == post.Body {
			return false
		}

	}
	return true
}
