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
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.EmmitBalanceQuery, info.GetAmount(), id)
}

func (a *BalanceAdapter) InsertBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest) {
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	id := uuid.NewString()
	con.Exec(ctx, sql.InsertBalanceQuery, id, info.GetAddress(), info.GetCurrency(), info.GetAmount())
}

func (a *BalanceAdapter) GetBalancesInfo(ctx context.Context, address string) []*sql.BalanceModel {
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	var balancesInfo []*sql.BalanceModel
	balanceRows, err := con.Query(ctx, sql.GetBalancesQuery, address)
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
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	var balanceInfo sql.BalanceModel
	err := con.QueryRow(ctx, sql.GetBalanceQuery, address, currency).Scan(&balanceInfo.Id, &balanceInfo.Currency, &balanceInfo.ActualBalance, &balanceInfo.FreezeBalance)
	if err == pgx.ErrNoRows {
		return nil
	}
	return &balanceInfo

}

func (a *BalanceAdapter) LockTransferBalance(ctx context.Context, id string, amount float64) {
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.LockBalanceQuery, id, amount)
}

func (a *BalanceAdapter) UnlockTransferBalance(ctx context.Context, address string, amount float64, cur string) {
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.UnLockBalanceQuery, address, amount, cur)
}
