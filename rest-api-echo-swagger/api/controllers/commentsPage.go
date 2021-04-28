package controllers

import (
	"net/http"
	"strconv"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
	"github.com/labstack/echo/v4"
)

// GetComments godoc
// @Summary Get all Comments
// @Description Get all Comments fron DB if exist
// @Tags Comments
// @Produce json/string
// @Success 200 {object} models.Comment
// @Router /comments/ [get]
func GetComments(c echo.Context) error {
	data := getCommentsAll()
	if len(data) == 0 {
		return c.JSON(http.StatusOK, "there is no comments")
	}
	return c.JSON(http.StatusOK, data)
}

// GetCommentByID godoc
// @Summary Retrieves Comment based on given ID
// @Tags Comments
// @Produce json
// @Param id path integer true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func GetCommentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	data := getCommentByID(id)
	return c.JSON(http.StatusOK, data)

}

// PostComment godoc
// @Summary Create a Comment
// @Description Create a new Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 201 {object} models.Comment
// @Router /comments/ [post]
func CreateComment(c echo.Context) error {

	com := new(models.Comment)
	c.Bind(&com)

	models.NewComment(com)
	return c.JSON(http.StatusOK, "comment created")

}

// PutComment godoc
// @Summary Update a Comment
// @Description Update a existing Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param id path integer true "Comment ID"
// @Success 201 {object} models.Comment
// @Router /comments/{id} [put]
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

// DeleteComment godoc
// @Summary delete a Comment
// @Description delete existing Comment
// @Tags Comments
// @Produce json
// @Param id path integer true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [delete]
func DeleteComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	data := deleteComment(id)
	return c.JSON(http.StatusOK, data)

}
