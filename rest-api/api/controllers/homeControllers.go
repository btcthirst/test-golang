package controllers

import (
	"net/http"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utilites"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	//test data
	text := "my home page"

	utilites.ToJSON(w, text, http.StatusOK)
}
