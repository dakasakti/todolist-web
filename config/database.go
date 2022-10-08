package config

import (
	"fmt"

	"github.com/dakasakti/todolist-web/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMySQL(config AppConfig) *gorm.DB {
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

func InitPostgreSQL(config AppConfig) *gorm.DB {
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DB_Address,
		config.DB_Username,
		config.DB_Password,
		config.DB_Name,
		config.DB_Port,
	)

	db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting to database", err)
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{}, &entities.Post{})
}
