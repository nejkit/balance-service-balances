package sql

import "time"

type WalletModel struct {
	Id        string
	Created   time.Time
	IsDeleted bool
}

type BalanceModel struct {
	Id            string
	WalletAddress string
	Currency      string
	ActualBalance float64
	FreezeBalance float64
}
