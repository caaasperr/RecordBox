package utils

import (
	"fmt"
	"myvinyl/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() error {
	Logger.Info("Initializing database...")
	Logger.Info(" - Get database information...")
	SetDBEnv()
	Logger.Info(" - Successfully got database information")
	Logger.Info(" - Connecting database ")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_NAME, DB_PASSWORD, DB_TABLE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("  - Database Conncection Error! Check database settings...")
		os.Exit(1)
		return err
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Vinyl{})
	DB.AutoMigrate(&models.Genre{})
	DB.AutoMigrate(&models.Bookshelf{})
	DB.AutoMigrate(&models.Shelfslot{})
	Logger.Info("Database Connected Successfully!")
	return nil
}
