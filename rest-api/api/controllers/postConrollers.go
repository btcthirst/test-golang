package controllers

import (
	"net/http"
	"text/template"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./api/controllers/templates/postPage.html", "./api/controllers/templates/header.html", "./api/controllers/templates/footer.html")
	if err != nil {
		panic(err)
	}

	posts := models.GetAll(models.Posts)

	t.ExecuteTemplate(w, "post", posts)
}
