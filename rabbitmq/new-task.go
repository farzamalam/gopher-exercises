package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {

	// Create a Connection
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel")
	defer ch.Close()

	// Create a Queue
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to create a queue")

	// Create payload
	body := bodyFrom(os.Args)

	// Publish message / payload
	err = ch.Publish(
		"", // Exchange,
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s ", body)
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
