package config

import "os"

// App represents the configuration for the application
type App struct {
	DB    db    `json:"db"`
	Queue queue `json:"queue"`
}

type (
	db struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		Port     string `json:"port"`
		SSLMode  string `json:"sslmode"`
		TimeZone string `json:"timezone"`
	}

	queue struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Port     string `json:"port"`
	}
)

// NewAppConfig creates a new DBConfig
func NewAppConfig() *App {
	db := db{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
	queue := queue{
		Host:     os.Getenv("QUEUE_HOST"),
		User:     os.Getenv("QUEUE_USER"),
		Password: os.Getenv("QUEUE_PASSWORD"),
		Port:     os.Getenv("QUEUE_PORT"),
	}

	return &App{db, queue}
}
