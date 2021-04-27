package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	"github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/api/models"
	_ "github.com/btcthirst/practical-tasks-nix/rest-api-echo-swagger/docs"
)

// @title My API for practice
// @version 1.0
// @description This is a sample server Posts server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email btcthirst@gmail.com

// @ license.name Apache 2.0
// @ license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /

func main() {
	e := echo.New()
	//DB init
	models.NewDB()
	//models.AutoMigrations()

	//start page
	e.GET("/", startPage)

	//swagger page
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":9000"))
}

//start page

func startPage(c echo.Context) error {

	return c.String(http.StatusOK, "Hello Someone!")
}
