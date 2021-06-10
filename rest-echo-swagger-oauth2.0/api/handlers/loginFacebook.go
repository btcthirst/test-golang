package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	oauthConf        *oauth2.Config
	oauthStateString = "thisshouldberandom"
)

func InitEnvFacebook() {
	err := godotenv.Load("api/.env")
	if err != nil {
		log.Fatal("Erorr load .env file")
	}
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/api/v1/callbackFacebook",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}

}

func LoginFacebook(c echo.Context) error {
	InitEnvFacebook()
	Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	Url.RawQuery = parameters.Encode()
	url := Url.String()
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackFacebook(c echo.Context) error {
	state := c.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)

		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1")
	}

	code := c.FormValue("code")

	token, err := oauthConf.Exchange(c.Request().Context(), code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)

		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1")
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)

		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1")
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)

		return c.Redirect(http.StatusTemporaryRedirect, "/api/v1")
	}

	log.Printf("parseResponseBody: %s\n", string(response))
	return c.JSONBlob(http.StatusOK, response)
}
