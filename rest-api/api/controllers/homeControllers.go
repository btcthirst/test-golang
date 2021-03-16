package controllers

import (
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./api/controllers/templates/homePage.html", "./api/controllers/templates/header.html", "./api/controllers/templates/footer.html")
	if err != nil {
		panic(err)
	}

	//test data
	P := "more more test text"

	t.ExecuteTemplate(w, "home", P)
}
