package controllers

import (
	"fmt"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
)

func getPostsAll() []models.Post {
	var posts []models.Post
	db := models.GetDB()
	db.Find(&posts)

	return posts
}

func getPostByID(id uint64) models.Post {
	var post models.Post
	db := models.GetDB()
	db.Where("id=?", id).Find(&post)
	//fmt.Println(id, user)
	return post
}

func updatePost(id uint64, pt *models.Post) (string, error) {
	db := models.GetDB()
	var post, p models.Post
	p = *pt
	db.Where("id=?", id).Find(&post)
	if post.Title != p.Title {
		return "title", db.Model(&models.Post{}).Where("id=?", id).Update("title", p.Title).Error
	}
	if post.Body != p.Body {
		return "body", db.Model(&models.Post{}).Where("id=?", id).Update("body", p.Body).Error
	}
	if post.ImageURL != p.ImageURL {
		return "image", db.Model(&models.Post{}).Where("id=?", id).Update("image_url", p.ImageURL).Error
	}
	return "nothing", nil
}

func deletePost(id uint64) string {
	db := models.GetDB()
	var p models.Post
	db.Where("id=?", id).Find(&p)
	if p.ID == 0 {
		text := fmt.Sprintf("post with id : %v does not exist", id)
		return text
	}

	c := db.Where("id=?", id).Delete(&models.Post{}).Error

	if c == nil {
		text := fmt.Sprintf("post with id : %v deleted", id)
		return text
	}

	return "something wrong"

}
