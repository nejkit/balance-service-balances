package postgres

import (
	"balance-service/external/balances"
	"balance-service/sql"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletAdapter struct {
	conn *pgxpool.Pool
}

func NewWalletAdapter(connection *pgxpool.Pool) WalletAdapter {
	return WalletAdapter{conn: connection}
}

func (a *WalletAdapter) InsertWalletInfo(ctx context.Context, info *balances.EmmitBalanceRequest) {
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.InsertWalletQuery, info.GetAddress(), time.Now())
}

func (a *WalletAdapter) GetWalletInfo(ctx context.Context, address string) *sql.WalletModel {
	var walletInfo sql.WalletModel
	con, err := a.conn.Acquire(ctx)
	if err != nil {
		panic(err.Error())
	}
	defer con.Release()
	err = con.QueryRow(ctx, sql.GetWalletQuery, address).Scan(&walletInfo.Id, &walletInfo.Created, &walletInfo.IsDeleted)
	if err == pgx.ErrNoRows {
		return nil
	}
	fmt.Println("Address from db: ", walletInfo.Id)
	return &walletInfo
}
