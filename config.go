package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// DBConfig represents the configuration for the application
type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
	TimeZone string `json:"timezone"`
}

// NewDBConfig creates a new DBConfig
func NewDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
}
