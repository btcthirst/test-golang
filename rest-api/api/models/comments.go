package models

import "time"

//Comment is an...
type Comment struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Comment string `gorm:"size:255 ;not null" json:"comments"`
	UserID  uint64 `gorm:"not null" json:"user_id"`

	PostID uint64 `gorm:"not null" json:"post_id"`

	CreatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"updated_at"`
}

func NewComment(comment Comment) error {

	db := Connect()

	err := db.Create(&comment).Error

	return err
}

func CheckComment(comment Comment) bool {
	//point to update func
	db := Connect()
	var comments []Comment
	db.Find(&comments)
	for _, c := range comments {
		if c.Comment == comment.Comment {
			return false
		}

	}
	return true
}
