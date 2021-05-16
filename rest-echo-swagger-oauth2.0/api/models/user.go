package models

import (
	"gorm.io/gorm"
)

type UserIn struct {
	Name     string
	Email    string
	Password string
	ID       uint64
}

type UserOut struct {
	Name  string
	Email string
	ID    uint64
}

type User struct {
	Name     string
	Email    string
	Password string
	gorm.Model
	ID uint64
	//CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	//UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

func (u User) Create() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "error connect", err
	}
	err = db.Create(&u).Error
	if err != nil {
		return "error create", err
	}

	return "create success", nil

}

func (u User) Update() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "error connect", err
	}
	err = db.Model(&u).Updates(map[string]interface{}{
		"name":     u.Name,
		"email":    u.Email,
		"password": u.Password,
	}).Error
	if err != nil {
		return "error update", err
	}

	return "update success", nil

}

func (u User) Delete() (string, error) {
	db, err := ConnectDB()
	if err != nil {
		return "error connect", err
	}
	err = db.Delete(&u).Error
	if err != nil {
		return "error delete", err
	}

	return "delete success", nil

}
