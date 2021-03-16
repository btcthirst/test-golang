package models

func AutoMigrations() {
	db := Connect()
	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}
	if db.Migrator().HasTable(&Post{}) {
		db.Migrator().DropTable(&Post{})
	}
	if db.Migrator().HasTable(&Comment{}) {
		db.Migrator().DropTable(&Comment{})
	}
	db.Debug().AutoMigrate(&User{}, &Post{}, &Comment{})
}
