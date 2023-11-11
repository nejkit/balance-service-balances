package rmq

import (
	"context"
	"sync"

	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func SendWalletInfo(response <-chan *balances.GetWalletInfoResponse, channel *amqp091.Channel) {
	lock := &sync.Mutex{}
	for msg := range response {
		lock.Lock()
		body, err := proto.Marshal(msg)
		if err != nil {

		}

		channel.PublishWithContext(
			context.Background(),
			"e.balances.forward",
			"r.balances.balance-service.GetWalletInfoResponse.#",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		lock.Unlock()
	}
}
