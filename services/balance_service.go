package services

import (
	"balance-service/sql"

	pgx "github.com/jackc/pgx/v5"
	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func InitConnection(connectionString string) *pgx.Conn {
	return sql.GetConnection(connectionString)
}

func EmmitBalance(request *balances.EmmitBalanceRequest, connection *pgx.Conn, logger *logrus.Logger) {

	if request.GetAmount() <= 0 {

	}

	if !sql.CheckExistsWallet(request.GetAddress(), connection) {
		sql.InsertWalletData(request.GetAddress(), connection)
	}

	balanceExists, id := sql.CheckExistsBalance(request.GetAddress(), request.GetCurrency(), connection, logger)
	if balanceExists {
		sql.EmmitBalanceData(id, request.GetAmount(), connection)
	} else {
		sql.InsertBalanceData(request.GetAddress(), request.GetCurrency(), request.GetAmount(), connection, logger)
	}
}

func GetInfoAboutBalance(inMessages <-chan amqp091.Delivery, outMessages chan<- *balances.GetWalletInfoResponse, connection *pgx.Conn, logger *logrus.Logger) {
	forever := make(chan bool)
	for msg := range inMessages {
		var request balances.GetWalletInfoRequest
		err := proto.Unmarshal(msg.Body, &request)
		logger.Info("Address:", request.GetAddress())
		if err != nil {
			msg.Nack(false, false)
		} else {
			msg.Ack(false)
		}
		walletInfo := sql.GetWalletInfo(request.GetAddress(), connection, logger)
		balancesInfo := sql.GetBalancesByWallet(request.GetAddress(), connection, logger)

		response := Mapper(walletInfo, balancesInfo)

		outMessages <- &balances.GetWalletInfoResponse{Id: request.GetId(), WalletInfo: &response}
		logger.Info("Put to sender chan: ", response.String())
	}
	<-forever
}
