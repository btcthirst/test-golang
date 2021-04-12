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

func CommentPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetComment(w, r)
	case http.MethodPost:
		PostComment(w, r)
	case http.MethodPut:
		PutComment(w, r)
	case http.MethodDelete:
		DeleteComment(w, r)
	default:
		fmt.Fprintf(w, "%v -there is no such method on this page", r.Method)
	}
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] != "" {
		id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		comments := models.GetByID(models.Comments, id)
		utilites.ToJSON(w, comments, http.StatusOK)
	} else {
		comments := models.GetAll(models.Comments)
		utilites.ToJSON(w, comments, http.StatusOK)
	}

}

func PostComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	defer r.Body.Close()

	//try to catch empty post request
	bug := json.NewDecoder(r.Body).Decode(&comment)
	if bug != nil {
		utilites.ToJSON(w, "can`t use empty post method", http.StatusOK)
	} else {
		if models.CheckComment(comment) {
			models.NewComment(comment)
			utilites.ToJSON(w, "comment created", http.StatusOK)
		} else {

			utilites.ToJSON(w, "not uniq comment", http.StatusOK)
		}

	}

}

func PutComment(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] == "" {
		utilites.ToJSON(w, "no comment to update", http.StatusOK)
	}
	id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := models.Connect()

	var comment models.Comment
	defer r.Body.Close()

	//try to catch empty post request
	bug := json.NewDecoder(r.Body).Decode(&comment)
	if bug != nil {
		utilites.ToJSON(w, "can`t use empty put method", http.StatusOK)
	} else {
		comment.ID = id
		db.Model(&comment).Update("comment", comment.Comment)
		utilites.ToJSON(w, "Comment created", http.StatusOK)

	}

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] == "" {
		utilites.ToJSON(w, "no comment to delete", http.StatusOK)
	}
	id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	db := models.Connect()
	db.Where("id=?", id).Delete(&models.Comment{})
	utilites.ToJSON(w, "comment deleted", http.StatusOK)
}
