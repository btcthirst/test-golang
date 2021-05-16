package handlers

import (
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {

	var (
		userIn = new(models.UserIn)
		user   models.User
	)

	c.Bind(&userIn)
	if userIn.Name == "" {
		return c.JSON(http.StatusBadRequest, "user is empty")
	}
	user.Name = userIn.Name
	user.Email = userIn.Email
	user.Password = userIn.Password
	s, err := user.Create()
	if err != nil {
		return c.JSON(http.StatusNotImplemented, err)
	}
	return c.JSON(http.StatusCreated, s)
}

func UpdateUser(c echo.Context) error {

	var (
		userIn = new(models.UserIn)
		user   models.User
	)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "can`t update")
	}

	c.Bind(&userIn)
	if userIn.Name == "" {
		return c.JSON(http.StatusBadRequest, "user is empty")
	}
	user.ID = id
	user.Name = userIn.Name
	user.Email = userIn.Email
	user.Password = userIn.Password
	s, err := user.Update()
	if err != nil {
		return c.JSON(http.StatusNotModified, err)
	}
	return c.JSON(http.StatusCreated, s)
}

func DeleteUser(c echo.Context) error {

	var (
		user models.User
	)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "can`t update")
	}

	user.ID = id

	s, err := user.Delete()
	if err != nil {
		return c.JSON(http.StatusNotModified, err)
	}
	return c.JSON(http.StatusOK, s)
}
