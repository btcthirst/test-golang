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

func UserPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPost:
		PostUser(w, r)
	case http.MethodPut:
		PutUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		fmt.Fprintf(w, "%v -there is no such method on this page", r.Method)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	slice := strings.Split(r.URL.String(), "/")
	if slice[len(slice)-1] != "" {
		id, err := strconv.ParseUint(slice[len(slice)-1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		users := models.GetByID(models.Users, id)
		utilites.ToJSON(w, users, http.StatusOK)
	} else {
		users := models.GetAll(models.Users)
		utilites.ToJSON(w, users, http.StatusOK)
	}

}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	defer r.Body.Close()

	//try to catch empty post request
	bug := json.NewDecoder(r.Body).Decode(&user)
	if bug != nil {
		utilites.ToJSON(w, "can`t use empty post", http.StatusOK)
	} else {
		if models.CheckUser(user) {
			models.NewUser(user)
			utilites.ToJSON(w, "user created", http.StatusOK)
		} else {

			utilites.ToJSON(w, "not uniq user", http.StatusOK)
		}

	}

	// var user models.User
	// body := utilites.BodyParser(r)
	// json.Unmarshal(body,&user)
	// models.NewUser(user)
	// utilites.ToJSON(w, "user created", http.StatusOK)

}

func PutUser(w http.ResponseWriter, r *http.Request) {
	utilites.ToJSON(w, "user udated", http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	utilites.ToJSON(w, "user deleted", http.StatusOK)
}
