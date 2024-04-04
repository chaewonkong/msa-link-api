package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/chaewonkong/msa-link-api/link"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	cfg := NewDBConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode,
		cfg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&link.Link{})
	if err != nil {
		log.Fatal(err)
	}

	r := link.NewRepository(db)
	h := link.NewHandler(r, logger)

	e.GET("/alive", h.HandleHealthCheck)
	e.POST("/link", h.HandleLinkAdd)

	e.Logger.Fatal(e.Start(":8080"))
}
