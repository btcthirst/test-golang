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

func UserPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		PostUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	case http.MethodPut:
		PutUser(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsuported method %v, to %v\n", r.Method, r.URL)
		log.Printf("Unsuported method %v, to %v\n", r.Method, r.URL)

	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := models.GetAll(models.Users)
	je := json.NewEncoder(w)
	je.Encode(users)
}

func PostUser(w http.ResponseWriter, r *http.Request) {

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to delete user %v", id)

	db := models.Connect()
	var users []models.User

	db.Find(&users)
	for _, user := range users {
		if user.ID == id {
			db.Delete(&user)
			log.Print("Success!!!")
		}
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to put user %v", id)

	db := models.Connect()
	var users []models.User

	db.Find(&users)

	body := utils.BodyParser(r)
	var user1 models.User
	err = json.Unmarshal(body, &user1)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	for _, user := range users {
		if user.ID == id {
			if user.UserName != user1.UserName {
				db.Model(&user).Update("user_name", user1.UserName)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")
			} else if user.Email != user1.Email {
				db.Model(&user).Update("email", user1.Email)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")

			} else if user.Password != user1.Password {
				db.Model(&user).Update("password", user1.Password)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")
			} else {
				fmt.Fprint(w, "there are nothing to change")
			}

		}
	}

}
