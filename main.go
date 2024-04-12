package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/chaewonkong/msa-link-api/config"
	"github.com/chaewonkong/msa-link-api/infrastructure"
	"github.com/chaewonkong/msa-link-api/link"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.NewAppConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// RabbitMQ
	queueURL := infrastructure.QueueURL(cfg.Queue.User, cfg.Queue.Password, cfg.Queue.Host, cfg.Queue.Port)
	queueConn := infrastructure.NewQueue(queueURL)

	// DB
	dsn := infrastructure.DSN(
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.Port,
		cfg.DB.SSLMode,
		cfg.DB.TimeZone,
	)
	db := infrastructure.NewDB(dsn)

	err := db.AutoMigrate(&link.Link{})
	if err != nil {
		log.Fatal(err)
	}

	r := link.NewRepository(db)
	h := link.NewHandler(r, queueConn, logger)

	e.GET("/alive", h.HandleHealthCheck)
	e.POST("/link", h.HandleLinkAdd)
	e.PATCH("/link", h.HandleLinkUpdate)

	e.Logger.Fatal(e.Start(":8080"))
}
