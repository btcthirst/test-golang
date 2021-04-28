package controllers

import (
	"fmt"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
)

func getUsersAll() []models.User {
	var users []models.User
	db := models.GetDB()
	db.Find(&users)

	return users
}

func getUsersByID(id uint64) models.User {
	var user models.User
	db := models.GetDB()
	db.Where("id=?", id).Find(&user)
	//fmt.Println(id, user)
	return user
}

func updateUser(id uint64, us *models.User) (string, error) {
	db := models.GetDB()
	var user, u models.User
	u = *us
	db.Where("id=?", id).Find(&user)
	if user.Name != u.Name {
		return "name", db.Model(&models.User{}).Where("id=?", id).Update("name", u.Name).Error
	}
	if user.Password != u.Password {
		return "password", db.Model(&models.User{}).Where("id=?", id).Update("password", u.Password).Error
	}
	if user.Email != u.Email {
		return "email", db.Model(&models.User{}).Where("id=?", id).Update("email", u.Email).Error
	}
	return "nothing", nil
}

func deleteUser(id uint64) string {
	db := models.GetDB()
	var u models.User
	db.Where("id=?", id).Find(&u)
	if u.ID == 0 {
		text := fmt.Sprintf("user with id : %v does not exist", id)
		return text
	}

	c := db.Where("id=?", id).Delete(&models.User{}).Error

	if c == nil {
		text := fmt.Sprintf("user with id : %v deleted", id)
		return text
	}

	return "something wrong"

}
