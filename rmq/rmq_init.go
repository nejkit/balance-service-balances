package rmq

import (
	"context"
	"sync"

	proto "github.com/nejkit/processing-proto/balances"
	amqp "github.com/rabbitmq/amqp091-go"
	googleProtoUtil "google.golang.org/protobuf/proto"
)

type con struct {
	Connection *amqp.Connection
}

var conInstance *con

func getConnection() *con {
	var lock = &sync.Mutex{}
	if conInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		amqpCon, err := amqp.Dial("amqp://admin:admin@message-broker:5672/")
		if err != nil {
			panic(err.Error())
		}
		conInstance = &con{
			Connection: amqpCon,
		}
		return conInstance
	}
	return conInstance
}

func getCh() *amqp.Channel {
	con := getConnection().Connection

	ch, err := con.Channel()

	if err != nil {
		panic(err.Error())
	}

	return ch
}

func InitRmq() {
	ch := getCh()
	defer ch.Close()

	ch.ExchangeDeclare(
		"e.balances.forward",
		"topic",
		true,
		false,
		false,
		false,
		nil)

	balanceGetWalletInfoRequestQ, _ := ch.QueueDeclare(
		"q.balances.request.GetWalletInfoRequest",
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
		"q.balances.request.EmmitBalanceRequest",
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
}

func InitListener(queueName string) <-chan amqp.Delivery {

	ch := getCh()

	listener, _ := ch.Consume(
		queueName,
		"balance-service",
		true,
		false,
		false,
		false,
		nil)

	return listener

}

func SendResponseGetWalletInfo(response *proto.GetWalletInfoResponse) {
	ch := getCh()
	body, err := googleProtoUtil.Marshal(response)
	if err != nil {
		return
	}

	ch.PublishWithContext(
		context.Background(),
		"e.balances.forward",
		"r.balances.balance-service.GetWalletInfoResponse.#",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}
