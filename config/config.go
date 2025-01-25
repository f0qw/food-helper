package config

import (
	"fmt"
	"food-helper/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
import "gorm.io/driver/mysql"

func InitDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/food?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 初始化表结构
	db.AutoMigrate(&models.Dish{}, &models.Element{})

	return db, nil
}
