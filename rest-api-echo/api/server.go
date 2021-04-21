package api

import (
	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func listenServ() {
	e := echo.New()
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

	e.Logger.Fatal(e.Start(":9000"))
}
