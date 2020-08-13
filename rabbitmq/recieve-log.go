package main

import (
	"log"

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
		"logs",   // exchange name.
		"fanout", // exchange type
		true,     // durable
		false,
		false,
		false,
		nil,
	)

	// Declare Queue.
	q, err := ch.QueueDeclare(
		"", // name
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to create the queue.")

	// Bind Queue.
	err = ch.QueueBind(
		q.Name,
		"", // routing key
		"logs",
		false,
		nil,
	)
	failOnError(err, "Failed to bind the Queue")

	// Consume messages.
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // cosumer
		true,   //auto-ack
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume.")
	// Print Consumed messages.
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recieved [x] %s\n ", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for the logs. To exit press CTRL + C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}
