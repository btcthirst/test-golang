package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
)

func CommentPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetComments(w, r)
	case http.MethodPost:
		PostComment(w, r)
	case http.MethodDelete:
		DeleteComment(w, r)
	case http.MethodPut:
		PutComment(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsuported method %v, to %v\n", r.Method, r.URL)
		log.Printf("Unsuported method %v, to %v\n", r.Method, r.URL)

	}
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments := models.GetAll(models.Comments)
	je := json.NewEncoder(w)
	je.Encode(comments)
}

func PostComment(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var comment models.Comment
	err := json.Unmarshal(body, &comment)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewComment(comment)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "Comment created success!", http.StatusCreated)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to delete comment %v", id)

	db := models.Connect()
	var comments []models.Comment

	db.Find(&comments)
	for _, comment := range comments {
		if comment.ID == id {
			db.Delete(&comment)
			log.Print("Success!!!")
			fmt.Fprintf(w, "comment with id : %v deleted", id)
		}
	}
}

func PutComment(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to put comment %v", id)

	db := models.Connect()
	var comments []models.Comment

	db.Find(&comments)

	body := utils.BodyParser(r)
	var comment1 models.Comment
	err = json.Unmarshal(body, &comment1)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	for _, comment := range comments {
		if comment.ID == id {
			if comment.Comment != comment1.Comment {
				db.Model(&comment).Update("comment", comment1.Comment)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")

			} else {
				fmt.Fprint(w, "there are nothing to change")
			}

		}
	}

}
