package models

import "time"

//User is an ...
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"type:varchar(150);not null;unique_index" json:"user_name"`
	Email     string    `gorm:"type:varchar(150);not null;unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(150);not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);default:current_timestamp();not null" json:"updated_at"`
}

func NewUser(user User) error {

	db := Connect()
	err := db.Create(&user).Error

	return err
}

func CheckUser(user User) bool {
	//point to update func
	db := Connect()
	var users []User
	db.Find(&users)
	for _, u := range users {
		if u.UserName == user.UserName || u.Email == user.Email {
			return false
		}

	}
	return true
}
