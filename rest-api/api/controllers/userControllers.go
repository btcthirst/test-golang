package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
	"github.com/gorilla/mux"
)

func UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	user := models.GetByID(models.Users, id)
	utils.ToJSON(w, user, http.StatusOK)

	//end xml
	//utils.ToXML(w, user, http.StatusOK)
}

func UsersGet(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll(models.Users)
	utils.ToJSON(w, users, http.StatusOK)

	//end xml
	//utils.ToXML(w, users, http.StatusOK)
}

func UserPost(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewUser(user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "User created success!", http.StatusCreated)
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := models.Delete(models.Users, id)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, rows, http.StatusOK)
}
