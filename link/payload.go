package link

// AddPayload represents the payload to add a new link
type AddPayload struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdatePayload represents the payload to update a link
type UpdatePayload struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	ThumbnailImage string `json:"thumbnail_image"`
}

// QueuePayload represents the payload to send to the queue
type QueuePayload struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

func (q *QueuePayload) FromEntity(l Link) {
	q.ID = l.ID
	q.URL = l.URL
}
