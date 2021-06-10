package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandl(c echo.Context) error {

	return c.Render(http.StatusOK, "loginPage", nil)
}

func HelloAuth(c echo.Context) error {
	return c.String(http.StatusOK, "you are logined")
}
