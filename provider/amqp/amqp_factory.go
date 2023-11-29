package amqp

import (
	"balance-service/external/balances"
	"balance-service/statics"
	"context"
	"fmt"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type AmqpFactory struct {
	logger *logrus.Logger
	con    *amqp091.Connection
}

func GetParserEmmitBalanceRequest() func([]byte) (*balances.EmmitBalanceRequest, error) {
	return func(body []byte) (*balances.EmmitBalanceRequest, error) {
		var request balances.EmmitBalanceRequest
		err := proto.Unmarshal(body, &request)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		fmt.Println(request.String())
		return &request, nil
	}
}

func GetParserWalletInfoRequest() func([]byte) (*balances.GetWalletInfoRequest, error) {
	return func(body []byte) (*balances.GetWalletInfoRequest, error) {
		var request balances.GetWalletInfoRequest
		err := proto.Unmarshal(body, &request)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		fmt.Println(request.String())
		return &request, nil
	}
}

func GetParserLockBalanceRequest() func([]byte) (*balances.LockBalanceRequest, error) {
	return func(b []byte) (*balances.LockBalanceRequest, error) {
		var request balances.LockBalanceRequest
		err := proto.Unmarshal(b, &request)
		if err != nil {
			return nil, err
		}
		return &request, nil
	}
}

func GetParserTransferRequest() func([]byte) (*balances.CreateTransferRequest, error) {
	return func(b []byte) (*balances.CreateTransferRequest, error) {
		var request balances.CreateTransferRequest
		if err := proto.Unmarshal(b, &request); err != nil {
			return nil, err
		}
		return &request, nil
	}
}

func NewAmqpFactory(connectionString string, loggger *logrus.Logger) AmqpFactory {
	var con *amqp091.Connection
	var err error
	for {
		con, err = amqp091.Dial(connectionString)
		if err != nil {
			loggger.Errorln("Rabbit unvailable. Reazon: ", err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	return AmqpFactory{logger: loggger, con: con}
}

func (f *AmqpFactory) RunReconnect(ctx context.Context, connectionString string) {
	for {
		select {
		case connErr := <-f.con.NotifyClose(make(chan *amqp091.Error)):
			{
				f.logger.Errorln("Connection was failed: message: ", connErr.Error())
				for {
					con, err := amqp091.Dial(connectionString)
					if err != nil {
						f.logger.Errorln("Reconnect failed! message: ", err.Error())
						time.Sleep(5 * time.Second)
						continue
					}
					f.con = con
					break
				}
			}
		case <-ctx.Done():
			break
		default:
			time.Sleep(10 * time.Second)
		}
	}
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
	return AmqpSender{channel: ch, logger: f.logger, exchange: exchange, rk: rk}
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
	ch.ExchangeDeclare(statics.ExNameBalances, "topic", true, false, false, false, nil)
	balanceGetWalletInfoRequestQ, _ := ch.QueueDeclare(statics.GetWalletInfoRequestQueue, true, false, false, false, nil)
	balanceGetWalletInfoResponse, _ := ch.QueueDeclare(statics.GetWalletInfoResponseQueue, true, false, false, false, nil)
	balanceEmmitRequestQ, _ := ch.QueueDeclare(statics.EmmitBalanceRequestQueue, true, false, false, false, nil)
	lockBalanceRequestQ, _ := ch.QueueDeclare(statics.LockBalanceRequestQueue, true, false, false, false, nil)
	lockBalanceResponseQ, _ := ch.QueueDeclare(statics.LockBalanceResponseQueue, true, false, false, false, nil)
	createTransferRequestQ, _ := ch.QueueDeclare(statics.TransferBalanceRequestQueue, true, false, false, false, nil)
	createTransferResponseQ, _ := ch.QueueDeclare(statics.TransferBalanceResponseQueue, true, false, false, false, nil)
	ch.QueueBind(lockBalanceResponseQ.Name, statics.RkLockBalanceResponse, statics.ExNameBalances, false, nil)
	ch.QueueBind(lockBalanceRequestQ.Name, statics.RkLockBalanceRequest, statics.ExNameBalances, false, nil)
	ch.QueueBind(balanceEmmitRequestQ.Name, statics.RkEmmitBalance, statics.ExNameBalances, false, nil)
	ch.QueueBind(balanceGetWalletInfoRequestQ.Name, statics.RkGetWalletInfoRequest, statics.ExNameBalances, false, nil)
	ch.QueueBind(balanceGetWalletInfoResponse.Name, statics.RkGetWalletInfoResponse, statics.ExNameBalances, false, nil)
	ch.QueueBind(createTransferRequestQ.Name, statics.RkTransferBalanceRequest, statics.ExNameBalances, false, nil)
	ch.QueueBind(createTransferResponseQ.Name, statics.RkTransferBalanceResponse, statics.ExNameBalances, false, nil)
}
