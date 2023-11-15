package amqp

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type AmqpSender struct {
	channel *amqp091.Channel
	logger  *logrus.Logger
}

func NewAmqpSender(ch *amqp091.Channel, log *logrus.Logger) AmqpSender {
	return AmqpSender{channel: ch, logger: log}
}

func (s *AmqpSender) SendMessage(ctx context.Context, body protoreflect.ProtoMessage) {
	bytes, err := proto.Marshal(body)
	if err != nil {
		s.logger.Warningln("Marchall message with error: ", err.Error())
		return
	}
	err = s.channel.PublishWithContext(ctx, "", "", false, false, amqp091.Publishing{ContentType: "text/plain", Body: bytes})
	if err != nil {
		s.logger.Warningln("Message not published. Error: ", err.Error())
	}
}
