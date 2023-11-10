package rmq

import "github.com/rabbitmq/amqp091-go"

func InitListener(queueName string, channel *amqp091.Channel) <-chan amqp091.Delivery {
	listener, _ := channel.Consume(
		queueName,
		"balance-service",
		true,
		false,
		false,
		false,
		nil)

	return listener
}
