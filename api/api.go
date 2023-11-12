package api

import (
	"balance-service/rmq"
	"balance-service/services"
	"balance-service/statics"

	"github.com/jackc/pgx/v5"
	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func InnitEmmitBalanceApi(rmqChannel *amqp091.Channel, pgxCon *pgx.Conn, logger *logrus.Logger) {
	listener := rmq.InitListener(statics.EmmitBalanceRequestQueue, rmqChannel)
	forever := make(chan bool)
	for msg := range listener {
		var request balances.EmmitBalanceRequest
		err := proto.Unmarshal(msg.Body, &request)
		if err != nil {
			msg.Nack(false, false)
		} else {
			msg.Ack(false)
		}
		go services.EmmitBalance(&request, pgxCon, logger)
	}
	<-forever
}

func InnitGetWalletInfoApi(rmqChannel *amqp091.Channel, pgxCon *pgx.Conn, logger *logrus.Logger) {
	listener := rmq.InitListener(statics.GetWalletInfoRequestQueue, rmqChannel)
	senderChan := make(chan *balances.GetWalletInfoResponse, 1)
	go services.GetInfoAboutBalance(listener, senderChan, pgxCon, logger)
	go rmq.SendWalletInfo(senderChan, rmqChannel, logger)
}
