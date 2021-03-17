package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
	"github.com/gorilla/mux"
)

func CommentGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	comment := models.GetByID(models.Comments, id)
	utils.ToJSON(w, comment, http.StatusOK)

	//end xml
	//utils.ToXML(w, comment, http.StatusOK)
}

func CommentsGet(w http.ResponseWriter, r *http.Request) {
	comments := models.GetAll(models.Comments)
	utils.ToJSON(w, comments, http.StatusOK)
}

func CommentPost(w http.ResponseWriter, r *http.Request) {
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

func CommentDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := models.Delete(models.Comments, id)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusOK)
}
