package sql

const (
	CheckExistsWalletQuery   = "select case when exists (select * from wallets where id = $1) then 1 else 0 end"
	InsertWalletQuery        = "insert into wallets(id, created) values($1, $2)"
	CheckExistsBalanceQuery  = "select id from balances where walletaddress = $1 and currency = $2"
	InsertBalanceQuery       = "insert into balances(id, walletaddress, currency, actualbalance) values($1, $2, $3, $4)"
	EmmitBalanceQuery        = "update balances set actualbalance = actualbalance + $1 where id = $2"
	GetWalletQuery           = "select id, created, isdeleted from wallets where id = $1"
	GetBalancesQuery         = "select id, currency, actualbalance, freezebalance from balances where walletaddress = $1"
	GetBalanceQuery          = "select id, currency, actualbalance, freezebalance from balances where walletaddress = $1 and currency = $2"
	LockBalanceQuery         = "update balances set actualbalance = actualbalance - $2, freezebalance = freezebalance + $2 where id = $1"
	UnLockBalanceQuery       = "update balances set actualbalance = actualbalance + $3, freezebalance = freezebalance - $3 where walletaddress = $1 and currency = $2"
	ChargeFreezeBalanceQuery = "update balances set freezebalance = freezebalance - $1 where id = $2"
)
