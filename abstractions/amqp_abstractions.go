package abstractions

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AmqpSender interface {
	SendMessage(ctx context.Context, body protoreflect.ProtoMessage)
}

type AmqpFactory interface {
	NewSender(exchange string, rk string, rmqChannel *amqp091.Channel) AmqpSender
	NewRmqChan() *amqp091.Channel
	NewConsumeChan(qName string, rmqChannel *amqp091.Channel) *amqp091.Delivery
	ClsConnection()
	InitRmq()
}
