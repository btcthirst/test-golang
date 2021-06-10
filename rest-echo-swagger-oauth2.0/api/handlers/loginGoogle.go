package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	//TODO: randomize it
	randomState = "random"
)

func InitEnvGoogle() {
	err := godotenv.Load("api/.env")
	if err != nil {
		log.Fatal("Erorr load .env file")
	}
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/api/v1/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	fmt.Println("init go")
}

func Login(c echo.Context) error {
	InitEnvGoogle()
	fmt.Println("login page", googleOauthConfig.AuthCodeURL(randomState)) //must be cleaned
	url := googleOauthConfig.AuthCodeURL(randomState)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackGoogle(c echo.Context) error {
	if c.FormValue("state") != randomState {
		fmt.Println("Callback 1 page") //must be cleaned
		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/")
	}
	token, err := googleOauthConfig.Exchange(c.Request().Context(), c.FormValue("code"))
	if err != nil {
		fmt.Println("no token") //must be cleaned
		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/")
	}
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("no resp") //must be cleaned
		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("no content") //must be cleaned
		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1/")
	}
	return c.JSONBlob(http.StatusOK, content)
}
