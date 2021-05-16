package models

import "log"

func Automigrator() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})
}
