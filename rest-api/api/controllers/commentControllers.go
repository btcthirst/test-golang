package controllers

import (
	"net/http"
	"text/template"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
)

func CommentPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./api/controllers/templates/commentPage.html", "./api/controllers/templates/header.html", "./api/controllers/templates/footer.html")
	if err != nil {
		panic(err)
	}

	comments := models.GetAll(models.Comments)

	t.ExecuteTemplate(w, "comment", comments)
}
