package rmq

import (
	"balance-service/statics"

	amqp "github.com/rabbitmq/amqp091-go"
)

func getChannel(connectionString string) *amqp.Channel {
	con, err := amqp.Dial(connectionString)
	if err != nil {

	}
	ch, _ := con.Channel()
	return ch
}

func InitRmq(connectionString string) *amqp.Channel {
	ch := getChannel(connectionString)

	ch.ExchangeDeclare(
		"e.balances.forward",
		"topic",
		true,
		false,
		false,
		false,
		nil)

	balanceGetWalletInfoRequestQ, _ := ch.QueueDeclare(
		statics.GetWalletInfoRequestQueue,
		true,
		false,
		false,
		false,
		nil)

	balanceGetWalletInfoResponse, _ := ch.QueueDeclare(
		"q.balances.response.GetWalletInfoResponse",
		true,
		false,
		false,
		false,
		nil)

	balanceEmmitRequestQ, _ := ch.QueueDeclare(
		statics.EmmitBalanceRequestQueue,
		true,
		false,
		false,
		false,
		nil)

	ch.QueueBind(
		balanceEmmitRequestQ.Name,
		"r.balances.#.EmmitBalanceRequest.#",
		"e.balances.forward",
		false,
		nil)

	ch.QueueBind(
		balanceGetWalletInfoRequestQ.Name,
		"r.balances.#.GetWalletInfoRequest.#",
		"e.balances.forward",
		false,
		nil)

	ch.QueueBind(
		balanceGetWalletInfoResponse.Name,
		"r.balances.#.GetWalletInfoResponse.#",
		"e.balances.forward",
		false,
		nil)

	return ch
}
