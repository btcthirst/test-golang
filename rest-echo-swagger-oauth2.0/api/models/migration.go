package models

import "log"

func Automigrator() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db.Migrator().DropTable(&User{})
	db.Migrator().CreateTable(&User{})

	////test create
	testCreator()
}

func testCreator() {
	///test///
	user1 := User{
		Name:     "admic",
		Email:    "admic@thirst.com",
		Password: "hardPass",
		//CreatedAt time.Time `gorm:"autoCreateTime:nano"`
		//UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	}
	user1.Create()
	user := User{
		Name:     "test1",
		Email:    "test@example.com",
		Password: "pass",
		//CreatedAt time.Time `gorm:"autoCreateTime:nano"`
		//UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
	}
	user.Create()
	///test///
}
