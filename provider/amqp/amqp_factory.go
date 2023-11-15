package amqp

import (
	"balance-service/statics"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type AmqpFactory struct {
	logger *logrus.Logger
	con    *amqp091.Connection
}

func NewAmqpFactory(connectionString string, loggger *logrus.Logger) AmqpFactory {
	var con *amqp091.Connection
	var err error
	for {
		con, err = amqp091.Dial(connectionString)
		if err != nil {
			loggger.Errorln("Rabbit unvailable. Reazon: ", err.Error())
			continue
		}
		break
	}
	return AmqpFactory{logger: loggger, con: con}
}

func (f *AmqpFactory) NewRmqChan() *amqp091.Channel {
	ch, err := f.con.Channel()
	if err != nil {
		f.logger.Errorln("Channel not created. Error: ", err.Error())
		return nil
	}
	return ch
}

func (f *AmqpFactory) NewSender(exchange string, rk string) AmqpSender {
	ch := f.NewRmqChan()
	return AmqpSender{channel: ch, logger: f.logger}
}

func (f *AmqpFactory) NewConsumeChan(qName string, rmqChannel *amqp091.Channel) <-chan amqp091.Delivery {
	consumeChan, err := rmqChannel.Consume(
		qName, "", false, false, false, false, nil)
	if err != nil {
		f.logger.Errorln("Cannot create a consumer. Error: ", err.Error())
	}
	return consumeChan
}

func (f *AmqpFactory) ClsConnection() {
	f.con.Close()
}

func (f *AmqpFactory) InitRmq() {
	ch := f.NewRmqChan()
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
}
