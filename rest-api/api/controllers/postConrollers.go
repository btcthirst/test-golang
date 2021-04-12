package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utilites"
)

func PostPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetPost(w, r)
	case http.MethodPost:
		PostPost(w, r)
	case http.MethodPut:
		PutPost(w, r)
	case http.MethodDelete:
		DeletePost(w, r)
	default:
		fmt.Fprintf(w, "%v -there is no such method on this page", r.Method)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] != "" {
		id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		posts := models.GetByID(models.Posts, id)
		utilites.ToJSON(w, posts, http.StatusOK)
	} else {
		posts := models.GetAll(models.Posts)
		utilites.ToJSON(w, posts, http.StatusOK)
	}

}

func PostPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	defer r.Body.Close()

	//try to catch empty post request
	bug := json.NewDecoder(r.Body).Decode(&post)
	if bug != nil {
		utilites.ToJSON(w, "can`t use empty post method", http.StatusOK)
	} else {
		if models.CheckPost(post) {
			models.NewPost(post)
			utilites.ToJSON(w, "post created", http.StatusOK)
		} else {

			utilites.ToJSON(w, "not uniq post", http.StatusOK)
		}

	}

}

func PutPost(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] == "" {
		utilites.ToJSON(w, "no posts to update", http.StatusOK)
	}
	id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := models.Connect()

	var post models.Post
	defer r.Body.Close()

	//try to catch empty post request
	bug := json.NewDecoder(r.Body).Decode(&post)
	if bug != nil {
		utilites.ToJSON(w, "can`t use empty put method", http.StatusOK)
	} else {
		post.ID = id
		db.Model(&post).Updates(models.Post{Title: post.Title, ImageURL: post.ImageURL, Body: post.Body})
		utilites.ToJSON(w, "post updated", http.StatusOK)

	}

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] == "" {
		utilites.ToJSON(w, "no posts to delete", http.StatusOK)
	}
	id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := models.Connect()
	db.Where("id=?", id).Delete(&models.Post{})
	utilites.ToJSON(w, "post deleted", http.StatusOK)
}
