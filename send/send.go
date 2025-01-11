// Producer

package main

import (
	"context" // Package for managing context and timeouts
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go" // RabbitMQ client package
)

func main() {
	// connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:1234@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us.
	// Next we create a channel, which is where most of the API for getting things done resides:
	channel, err := conn.Channel() // Open a channel to communicate with RabbitMQ
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// we must declare a queue for us to send to; then we can publish a message to the queue
	// Note: A queue will only be created if it doesn't exist already.
	queue, err := channel.QueueDeclare(
		"hello", // name of the queue
		false,   // durable (queue will survive a broker restart if true)
		false,   // delete when unused
		false,   // exclusive (used by only one connection and the queue will be deleted when that connection closes)
		false,   // no-wait (do not wait for a server response)
		nil,     // arguments (optional additional arguments)
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	context, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Create a context with a 3-second timeout
	defer cancel()

	contentSentByPublisher := "This message is being published by Publisher."

	err = channel.PublishWithContext(
		context,    // Context for managing timeout
		"",         // exchange (default exchange)
		queue.Name, // routing key (queue name)
		false,      // mandatory (if true, the server will return an unroutable message)
		false,      // immediate (if true, the server will return an undeliverable message)
		amqp.Publishing{
			ContentType: "text/plain",                   // Content type of the message
			Body:        []byte(contentSentByPublisher), // Message body as a byte array
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf("Message Published: %s\n", contentSentByPublisher)
}
