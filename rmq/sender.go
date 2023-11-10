package rmq

import (
	"context"

	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func SendWalletInfo(response *balances.GetWalletInfoResponse, channel *amqp091.Channel) {
	body, err := proto.Marshal(response)
	if err != nil {
		return
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
}
