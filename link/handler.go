package link

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler represents the handler for link
type Handler struct {
	repo   *Repository
	logger *slog.Logger
}

// NewHandler creates a new link handler
func NewHandler(r *Repository, l *slog.Logger) *Handler {
	return &Handler{r, l}
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

	// save link to database
	return c.JSON(http.StatusOK, res)
}
