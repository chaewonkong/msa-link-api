package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/chaewonkong/msa-link-api/config"
	"github.com/chaewonkong/msa-link-api/link"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	cfg := config.NewAppConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// RabbitMQ
	queue := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Queue.User, cfg.Queue.Password, cfg.Queue.Host, cfg.Queue.Port)

	queueConn, err := amqp.Dial(queue)
	if err != nil {
		log.Fatal(err)
	}
	defer queueConn.Close()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.Port,
		cfg.DB.SSLMode,
		cfg.DB.TimeZone,
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
	h := link.NewHandler(r, queueConn, logger)

	e.GET("/alive", h.HandleHealthCheck)
	e.POST("/link", h.HandleLinkAdd)

	e.Logger.Fatal(e.Start(":8080"))
}
