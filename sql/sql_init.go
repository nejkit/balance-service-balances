package sql

import (
	"context"
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v5"
	idGen "github.com/matoous/go-nanoid"
	proto "github.com/nejkit/processing-proto/balances"
)

func InitDb() *pgx.Conn {
	con, err := pgx.Connect(context.Background(), "postgres://postgre:admin@postgres:5432/servicebalances")
	if err != nil {
		panic(err.Error())
	}

	return con
}

func EmmitBalance(request *proto.EmmitBalanceRequest, con *pgx.Conn) error {

	var count int
	err := con.QueryRow(context.Background(), "select count(*) from wallets where id = $1", request.GetAddress()).Scan(&count)
	if err != nil {
		return err
	}

	fmt.Printf("Data is exists: %d", count)

	if count == 0 {
		_, err = con.Exec(context.Background(), "insert into wallets(id, created) values($1, $2)", request.GetAddress(), time.Now().UnixMilli())
		if err != nil {
			return err
		}
	}
	err = con.QueryRow(context.Background(), "select count(*) from balances where id=$1 and currency=$2",
		request.GetAddress(), request.GetCurrency()).Scan(&count)

	fmt.Printf("Balance is exists: %d", count)

	if err != nil {
		return err
	}
	idBalance, _ := idGen.ID(10)
	if count == 0 {
		_, err := con.Exec(context.Background(), "insert into balances(id, walletaddress, currency, actualbalance) values($1, $2, $3, $4)",
			idBalance, request.GetAddress(), request.GetCurrency(), request.GetAmount())
		if err != nil {
			return err
		}

	} else {
		_, err := con.Exec(context.Background(), "update balances set amount = amount + $1 where walletaddress = $2 and currency = $3",
			request.GetAmount(), request.GetAddress(), request.GetCurrency())

		if err != nil {
			return err
		}
	}

	return nil
}
