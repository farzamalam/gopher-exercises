package main

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", msg, err)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	failOnError(err, "Failed to Connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to Create a queue.")

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set Qos.")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer.")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer.")

			log.Printf("[.] fib(%d) \n ", n)
			response := fib(n)

			err = ch.Publish(
				"",
				"rpc_queue",
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					ReplyTo:       q.Name,
					Body:          []byte(strconv.Itoa((response))),
				},
			)
			failOnError(err, "Failed to publish a message.")
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Awaiting RPC requests.")
	<-forever
}
