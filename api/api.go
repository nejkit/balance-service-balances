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
	for msg := range listener {
		var request balances.EmmitBalanceRequest
		proto.Unmarshal(msg.Body, &request)
		go services.EmmitBalance(&request, pgxCon)
	}
}

func InnitGetWalletInfoApi(rmqChannel *amqp091.Channel, pgxCon *pgx.Conn, logger *logrus.Logger) {
	listener := rmq.InitListener(statics.GetWalletInfoRequestQueue, rmqChannel)
	senderChan := make(chan *balances.GetWalletInfoResponse)
	go services.GetInfoAboutBalance(listener, senderChan, pgxCon)
	go rmq.SendWalletInfo(senderChan, rmqChannel)
}
