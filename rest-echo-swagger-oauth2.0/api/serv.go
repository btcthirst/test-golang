package api

import (
	"github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/handlers"
	"github.com/labstack/echo/v4"
)

func StartServ() {
	e := echo.New()

	api := e.Group("/api/v1")
	users := api.Group("/users")
	users.POST("/", handlers.CreateUser)
	users.PUT("/:id", handlers.UpdateUser)
	users.DELETE("/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
