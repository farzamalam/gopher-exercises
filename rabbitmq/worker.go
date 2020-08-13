package main

import (
	"bytes"
	"log"
	"time"

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

	// Create Queue.
	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	// Prefetch
	err = ch.Qos(
		1,
		0,
		false,
	)

	// Consume messages.
	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // Message Acknowledgement
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume the msgs.")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Recieved a message %s \n", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done.")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for the messages. To exit press CTRL + C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}
