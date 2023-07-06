package main

import (
	"fmt"
	"main/rabbitmq-demo/Lib"
	"time"

	"github.com/streadway/amqp"
)

func SendMail(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		fmt.Printf("向userID=%s的用户发送邮件\n", string(msg.Body))
		time.Sleep(time.Second * 1)
		msg.Ack(false)
	}

}

func main() {
	mq := Lib.NewMq()
	mq.Consume(Lib.QUEUE_NEWUSER, "c1", SendMail)

	defer mq.Channel.Close()

}
