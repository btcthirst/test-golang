package api

import "github.com/btcthirst/practical-tasks-nix/rest-api/api/models"

func Run() {
	//creating tables in db
	models.AutoMigrations()

	//try for create some in tables
	models.NewUser(models.User{UserName: "Testix", Email: "TestixShdw@email.com", Password: "123456"})
	models.NewUser(models.User{UserName: "Twostix", Email: "TwostixShdw@email.com", Password: "123456"})
	models.NewPost(models.Post{Title: "some description", ImageURL: "https://source.unsplash.com/random/800x600", Body: "some text without idea sub image of", UserID: 1})
	models.NewComment(models.Comment{Comment: "just test comment", UserID: 1, PostID: 1})

	listen(9000)
}
