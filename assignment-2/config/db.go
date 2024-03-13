package config

import (
	"assignment-2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	err = db.AutoMigrate(&model.Order{}, &model.Item{})
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}

	return db, nil
}
