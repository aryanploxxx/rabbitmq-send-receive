// Consumer

package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Setting up is the same as the publisher; we open a connection and a channel, and declare the queue from which we're going to consume.
	// This matches up with the queue that send publishes to.
	conn, err := amqp.Dial("amqp://user:1234@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer channel.Close()

	// We declare the queue here, as well. Because we might start the consumer before the publisher, we want to make sure the queue exists before we try to consume messages from it.
	queue, err := channel.QueueDeclare(
		"hello", // name of the queue
		false,   // durable (queue will survive a broker restart)
		false,   // delete when unused (queue will be deleted when no longer used)
		false,   // exclusive (queue will only be used by the connection that declared it)
		false,   // no-wait (do not wait for a server response)
		nil,     // arguments (optional arguments)
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err) // Logging and exiting if queue declaration fails
	}

	// We're about to tell the server to deliver us the messages from the queue. Since it will push us messages asynchronously, we will read the messages from a channel (returned by amqp::Consume) in a goroutine.
	msgs, err := channel.Consume(
		queue.Name, // queue name
		"",         // consumer tag (empty string means a unique tag will be generated)
		true,       // auto-ack (automatically acknowledge messages)
		false,      // exclusive (only this consumer can access the queue)
		false,      // no-local (if true, the server will not deliver messages to the connection that published them)
		false,      // no-wait (do not wait for a server response)
		nil,        // args (optional arguments)
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err) // Logging and exiting if consumer registration fails
	}

	var forever = make(chan struct{}) // Creating a channel to block the main function

	go func() {
		for d := range msgs { // Iterating over messages as they arrive
			log.Printf("Received a message: %s", d.Body) // Logging the received message
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C") // Informing the user that the consumer is waiting for messages
	<-forever                                                     // Blocking the main function to keep it running
}
