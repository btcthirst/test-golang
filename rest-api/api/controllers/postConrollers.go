package controllers

import (
	"net/http"
	"text/template"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./api/controllers/templates/postPage.html", "./api/controllers/templates/header.html", "./api/controllers/templates/footer.html")
	if err != nil {
		panic(err)
	}

	//test data
	P := "more more test text"

	t.ExecuteTemplate(w, "post", P)
}
