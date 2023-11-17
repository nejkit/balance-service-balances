package amqp

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type ParserFunc[T any] func([]byte) (*T, error)
type Handle[T any] func(context.Context, *T)

type AmqpListener[T any] struct {
	channel    *amqp091.Channel
	logger     *logrus.Logger
	Message    <-chan amqp091.Delivery
	parserFunc ParserFunc[T]
	handler    Handle[T]
}

func NewAmqpListener[T any](ctx context.Context, factory AmqpFactory, qName string, parser ParserFunc[T], handler Handle[T]) *AmqpListener[T] {
	rmqChan := factory.NewRmqChan()
	consumer, err := rmqChan.Consume(qName, "", false, false, false, false, nil)
	if err != nil {
		factory.logger.Warningln("Fail create consumer!")
		return nil
	}
	return &AmqpListener[T]{
		channel:    rmqChan,
		logger:     factory.logger,
		Message:    consumer,
		parserFunc: parser,
		handler:    handler,
	}
}

func (l *AmqpListener[T]) Run(ctx context.Context) {
	for {
		select {
		case msg, ok := <-l.Message:
			if !ok {
				l.logger.Errorln("Consumer failed!")
			}
			body, err := l.parserFunc(msg.Body)

			if err != nil {
				l.logger.Infoln("Message unsuccessfully parsed")
				msg.Nack(false, false)
				continue // Skip processing if parsing fails
			}
			msg.Ack(false)
			go l.handler(ctx, body)

		case <-ctx.Done():
			// Context canceled, clean up and return
			l.channel.Close()
			return

		default:
			// No message received, wait briefly before checking again
			time.Sleep(100 * time.Millisecond)
		}
	}
}
