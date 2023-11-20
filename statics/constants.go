package statics

const (
	ExNameBalances             = "e.balances.forward"
	EmmitBalanceRequestQueue   = "q.balances.request.EmmitBalanceRequest"
	GetWalletInfoRequestQueue  = "q.balances.request.GetWalletInfoRequest"
	GetWalletInfoResponseQueue = "q.balances.response.GetWalletInfoResponse"
	LockBalanceRequestQueue    = "q.balances.request.LockBalanceRequest"
	LockBalanceResponseQueue   = "q.balances.response.LockBalanceResponse"
	RkEmmitBalance             = "r.balances.#.EmmitBalanceRequest.#"
	RkGetWalletInfoRequest     = "r.balances.#.GetWalletInfoRequest.#"
	RkGetWalletInfoResponse    = "r.balances.#.GetWalletInfoResponse.#"
	RkLockBalanceRequest       = "r.#.request.LockBalanceRequest.#"
	RkLockBalanceResponse      = "r.#.response.LockBalanceResponse.#"
)
