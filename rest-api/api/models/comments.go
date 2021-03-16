package models

import "time"

//Comment is an...
type Comment struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Comment string `gorm:"size:255 ;not null" json:"comments"`
	UserID  uint32 `gorm:"not null" json:"user_id"`

	PostID uint32 `gorm:"not null" json:"post_id"`

	CreatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"updated_at"`
}

func NewComment(comment Comment) error {

	db := Connect()

	err := db.Create(&comment).Error

	return err
}
