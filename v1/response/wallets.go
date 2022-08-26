package response

type AllWalletsResponse struct {
	StatusCode int
	Data       []Wallet `json:"data"`
}

type WalletResponse struct {
	StatusCode int
	Data       Wallet `json:"data"`
}

type Wallet struct {
	WalletId    string    `json:"walletId"`
	EntityId    string    `json:"entityId"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Balances    []Balance `json:"balances"`
}

type Balance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type ErrorResponse struct {
	StatusCode int
	Code       int    `json:"code"`
	Message    string `json:"message"`
}
