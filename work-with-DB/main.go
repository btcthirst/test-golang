package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Articles struct {
	UserID    uint32    `json:"userId"`
	ID        uint32    `json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Body      string    `gorm:"type:text;not null" json:"body"`
	CreatedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp();not null" json:"updated_at"`
	DeletedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp()" json:"deleted_at"`
}

type Comments struct {
	PostID    uint32    `json:"postId"`
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Body      string    `json:"body"`
	CreatedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp();not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp();not null" json:"updated_at"`
	DeletedAt time.Time `gorm:"type: timestamp(0);default:current_timestamp()" json:"deleted_at"`
}

const (
	host     = "127.0.0.1"
	port     = "3306"
	database = "golang1"
	user     = "webmax"
	password = "qwerty1234#WebM4X"
)

//get data
func getData(urla, label string) {
	resp, err := http.Get(urla)
	mistakes(err)
	body, err := ioutil.ReadAll(resp.Body)
	mistakes(err)

	go FromJSON(body, label)
}

//unmarshal body
func FromJSON(b []byte, label string) {
	var (
		post []Articles
		comm []Comments
	)

	switch label {
	case "posts":
		json.Unmarshal(b, &post)
		for _, v := range post {
			urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", v.ID)

			go getData(urla, "comments")
		}
	case "comments":
		json.Unmarshal(b, &comm)
		db := conDB()
		go func() {
			db.Create(&comm)
		}()

	default:
		fmt.Println("in switch some wrong data")
	}

}

func conDB() *gorm.DB {
	//think about context

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	mistakes(err)
	return db
}

func mistakes(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func automigrator() {
	db := conDB()

	if db.Migrator().HasTable(&Articles{}) {
		db.Migrator().DropTable(&Articles{})
	}
	if db.Migrator().HasTable(&Comments{}) {
		db.Migrator().DropTable(&Comments{})
	}
	db.Debug().AutoMigrate(&Articles{}, &Comments{})

}

func main() {
	automigrator()

	var tn int64 //test number

	fmt.Scan(&tn)
	urla := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", tn)
	go getData(urla, "posts")

	fmt.Scan(&tn)

}
