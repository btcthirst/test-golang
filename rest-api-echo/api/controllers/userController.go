package controllers

import (
	"fmt"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/models"
)

func getUsersAll() []models.UserDB {
	var users []models.UserDB
	db := models.Connect()
	db.Find(&users)

	return users
}

func getUsersByID(id uint64) models.UserDB {
	var user models.UserDB
	db := models.Connect()
	db.Where("id=?", id).Find(&user)
	//fmt.Println(id, user)
	return user
}

func updateUser(id uint64, us *models.User) (string, error) {
	db := models.Connect()
	var user, u models.UserDB
	u.User = *us
	db.Where("id=?", id).Find(&user)
	if user.Name != u.Name {
		return "name", db.Model(&models.UserDB{}).Where("id=?", id).Update("name", u.Name).Error
	}
	if user.Password != u.Password {
		return "password", db.Model(&models.UserDB{}).Where("id=?", id).Update("password", u.Password).Error
	}
	if user.Email != u.Email {
		return "email", db.Model(&models.UserDB{}).Where("id=?", id).Update("email", u.Email).Error
	}
	return "nothing", nil
}

func deleteUser(id uint64) string {
	db := models.Connect()
	var u models.UserDB
	db.Where("id=?", id).Find(&u)
	if u.ID == 0 {
		text := fmt.Sprintf("user with id : %v does not exist", id)
		return text
	}

	c := db.Where("id=?", id).Delete(&models.UserDB{}).Error

	if c == nil {
		text := fmt.Sprintf("user with id : %v deleted", id)
		return text
	}

	return "something wrong"

}
