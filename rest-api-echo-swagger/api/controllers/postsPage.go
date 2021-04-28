package controllers

import (
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
	"github.com/labstack/echo/v4"
)

// GetPosts godoc
// @Summary Get all posts
// @Description Get all posts fron DB if exist
// @Tags posts
// @Produce json/string
// @Success 200 {object} models.Post
// @Router /posts/ [get]
func GetPosts(c echo.Context) error {
	data := getPostsAll()
	if len(data) == 0 {
		return c.String(http.StatusOK, "there is no posts, you must to create one")
	}
	return c.JSON(http.StatusOK, data)
}

// GetPostByID godoc
// @Summary Retrieves post based on given ID
// @Tags posts
// @Produce json
// @Param id path integer true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func GetPostByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	data := getPostByID(id)
	return c.JSON(http.StatusOK, data)
}

// CreatePost godoc
// @Summary Create a post
// @Description Create a new Post
// @Tags posts
// @Accept json
// @Produce json
// @Success 201 {object} models.Post
// @Router /posts/ [post]
func CreatePost(c echo.Context) error {
	p := new(models.Post)
	c.Bind(&p)
	models.NewPost(p)
	return c.JSON(http.StatusOK, "post created")
}

// PutPost godoc
// @Summary Update a post
// @Description Update a existing Post
// @Tags posts
// @Accept json
// @Produce json
// @Success 201 {object} models.Post
// @Router /posts/{id}  [put]
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

// DeletePost godoc
// @Summary delete a Post
// @Description delete existing Post
// @Tags posts
// @Produce json
// @Param id path integer true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{id} [delete]
func DeletePost(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	data := deletePost(id)
	return c.JSON(http.StatusOK, data)
}
