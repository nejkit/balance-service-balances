package postgres

import (
	"balance-service/external/balances"
	"balance-service/sql"
	"context"
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
	a.conn.Exec(ctx, sql.InsertWalletQuery, info.GetAddress(), time.Now())
}

func (a *WalletAdapter) GetWalletInfo(ctx context.Context, address string) *sql.WalletModel {
	var walletInfo sql.WalletModel
	err := a.conn.QueryRow(ctx, sql.GetWalletQuery, address).Scan(&walletInfo.Id, &walletInfo.Created, &walletInfo.IsDeleted)
	if err == pgx.ErrNoRows {
		return nil
	}
	return &walletInfo
}
