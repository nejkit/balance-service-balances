package sql

const (
	CheckExistsWalletQuery  = "select case when exists (select * from wallets where id = $1) then 1 else 0"
	InsertWalletQuery       = "insert into wallets(id, created) values($1, $2)"
	CheckExistsBalanceQuery = "select id from balances where waleltaddress = $1 and currency = $2"
	InsertBalanceQuery      = "insert into balances(id, walletaddress, currency, actualbalance) values($1, $2, $3, $4)"
	EmmitBalanceQuery       = "update balances set actualbalance = actualbalance + $1 where id = $2"
	GetWalletQuery          = "select id, created, isdeleted from wallets where id = $1"
	GetBalancesQuery        = "select id, currency, actualbalance, freezebalance from balances where walletaddress = $1"
)
