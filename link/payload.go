package link

// AddPayload represents the payload to add a new link
type AddPayload struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
