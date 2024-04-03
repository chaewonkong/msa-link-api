package link

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handler represents the handler for link
type Handler struct {
	r *Repository
}

// NewHandler creates a new link handler
func NewHandler(r *Repository) *Handler {
	return &Handler{r}
}

// HandleLinkAdd handles adding a new link
func (h *Handler) HandleLinkAdd(c echo.Context) error {
	l := new(AddPayload)
	if err := c.Bind(l); err != nil {
		return c.String(http.StatusBadRequest, "Invalid payload")
	}

	// save link to database
	return c.JSON(http.StatusOK, "ok")
}
