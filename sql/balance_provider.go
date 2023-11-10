package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func GetConnection(connectionString string) *pgx.Conn {
	con, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {

	}
	return con
}

func CheckExistsWallet(address string, connection *pgx.Conn) bool {

	exists := false
	connection.QueryRow(context.Background(), CheckExistsWalletQuery, address).Scan(exists)
	return exists
}

func CheckExistsBalance(address string, currency string, connection *pgx.Conn) (bool, string) {
	var id *string = nil
	var exists bool = false
	connection.QueryRow(context.Background(), CheckExistsBalanceQuery, address, currency).Scan(id)
	if id != nil {
		exists = true
	}
	return exists, *id
}

func InsertWalletData(address string, connection *pgx.Conn) {
	connection.Exec(context.Background(), InsertWalletQuery, address, time.Now())
}

func InsertBalanceData(address string, currency string, amount float64, connection *pgx.Conn) string {
	id, _ := uuid.NewRandom()
	connection.Exec(context.Background(), InsertBalanceQuery, id.String(), address, currency, amount)
	return id.String()
}

func EmmitBalanceData(id string, amount float64, con *pgx.Conn) {
	con.Exec(context.Background(), EmmitBalanceQuery, amount, id)
}

func GetWalletInfo(address string, con *pgx.Conn) WalletModel {
	var model WalletModel
	con.QueryRow(context.Background(), GetWalletQuery, address).Scan(&model.Id, &model.Created, &model.IsDeleted)
	return model
}

func GetBalancesByWallet(address string, con *pgx.Conn) []*BalanceModel {
	var balanceModels []*BalanceModel
	rows, _ := con.Query(context.Background(), GetBalancesQuery, address)
	for rows.Next() {
		var balanceModel BalanceModel
		rows.Scan(&balanceModel.Id, &balanceModel.Currency, &balanceModel.ActualBalance, &balanceModel.FreezeBalance)
		balanceModels = append(balanceModels, &balanceModel)
	}

	return balanceModels
}
