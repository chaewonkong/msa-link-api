package main

import (
	"fmt"
	"github.com/chaewonkong/msa-link-api/link"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()
	cfg := NewDBConfig()

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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // TODO: connect postgreSQL
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&link.Link{})
	if err != nil {
		log.Fatal(err)
	}

	r := link.NewRepository(db)
	h := link.NewHandler(r)
	e.POST("/link", h.HandleLinkAdd)

	e.Logger.Fatal(e.Start(":8080"))
}
