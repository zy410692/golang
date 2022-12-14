package Lib

import (
	"log"
	"main/rabbitmq-demo/SimpleService/AppInit"

	"github.com/streadway/amqp"
)

const (
	QUEUE_NEWUSER = "newuser"
)

type MQ struct {
	Channel *amqp.Channel
}

func NewMq() *MQ {
	c, err := AppInit.GetConn().Channel()
	if err != nil {
		log.Println(err)
		return nil
	}
	return &MQ{Channel: c}
}

func (this *MQ) SendMessage(queueName string, message string) error {

	_, err := this.Channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}
	return this.Channel.Publish("", queueName, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
