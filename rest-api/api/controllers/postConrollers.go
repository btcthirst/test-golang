package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
	"github.com/gorilla/mux"
)

func PostGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	post := models.GetByID(models.Posts, id)
	utils.ToJSON(w, post, http.StatusOK)

	//end xml
	//utils.ToXML(w, post, http.StatusOK)
}

func PostsGet(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAll(models.Posts)
	utils.ToJSON(w, posts, http.StatusOK)
}

func PostPost(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var post models.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewPost(post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "Post created success!", http.StatusCreated)
}

func PostDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := models.Delete(models.Posts, id)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusOK)
}
