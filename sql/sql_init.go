package sql

import (
	"context"
	"time"

	pgx "github.com/jackc/pgx/v5"
	proto "github.com/nejkit/processing-proto/balances"
)

func InitDb() *pgx.Conn {
	con, err := pgx.Connect(context.Background(), "postgres://postgres:admin@127.0.1.1:5432/servicebalance")
	if err != nil {
		panic(err.Error())
	}

	return con
}

func EmmitBalance(request *proto.EmmitBalanceRequest, con *pgx.Conn) error {

	var count int
	err := con.QueryRow(context.Background(), "select count(*) from wallets where id = ?", request.GetAddress()).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		con.Exec(context.Background(), "insert into wallets(id, created) values($1, $2)", request.GetAddress(), time.Now().UnixMilli())
	}
	err = con.QueryRow(context.Background(), "select count(*) from balances where id=$1 and currency=$2",
		request.GetAddress(), request.GetCurrency()).Scan(&count)

	if err != nil {
		return err
	}

	if count == 0 {
		con.Exec(context.Background(), "insert into balances(walletaddress, currency, actualbalance) values($1, $2, $3)",
			request.GetAddress(), request.GetCurrency(), request.GetAmount())
	} else {
		con.Exec(context.Background(), "update balances set amount = amount + $1 where walletaddress = $2 and currency = $3",
			request.GetAmount(), request.GetAddress(), request.GetCurrency())
	}

	return nil
}
