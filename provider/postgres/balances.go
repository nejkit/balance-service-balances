package postgres

import (
	"balance-service/external/balances"
	"balance-service/sql"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BalanceAdapter struct {
	conn *pgxpool.Pool
}

func NewBalanceAdapter(connPool *pgxpool.Pool) BalanceAdapter {
	return BalanceAdapter{conn: connPool}

}

func (a *BalanceAdapter) EmmitBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest, id string) {
	a.conn.Exec(ctx, sql.EmmitBalanceQuery, info.GetAmount(), id)
}

func (a *BalanceAdapter) InsertBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest) {
	id := uuid.NewString()
	a.conn.Exec(ctx, sql.InsertBalanceQuery, id, info.GetAddress(), info.GetCurrency(), info.GetAmount())
}

func (a *BalanceAdapter) GetBalancesInfo(ctx context.Context, address string) []*sql.BalanceModel {
	var balancesInfo []*sql.BalanceModel
	balanceRows, err := a.conn.Query(ctx, sql.GetBalancesQuery, address)
	if err == pgx.ErrNoRows {
		return nil
	}
	for balanceRows.Next() {
		var balanceInfo sql.BalanceModel
		balanceRows.Scan(&balanceInfo.Id, &balanceInfo.Currency, &balanceInfo.ActualBalance, &balanceInfo.FreezeBalance)
		balancesInfo = append(balancesInfo, &balanceInfo)
	}
	return balancesInfo
}

func (a *BalanceAdapter) GetBalanceInfo(ctx context.Context, address string, currency string) *sql.BalanceModel {
	var balanceInfo *sql.BalanceModel
	err := a.conn.QueryRow(ctx, sql.GetBalanceQuery, address, currency).Scan(&balanceInfo.Id, &balanceInfo.Currency, &balanceInfo.ActualBalance, &balanceInfo.FreezeBalance)
	if err == pgx.ErrNoRows {
		return nil
	}
	return balanceInfo

}

func (a *BalanceAdapter) LockTransferBalance(ctx context.Context, id string, amount float64) {
	a.conn.Exec(ctx, sql.LockBalanceQuery, id, amount)
}

func (a *BalanceAdapter) UnlockTransferBalance(ctx context.Context, address string, amount float64, cur string) {
	a.conn.Exec(ctx, sql.UnLockBalanceQuery, address, amount, cur)
}
