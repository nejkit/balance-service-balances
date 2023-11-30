package statics

const (
	ExNameBalances               = "e.balances.forward"
	EmmitBalanceRequestQueue     = "q.balances.request.EmmitBalanceRequest"
	GetWalletInfoRequestQueue    = "q.balances.request.GetWalletInfoRequest"
	GetWalletInfoResponseQueue   = "q.balances.response.GetWalletInfoResponse"
	LockBalanceRequestQueue      = "q.balances.request.LockBalanceRequest"
	LockBalanceResponseQueue     = "q.balances.response.LockBalanceResponse"
	UnLockBalanceRequestQueue    = "q.balances.request.UnLockBalanceRequest"
	UnLockBalanceResponseQueue   = "q.balances.response.UnLockBalanceResponse"
	TransferBalanceRequestQueue  = "q.balances.request.TransferBalanceRequest"
	TransferBalanceResponseQueue = "q.balances.response.TransferBalanceResponse"
	RkTransferBalanceRequest     = "r.balances.#.TransferBalanceRequest.#"
	RkTransferBalanceResponse    = "r.balances.#.TransferBalanceResponse.#"
	RkEmmitBalance               = "r.balances.#.EmmitBalanceRequest.#"
	RkGetWalletInfoRequest       = "r.balances.#.GetWalletInfoRequest.#"
	RkGetWalletInfoResponse      = "r.balances.#.GetWalletInfoResponse.#"
	RkLockBalanceRequest         = "r.#.request.LockBalanceRequest.#"
	RkLockBalanceResponse        = "r.#.response.LockBalanceResponse.#"
	RkUnLockBalanceRequest       = "r.#.request.UnLockBalanceRequest.#"
	RkunLockBalanceResponse      = "r.#.response.UnLockBalanceResponse.#"
)
