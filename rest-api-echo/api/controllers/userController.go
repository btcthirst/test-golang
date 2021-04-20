package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, I am a User!")
}
