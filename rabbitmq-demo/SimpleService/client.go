package main

import (
	"fmt"
	"log"
	"main/rabbitmq-demo/SimpleService/AppInit"
)

func main() {
	conn := AppInit.GetConn()
	defer conn.Close()
	c, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := c.Consume("ferry", "c1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs {
		msg.Ack(false)
		fmt.Println(msg.DeliveryTag, string(msg.Body))
	}
}
