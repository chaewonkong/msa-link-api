package infrastructure

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// QueueURL returns a formatted string for a RabbitMQ connection.
func QueueURL(user, password, host, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
}

// NewQueue returns a new amqp.Connection instance.
func NewQueue(queueURL string) *amqp.Connection {
	// Connect to the RabbitMQ server

	queueConn, err := amqp.Dial(queueURL)
	if err != nil {
		log.Fatal(err)
	}

	return queueConn
}
