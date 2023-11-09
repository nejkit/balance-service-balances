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
		con, err := pgx.Connect(context.Background(), "postgres://postgre:admin@postgres:5432/servicebalances")
		if err != nil {
			panic(err.Error())
		}
		connection = &postgresConnection{Connection: con}
		return *connection
	}
	return *connection
}

func EmmitBalance(request *proto.EmmitBalanceRequest) error {

	if request.GetAmount() <= 0 {
		return nil
	}
	con := getConnection().Connection
	var count int
	err := con.QueryRow(context.Background(), "select count(*) from wallets where id = $1", request.GetAddress()).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		fmt.Printf("Data is new. Try insert to db Address: %s", request.GetAddress())
		fmt.Println()
		_, err = con.Exec(context.Background(), "insert into wallets(id, created) values($1, $2)", request.GetAddress(), time.Now())
		if err != nil {
			return err
		}
	}
	err = con.QueryRow(context.Background(), "select count(*) from balances where walletaddress=$1 and currency=$2",
		request.GetAddress(), request.GetCurrency()).Scan(&count)

	if err != nil {
		return err
	}
	idBalance, _ := idGen.ID(10)
	if count == 0 {
		fmt.Printf("Balance with currency %s not exists. Try insert to db balance with Address: %s", request.GetCurrency(), request.GetAddress())
		_, err := con.Exec(context.Background(), "insert into balances(id, walletaddress, currency, actualbalance) values($1, $2, $3, $4)",
			idBalance, request.GetAddress(), request.GetCurrency(), request.GetAmount())
		if err != nil {
			return err
		}

	} else {
		fmt.Printf("Balance with currency %s exists. Try update balance with Address: %s", request.GetCurrency(), request.GetAddress())
		_, err := con.Exec(context.Background(), "update balances set actualbalance = actualbalance + $1 where walletaddress = $2 and currency = $3",
			request.GetAmount(), request.GetAddress(), request.GetCurrency())

		if err != nil {
			return err
		}
	}

	return nil
}

func GetInfoAboutBalance(address string) *proto.WalletInfo {
	con := getConnection().Connection

	walletInfo := &WalletModel{}
	con.QueryRow(context.Background(), "select id, created, isdeleted from wallets where id = $1", address).Scan(&walletInfo.Id, &walletInfo.Created, &walletInfo.IsDeleted)
	rows, _ := con.Query(context.Background(), "select id, currency, actualbalance, freezebalance from balances where walletaddress = $1", address)
	balancesList := []*proto.BalanceInfo{}
	for rows.Next() {
		balanceModel := &BalanceModel{}
		if err := rows.Scan(&balanceModel.Id, &balanceModel.Currency, &balanceModel.ActualBalance, &balanceModel.FreezeBalance); err != nil {
			panic(err.Error())
		}
		balanceProto := proto.BalanceInfo{
			Id:            balanceModel.Id,
			Currency:      balanceModel.Currency,
			ActualBalance: float32(balanceModel.ActualBalance),
			FreezeBalance: float32(balanceModel.FreezeBalance),
		}
		balancesList = append(balancesList, &balanceProto)
	}
	response := proto.WalletInfo{
		Address:      walletInfo.Id,
		Created:      uint64(walletInfo.Created.UnixMilli()),
		IsDeleted:    walletInfo.IsDeleted,
		BalanceInfos: balancesList,
	}
	fmt.Println(response.String())
	return &response

}
