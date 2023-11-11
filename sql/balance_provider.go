package sql

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
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

func CheckExistsBalance(address string, currency string, connection *pgx.Conn, logger *logrus.Logger) (bool, string) {
	var id string
	err := connection.QueryRow(context.Background(), CheckExistsBalanceQuery, address, currency).Scan(&id)
	if err == pgx.ErrNoRows {
		return false, ""
	}
	if err != nil {
		logger.Fatal(err.Error())
		return false, ""
	} else {
		return true, id
	}
}

func InsertWalletData(address string, connection *pgx.Conn) {
	connection.Exec(context.Background(), InsertWalletQuery, address, time.Now())
}

func InsertBalanceData(address string, currency string, amount float64, connection *pgx.Conn, logger *logrus.Logger) string {
	id, _ := uuid.NewRandom()
	_, err := connection.Exec(context.Background(), InsertBalanceQuery, id.String(), address, currency, amount)
	if err != nil {
		logger.Info(err.Error())
	}
	return id.String()
}

func EmmitBalanceData(id string, amount float64, con *pgx.Conn) {
	con.Exec(context.Background(), EmmitBalanceQuery, amount, id)
}

func GetWalletInfo(address string, con *pgx.Conn, logger *logrus.Logger) *WalletModel {
	var model WalletModel
	err := con.QueryRow(context.Background(), GetWalletQuery, address).Scan(&model.Id, &model.Created, &model.IsDeleted)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("WalletModel: ", "Id: ", model.Id, " Created: ", model.Created, " Deleted: ", model.IsDeleted)
	return &model
}

func GetBalancesByWallet(address string, con *pgx.Conn, logger *logrus.Logger) []*BalanceModel {
	var balanceModels []*BalanceModel
	rows, err := con.Query(context.Background(), GetBalancesQuery, address)
	if err != nil {
		logger.Fatal(err.Error())
	}
	for rows.Next() {
		var balanceModel BalanceModel
		rows.Scan(&balanceModel.Id, &balanceModel.Currency, &balanceModel.ActualBalance, &balanceModel.FreezeBalance)
		balanceModels = append(balanceModels, &balanceModel)
		logger.Info("BalanceModel: ", "Id: ", balanceModel.Id, " Cur: ", balanceModel.Currency, " Balance: ", balanceModel.ActualBalance)
	}

	return balanceModels
}
