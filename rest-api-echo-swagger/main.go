package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/controllers"
	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
	_ "github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/docs"
)

// @title My API for practice
// @version 1.0
// @description This is a sample server Posts server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email btcthirst@gmail.com

// @host localhost:1982
// @BasePath /

func main() {
	e := echo.New()
	//DB init
	models.NewDB()
	models.AutoMigrations()

	//start page
	e.GET("/", startPage)

	//swagger page
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//users page
	e.GET("/users/", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUsersByID)
	e.POST("/users/", controllers.PostUser)
	e.PUT("/users/:id", controllers.PutUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	//post page
	e.GET("/posts/", controllers.GetPosts)
	e.GET("/posts/:id", controllers.GetPostByID)
	e.POST("/posts/", controllers.CreatePost)
	e.PUT("/posts/:id", controllers.PutPost)
	e.DELETE("/posts/:id", controllers.DeletePost)
	//comment
	e.GET("/comments/", controllers.GetComments)
	e.GET("/comments/:id", controllers.GetCommentByID)
	e.POST("/comments/", controllers.CreateComment)
	e.PUT("/comments/:id", controllers.PutComment)
	e.DELETE("/comments/:id", controllers.DeleteComment)

	e.Logger.Fatal(e.Start(":1982"))
}

//start page

func startPage(c echo.Context) error {

	return c.String(http.StatusOK, "Hello Someone!")
}
