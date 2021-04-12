package models

const (
	//Users ..
	Users = "users"
	//Posts ...
	Posts = "posts"
	//Feedbacks ...
	Comments = "comments"
)

func GetAll(table string) interface{} {
	var (
		userAll    []User
		postAll    []Post
		commentAll []Comment
	)
	db := Connect()
	switch table {
	case Users:
		db.Find(&userAll)
		return userAll
	case Posts:
		db.Find(&postAll)
		return postAll
	case Comments:
		db.Find(&commentAll)
		return commentAll
	}
	return nil
}

func GetByID(table string, id uint64) interface{} {
	var (
		user    User
		post    Post
		comment Comment
	)
	db := Connect()
	switch table {
	case Users:
		db.Where("id=?", id).Find(&user)
		return user
	case Posts:
		db.Where("id=?", id).Find(&post)
		return post
	case Comments:
		db.Where("id=?", id).Find(&comment)
		return comment
	}
	return nil
}
