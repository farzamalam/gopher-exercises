package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Create a connection.
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	failOnError(err, "Failed to Connect to RabbitMQ.")
	defer conn.Close()

	// Create a channel.
	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel.")
	defer ch.Close()

	// Declare an exchange.
	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to Create an Exchange.")

	// Declare a Queue.
	q, err := ch.QueueDeclare(
		"", // name
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to create a queue.")

	// Bind Queue.
	if len(os.Args) < 2 {
		log.Fatalf("Usage : %s [binding_key] ... ", os.Args[0])
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s ", q.Name, "logs_topic", s)
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)

		failOnError(err, "Failed to bind a queue.")

	}
	// Consume message.
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume the message.")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s\n ", d.Body)
		}
	}()
	log.Printf("[*] Waiting for the logs. Press CTRL + C to exit.")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}
