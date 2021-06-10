package handlers

import (
	"net/http"

	"github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/models"
	"github.com/labstack/echo/v4"
)

func LoginBasic(c echo.Context) error {
	url := "/api/v1/users/"

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func MySimpleAuth(userName, password string, c echo.Context) (bool, error) {
	/*
		if userName == "noname" && password == "nopass" {
			return false, nil
		}
	*/
	u := models.User{}
	models.DB.Where("name=?", userName).First(&u)
	if userName == u.Name && password == u.Password {
		return true, nil
	}
	return false, nil
}
