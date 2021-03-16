package controllers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
)

func UserPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./api/controllers/templates/userPage.html", "./api/controllers/templates/header.html", "./api/controllers/templates/footer.html")
	if err != nil {
		panic(err)
	}

	users := models.GetAll(models.Users)
	//utils.ToJSON(w, users, http.StatusOK)

	//test data
	//P := "more more test users"

	t.ExecuteTemplate(w, "user", users)
}

func PostUser(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "User created success!", http.StatusCreated)
}
