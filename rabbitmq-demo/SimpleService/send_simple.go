package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@192.168.31.227:5672/")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}
	queue, err := c.QueueDeclare("ferry", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	err = c.Publish("", queue.Name, false, false,
		amqp.Publishing{
			ContentType: "test/plain",
			Body:        []byte("test001"),
		})

	if err != nil {
		log.Fatal(err)
	}

}
