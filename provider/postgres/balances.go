package postgres

import (
	"balance-service/external/balances"
	"balance-service/sql"
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BalanceAdapter struct {
	conn *pgxpool.Pool
	mtx  sync.Mutex
}

func NewBalanceAdapter(connPool *pgxpool.Pool) BalanceAdapter {
	return BalanceAdapter{conn: connPool, mtx: sync.Mutex{}}

}

func (a *BalanceAdapter) EmmitBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest, id string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.EmmitBalanceQuery, info.GetAmount(), id)
}

func (a *BalanceAdapter) InsertBalanceInfo(ctx context.Context, info *balances.EmmitBalanceRequest) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
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
	a.mtx.Lock()
	defer a.mtx.Unlock()
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.LockBalanceQuery, id, amount)
}

func (a *BalanceAdapter) TransferMoney(
	ctx context.Context,
	senderOptions *balances.TransferOptions,
	receiptOptions *balances.TransferOptions,
	senderId string,
	receiptId string) error {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()

	tx, err := con.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadUncommitted})
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, sql.ChargeFreezeBalanceQuery, senderOptions.GetAmount(), senderId)
	_, err = tx.Exec(ctx, sql.ChargeFreezeBalanceQuery, receiptOptions.GetAmount(), receiptId)
	_, err = tx.Exec(ctx, sql.EmmitBalanceQuery, senderOptions.GetAmount(), receiptId)
	_, err = tx.Exec(ctx, sql.EmmitBalanceQuery, receiptOptions.GetAmount(), senderId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	tx.Commit(ctx)
	return nil
}

func (a *BalanceAdapter) UnlockTransferBalance(ctx context.Context, address string, amount float64, cur string) {
	a.mtx.Lock()
	defer a.mtx.Unlock()
	con, _ := a.conn.Acquire(ctx)
	defer con.Release()
	con.Exec(ctx, sql.UnLockBalanceQuery, address, amount, cur)
}
