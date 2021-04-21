package controllers

import (
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/models"
	"github.com/labstack/echo/v4"
)

func GetComments(c echo.Context) error {
	data := getCommentsAll()
	if len(data) == 0 {
		return c.JSON(http.StatusOK, "there is no comments")
	}
	return c.JSON(http.StatusOK, data)
}

func GetCommentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	data := getCommentByID(id)
	return c.JSON(http.StatusOK, data)

}

//POST method
func CreateComment(c echo.Context) error {

	com := new(models.Comment)
	c.Bind(&com)

	models.NewComment(com)
	return c.JSON(http.StatusOK, "comment created")

}

func PutComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	com := new(models.Comment)
	c.Bind(&com)
	data, err := updateComment(id, com)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return c.JSON(http.StatusOK, data)
}

func DeleteComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	data := deleteComment(id)
	return c.JSON(http.StatusOK, data)

}
