package services

import (
	"balance-service/sql"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/nejkit/processing-proto/balances"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

func InitConnection(connectionString string) *pgx.Conn {
	return sql.GetConnection(connectionString)
}

func EmmitBalance(request *balances.EmmitBalanceRequest, connection *pgx.Conn) {

	if request.GetAmount() <= 0 {

	}

	if !sql.CheckExistsWallet(request.GetAddress(), connection) {
		sql.InsertWalletData(request.GetAddress(), connection)
	}

	balanceExists, id := sql.CheckExistsBalance(request.GetAddress(), request.GetCurrency(), connection)
	if balanceExists {
		sql.EmmitBalanceData(id, request.GetAmount(), connection)
	} else {
		sql.InsertBalanceData(request.GetAddress(), request.GetCurrency(), request.GetAmount(), connection)
	}
}

func GetInfoAboutBalance(inMessages <-chan amqp091.Delivery, outMessages chan<- *balances.GetWalletInfoResponse, connection *pgx.Conn) {
	for msg := range inMessages {
		var request balances.GetWalletInfoRequest
		err := proto.Unmarshal(msg.Body, &request)
		if err != nil {

		}
		walletInfo := sql.GetWalletInfo(request.GetAddress(), connection)
		balancesInfo := sql.GetBalancesByWallet(request.GetAddress(), connection)

		response := Mapper(walletInfo, balancesInfo)

		fmt.Println(response.String())
		outMessages <- &balances.GetWalletInfoResponse{Id: request.GetId(), WalletInfo: &response}
	}

}
