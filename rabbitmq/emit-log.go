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
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create Channel.
	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel.")
	defer ch.Close()

	// Create Exchange.
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // exchange type
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to create an exchange.")

	// Make Payload.
	body := bodyFrom(os.Args)

	// Send message.
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

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
		log.Fatalf("%s : %s ", msg, err)
	}
}
