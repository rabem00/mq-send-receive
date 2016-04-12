package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileNamePtr := flag.String("file", "", "enter a /path/filename to send")
	flag.Parse()

	dat, err := ioutil.ReadFile(*fileNamePtr)
	check(err)
	fmt.Print(string(dat))

	conn, err := amqp.Dial("amqp://guest:guest@192.168.1.134:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// The false puts the client out of NoWait mode such that the client waits
	// for an ACK or NACK after each message.
	ch.Confirm(false)

	q, err := ch.QueueDeclare(
		"zgr", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := dat
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		true,   // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/xml",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
