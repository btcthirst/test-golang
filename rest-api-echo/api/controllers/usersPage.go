package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {

	data := getUsersAll()
	if len(data) == 0 {
		return c.String(http.StatusOK, "there is no users, you must to create someone")
	}

	return c.JSON(http.StatusOK, data)
}

func GetUsersByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, "errrorrrr")
	}

	data := getUsersByID(id)
	if data.ID == 0 {
		text := fmt.Sprintf("there is no user with id : %v", id)
		return c.String(http.StatusOK, text)
	}
	return c.JSON(http.StatusOK, data)
}

func PostUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	models.NewUser(user)

	return c.JSON(http.StatusOK, "user created")
}

func PutUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, "errrorrrr")
	}
	user := new(models.User)
	c.Bind(user)

	d, err := updateUser(id, user)
	if err != nil {
		return err
	}
	if d == "nothing" {
		return c.JSON(http.StatusOK, "nothing to update")
	}
	data := fmt.Sprintf("the user %v updated", d)
	return c.JSON(http.StatusOK, data)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, "errrorrrr")
	}

	data := deleteUser(id)

	return c.JSON(http.StatusOK, data)
}
