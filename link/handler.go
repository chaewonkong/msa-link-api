package link

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/chaewonkong/msa-link/lib/transport/queue"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/labstack/echo/v4"
)

// Handler represents the handler for link
type Handler struct {
	repo   *Repository
	queue  *queue.RabbitMQ
	logger *slog.Logger
}

// NewHandler creates a new link handler
func NewHandler(r *Repository, q *queue.RabbitMQ, l *slog.Logger) *Handler {
	return &Handler{r, q, l}
}

// HandleHealthCheck handles health check
func (h *Handler) HandleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// HandleLinkAdd handles adding a new link
func (h *Handler) HandleLinkAdd(c echo.Context) error {
	l := new(AddPayload)
	if err := c.Bind(l); err != nil {
		h.logger.Error("error occurred while binding payload", err)
		return c.String(http.StatusBadRequest, "Invalid payload")
	}

	res, err := h.repo.Add(*l)
	if err != nil {
		h.logger.Error("error occurred while saving link", err)
		return c.String(http.StatusInternalServerError, "error occurred while saving link")
	}

	// convert saved res to QueuePayload
	QueuePayload := new(QueuePayload)
	QueuePayload.FromEntity(res)
	bytes, err := json.Marshal(QueuePayload)
	if err != nil {
		h.logger.Error("error occurred while marshalling payload", err)
	}

	q, err := h.queue.Ch.QueueDeclare(
		"link_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		h.logger.Error("error occurred while declaring queue", err)
	}

	err = h.queue.Ch.PublishWithContext(
		c.Request().Context(),
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
	if err != nil {
		h.logger.Error("error occurred while publishing message", err)
	}

	// save link to database
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) HandleLinkUpdate(c echo.Context) error {
	l := new(UpdatePayload)
	ctp := c.Request().Header.Get("Content-Type")
	_ = ctp
	if err := c.Bind(l); err != nil {
		h.logger.Error("error occurred while binding payload", err)
		return c.String(http.StatusBadRequest, "Invalid payload")
	}

	res, err := h.repo.Update(*l)
	if err != nil {
		h.logger.Error("error occurred while updating link", err)
		return c.String(http.StatusInternalServerError, "error occurred while updating link")
	}

	return c.JSON(http.StatusOK, res)
}
