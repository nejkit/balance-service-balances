package sql

import (
	"context"
	"fmt"
	"sync"
	"time"

	pgx "github.com/jackc/pgx/v5"
	idGen "github.com/matoous/go-nanoid"
	proto "github.com/nejkit/processing-proto/balances"
)

type postgresConnection struct {
	Connection *pgx.Conn
}

var connection *postgresConnection

func getConnection() postgresConnection {
	var lock = sync.Mutex{}
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		con, err := pgx.Connect(context.Background(), "postgres://postgre:admin@posgres:5432/servicebalances")
		if err != nil {
			panic(err.Error())
		}
		connection = &postgresConnection{Connection: con}
		return *connection
	}
	return *connection
}

func EmmitBalance(request *proto.EmmitBalanceRequest) *proto.EmmitBalanceResponse {

	if request.GetAmount() <= 0 {
		return &proto.EmmitBalanceResponse{
			Id:        request.Id,
			ErrorCode: "InvalidAmountEmmit",
		}
	}
	con := getConnection().Connection
	var count int
	err := con.QueryRow(context.Background(), "select count(*) from wallets where id = $1", request.GetAddress()).Scan(&count)
	if err != nil {
		return &proto.EmmitBalanceResponse{
			Id:        request.Id,
			ErrorCode: "InternalError",
		}
	}

	if count == 0 {
		fmt.Printf("Data is new. Try insert to db Address: %s", request.GetAddress())
		fmt.Println()
		_, err = con.Exec(context.Background(), "insert into wallets(id, created) values($1, $2)", request.GetAddress(), time.Now())
		if err != nil {
			return &proto.EmmitBalanceResponse{
				Id:        request.Id,
				ErrorCode: "InternalError",
			}
		}
	}
	err = con.QueryRow(context.Background(), "select count(*) from balances where walletaddress=$1 and currency=$2",
		request.GetAddress(), request.GetCurrency()).Scan(&count)

	if err != nil {
		return &proto.EmmitBalanceResponse{
			Id:        request.Id,
			ErrorCode: "InternalError",
		}
	}
	idBalance, _ := idGen.ID(10)
	if count == 0 {
		fmt.Printf("Balance with currency %s not exists. Try insert to db balance with Address: %s", request.GetCurrency(), request.GetAddress())
		_, err := con.Exec(context.Background(), "insert into balances(id, walletaddress, currency, actualbalance) values($1, $2, $3, $4)",
			idBalance, request.GetAddress(), request.GetCurrency(), request.GetAmount())
		if err != nil {
			return &proto.EmmitBalanceResponse{
				Id:        request.Id,
				ErrorCode: "InternalError",
			}
		}

	} else {
		fmt.Printf("Balance with currency %s exists. Try update balance with Address: %s", request.GetCurrency(), request.GetAddress())
		_, err := con.Exec(context.Background(), "update balances set actualbalance = actualbalance + $1 where walletaddress = $2 and currency = $3",
			request.GetAmount(), request.GetAddress(), request.GetCurrency())

		if err != nil {
			return &proto.EmmitBalanceResponse{
				Id:        request.Id,
				ErrorCode: "InternalError",
			}
		}
	}

	return &proto.EmmitBalanceResponse{
		Id:    request.Id,
		State: proto.EmmitBalanceState_DONE,
	}
}

func GetInfoAboutBalance(address string) BalanceModel {

}
