package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Create a Connection.
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	failOnError(err, "Failed to connect to RabbitMQ.")
	defer conn.Close()

	// Create a Channel.
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel.")
	defer ch.Close()

	// Declare Exchange.
	err = ch.ExchangeDeclare(
		"logs_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to Declare Exchange.")

	// Declare Queue.
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to create Queue")

	// Bind Queue.
	if len(os.Args) < 2 {
		log.Fatal("Use [info] [warning] [error] ")
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s with exchange %s using routing key %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_direct",
			false,
			nil,
		)
		failOnError(err, "Failed in QueueBind")
	}
	// Consume messages.
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register msgs.")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s\n ", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}
