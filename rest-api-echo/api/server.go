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

	e.Logger.Fatal(e.Start(":9000"))
}
