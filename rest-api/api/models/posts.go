package models

import "time"

//Post is an...
type Post struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	ImageURL string `gorm:"type:varchar(255)" json:"image_url"`
	Body     string `gorm:"type:text" json:"body"`
	UserID   uint32 `gorm:"not null" json:"user_id"`

	CreatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"updated_at"`
}

func NewPost(post Post) error {

	db := Connect()

	err := db.Create(&post).Error

	return err
}
