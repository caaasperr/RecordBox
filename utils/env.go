package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var LASTFM_API_KEY string
var DB_NAME string
var DB_PASSWORD string
var DB_TABLE string

func SetLastFmEnv() {
	err := godotenv.Load()
	if err != nil {
		Logger.Warn(" - Error loading .env file: ", err)
	}

	apiKey := os.Getenv("LASTFM_API_KEY")
	if apiKey == "" {
		Logger.Warn(" - LASTFM_API_KEY is not set in .env file")
	}
	LASTFM_API_KEY = apiKey
}

func SetDBEnv() {
	err := godotenv.Load()
	if err != nil {
		Logger.Warn(" - Error loading .env file: ", err)
	}

	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbTable := os.Getenv("DB_TABLE")

	if dbName == "" {
		Logger.Warn(" - DB_NAME is not set in .env file")
	}
	if dbPassword == "" {
		Logger.Warn(" - DB_PASSWORD is not set in .env file")
	}
	if dbTable == "" {
		Logger.Warn(" - DB_TABLE is not set in .env file")
	}
	DB_NAME = dbName
	DB_PASSWORD = dbPassword
	DB_TABLE = dbTable
}
