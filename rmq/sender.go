package rmq

import (
	"context"
	"sync"

	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func SendWalletInfo(response <-chan *balances.GetWalletInfoResponse, channel *amqp091.Channel, logger *logrus.Logger) {
	lock := &sync.Mutex{}
	forever := make(chan bool)
	for msg := range response {
		lock.Lock()
		body, err := proto.Marshal(msg)
		if err != nil {
			logger.Fatal(err.Error())
		}

		err = channel.PublishWithContext(
			context.Background(),
			"e.balances.forward",
			"r.balances.balance-service.GetWalletInfoResponse.#",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		if err != nil {
			logger.Fatal(err.Error())
		} else {
			logger.Info("Send successfuly!")
		}
		lock.Unlock()
	}
	<-forever
}
