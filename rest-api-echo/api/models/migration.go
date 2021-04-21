package models

func AutoMigrations() {
	db := Connect()
	if db.Migrator().HasTable(&UserDB{}) {
		db.Migrator().DropTable(&UserDB{})
	}
	//if db.Migrator().HasTable(&Post{}) {
	//	db.Migrator().DropTable(&Post{})
	//}
	//if db.Migrator().HasTable(&Comment{}) {
	//	db.Migrator().DropTable(&Comment{})
	//}
	db.Debug().AutoMigrate(&UserDB{} /*, &Post{}, &Comment{}*/)
}
