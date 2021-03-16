package routers

import (
	"net/http"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/controllers"
)

func MyHandlers() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/users", controllers.UserPage)
	http.HandleFunc("/posts", controllers.PostPage)
	http.HandleFunc("/comments", controllers.CommentPage)
}
