package postgres

import (
	"balance-service/sql"
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/nejkit/processing-proto/balances"
)

type WalletAdapter struct {
	conn *pgx.Conn
}

func NewWalletAdapter(connection *pgx.Conn) WalletAdapter {
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
