package configs

import (
	"github.com/leeroyakbar/messaging-app/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {

	dsn := "host=localhost user=postgres password=Akbar_123 dbname=messaging_app_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		println("Database connected")
	}

	err = db.AutoMigrate(
		&models.Contact{},
		&models.GroupMember{},
		&models.Group{},
		&models.Media{},
		&models.Message{},
	)
	if err != nil {
		panic("Migration failed")
	} else {
		println("Migration success")
	}
}
