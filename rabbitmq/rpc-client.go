package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := bodyFrom(os.Args)
	log.Printf(" [x] Requesting fib(%d) \n ", n)
	res, err := fibonacciRPC(n)
	failOnError(err, "Failed to handle RPC request.")
	log.Printf(" [.] Got %d", res)
}

func bodyFrom(args []string) int {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnError(err, "Failed to convert args to integer.")
	return n
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.99.100:5672")
	failOnError(err, "Failed to Connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create a channel.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue.")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer.")
	corrId := randomString(32)

	err = ch.Publish(
		"",          // Exchange
		"rpc_queue", // routing key
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		},
	)
	failOnError(err, "Failed to publish a message.")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err := strconv.Atoi(string(d.Body))
			log.Println("res  : ", res)
			failOnError(err, "Failed to convert body to integer.")
			break
		}
	}
	return
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
