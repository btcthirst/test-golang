package routers

import (
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/controllers"
	"github.com/gorilla/mux"
)

func MyHandlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//homePage
	r.HandleFunc("/", controllers.HomePage)
	//users routers
	r.HandleFunc("/users/{id:[0-9]+}", controllers.UserGet).Methods("GET")
	r.HandleFunc("/users", controllers.UsersGet).Methods("GET")
	r.HandleFunc("/users", controllers.UserPost).Methods("POST")
	r.HandleFunc("/users", controllers.UserDelete).Methods("DELETE")
	//posts routers
	r.HandleFunc("/posts/{id:[0-9]+}", controllers.PostGet).Methods("GET")
	r.HandleFunc("/posts", controllers.PostsGet).Methods("GET")
	r.HandleFunc("/posts", controllers.PostPost).Methods("POST")
	r.HandleFunc("/posts", controllers.PostDelete).Methods("DELETE")
	//comments routers
	r.HandleFunc("/comments/{id:[0-9]+}", controllers.CommentGet).Methods("GET")
	r.HandleFunc("/comments", controllers.CommentsGet).Methods("GET")
	r.HandleFunc("/comments", controllers.CommentPost).Methods("POST")
	r.HandleFunc("/comments", controllers.CommentDelete).Methods("DELETE")

	return r
}
