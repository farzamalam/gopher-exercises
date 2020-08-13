package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	// Create Connection.
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	failOnError(err, "Failed to connect to RabbitMQ.")
	defer conn.Close()

	// Create Channel.
	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel.")
	defer ch.Close()

	// Declare Exchange.
	err = ch.ExchangeDeclare(
		"logs_direct", // exchange name
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare the exchange.")

	// Create msg.
	body := bodyFrom(os.Args)

	// Publish to Exchange.
	err = ch.Publish(
		"logs_direct",         // exchange name
		severityFrom(os.Args), // routing keys
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
