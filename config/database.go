package config

import (
	"fmt"

	"github.com/dakasakti/postingan/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config AppConfig) *gorm.DB {
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Address,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{}, &entities.Post{})
}
