package sql

type WalletModel struct {
	Id        string
	Created   uint64
	IsDeleted bool
}

type BalanceModel struct {
	Id            string
	WalletAddress string
	Currency      string
	ActualBalance float64
	FreezeBalance float64
}
