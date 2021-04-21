package controllers

import (
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/models"
	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	data := getPostsAll()
	if len(data) == 0 {
		return c.String(http.StatusOK, "there is no posts, you must to create one")
	}
	return c.JSON(http.StatusOK, data)
}

func GetPostByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	data := getPostByID(id)
	return c.JSON(http.StatusOK, data)
}

//POST method
func CreatePost(c echo.Context) error {
	p := new(models.Post)
	c.Bind(&p)
	models.NewPost(p)
	return c.JSON(http.StatusOK, "post created")
}

func PutPost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	p := new(models.Post)
	c.Bind(&p)

	data, err := updatePost(id, p)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	return c.JSON(http.StatusOK, data)
}

func DeletePost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	data := deletePost(id)
	return c.JSON(http.StatusOK, data)
}
