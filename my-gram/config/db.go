package config

import (
	"fmt"
	"my-gram/model/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
