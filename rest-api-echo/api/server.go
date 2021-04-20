package api

import (
	"github.com/btcthirst/practical-tasks-nix/rest-api-echo/api/controllers"
	"github.com/labstack/echo/v4"
)

func listenServ() {
	e := echo.New()
	e.GET("/users/", controllers.GetUsers)

	e.Logger.Fatal(e.Start(":9000"))
}
