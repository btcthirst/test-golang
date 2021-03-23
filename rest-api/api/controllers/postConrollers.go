package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/btcthirst/practical-tasks-nix/rest-api/api/models"
	"github.com/btcthirst/practical-tasks-nix/rest-api/api/utils"
)

func PostPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetPosts(w, r)
	case http.MethodPost:
		PostPost(w, r)
	case http.MethodDelete:
		DeletePost(w, r)
	case http.MethodPut:
		PutPost(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsuported method %v, to %v\n", r.Method, r.URL)
		log.Printf("Unsuported method %v, to %v\n", r.Method, r.URL)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAll(models.Posts)

	w.Header().Set("Content-Type", "application/json")
	je := json.NewEncoder(w)
	je.Encode(posts)

}

func PostPost(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var post models.Post
	err := json.Unmarshal(body, &post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = models.NewPost(post)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(w, "Post created success!", http.StatusCreated)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to delete post %v", id)

	db := models.Connect()
	var posts []models.Post

	db.Find(&posts)
	for _, post := range posts {
		if post.ID == id {
			db.Delete(&post)
			log.Print("Success!!!")
			fmt.Fprintf(w, " the post id: %v deleted success!", id)
		}
	}
}

func PutPost(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("Request to put post %v", id)

	db := models.Connect()
	var posts []models.Post

	db.Find(&posts)

	body := utils.BodyParser(r)
	var post1 models.Post
	err = json.Unmarshal(body, &post1)
	if err != nil {
		utils.ToJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	for _, post := range posts {
		if post.ID == id {
			if post.Title != post1.Title {
				db.Model(&post).Update("title", post1.Title)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")
			} else if post.ImageURL != post1.ImageURL {
				db.Model(&post).Update("image_url", post1.ImageURL)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")

			} else if post.Body != post1.Body {
				db.Model(&post).Update("password", post1.Body)
				log.Print("Success!!!")
				fmt.Fprint(w, "Success!!!")
			} else {
				fmt.Fprint(w, "there are nothing to change")
			}

		}
	}
}
