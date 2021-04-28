package controllers

import (
	"fmt"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
)

func getCommentsAll() []models.Comment {
	var comments []models.Comment
	db := models.GetDB()
	db.Find(&comments)

	return comments
}

func getCommentByID(id uint64) models.Comment {
	var comment models.Comment
	db := models.GetDB()
	db.Where("id=?", id).Find(&comment)
	//fmt.Println(id, user)
	return comment
}

func updateComment(id uint64, cm *models.Comment) (string, error) {
	db := models.GetDB()
	var comment, cnt models.Comment
	cnt = *cm
	db.Where("id=?", id).Find(&comment)
	if comment.Body != cnt.Body {
		return "text", db.Model(&models.Comment{}).Where("id=?", id).Update("body", cnt.Body).Error
	}

	return "nothing", nil
}

func deleteComment(id uint64) string {
	db := models.GetDB()
	var cm models.Comment
	db.Where("id=?", id).Find(&cm)
	if cm.ID == 0 {
		text := fmt.Sprintf("comment with id : %v does not exist", id)
		return text
	}

	err := db.Where("id=?", id).Delete(&models.User{}).Error

	if err == nil {
		text := fmt.Sprintf("comment with id : %v deleted", id)
		return text
	}

	return "something wrong"

}
