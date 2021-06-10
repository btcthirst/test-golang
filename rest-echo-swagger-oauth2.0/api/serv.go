package api

import (
	"html/template"
	"io"

	"github.com/btcthirst/practical-tasks-nix/rest-echo-swagger-oauth2.0/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func StartServ() {
	e := echo.New()

	api := e.Group("/api/v1")
<<<<<<< Updated upstream
=======

	t := &Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.Renderer = t
	api.GET("/", handlers.HomeHandl)

	/////////////////Basic Auth///////////////////
	api.GET("/loginBasic", handlers.LoginBasic, middleware.BasicAuth(handlers.MySimpleAuth))

	api.GET("/helloAuth", handlers.HelloAuth)
	/////////////////Basic Auth///////////////////

	/////// TEST OAUTH 2.0 ////////////////

	api.GET("/loginGoogle", handlers.Login)
	api.GET("/callback", handlers.CallbackGoogle)

	api.GET("/loginFacebook", handlers.LoginFacebook)
	api.GET("/callbackFacebook", handlers.CallbackFacebook)

	/////// TEST OAUTH 2.0 ////////////////
	//users
>>>>>>> Stashed changes
	users := api.Group("/users")
	users.POST("/", handlers.CreateUser)
	users.PUT("/:id", handlers.UpdateUser)
	users.DELETE("/:id", handlers.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
