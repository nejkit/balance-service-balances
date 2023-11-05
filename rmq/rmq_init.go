package rmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRmq() {
	con, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err.Error())
	}

	ch, err := con.Channel()

	if err != nil {
		panic(err.Error())
	}

	err = ch.ExchangeDeclare(
		"e.balances.forward",
		"topic",
		true,
		false,
		false,
		false,
		nil)

	balanceEmmitRequestQ, err := ch.QueueDeclare(
		"q.balances.request.EmmitBalanceRequest",
		true,
		false,
		false,
		false,
		nil)

	err = ch.QueueBind(
		balanceEmmitRequestQ.Name,
		"r.balances.#.EmmitBalanceRequest.#",
		"e.balances.forward",
		false,
		nil)

	balanceEmmitResponceQ, err := ch.QueueDeclare(
		"q.balances.response.EmmitBalanceResponse",
		true,
		false,
		false,
		false,
		nil)

	err = ch.QueueBind(
		balanceEmmitResponceQ.Name,
		"r.balances.#.EmmitBalanceResponse.#",
		"e.balances.forward",
		false,
		nil)
}

func InitListener() <-chan amqp.Delivery {
	con, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		panic(err.Error())
	}

	ch, err := con.Channel()

	if err != nil {
		panic(err.Error())
	}

	listener, err := ch.Consume(
		"q.balances.request.EmmitBalanceRequest",
		"",
		true,
		false,
		false,
		false,
		nil)

	return listener

}
