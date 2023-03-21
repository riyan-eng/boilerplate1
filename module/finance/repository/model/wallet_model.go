package model

type WalletTopUpRequest struct {
	Description     string
	TransactionCode string
	CoaDebit        string
	CoaCredit       string
	Amount          float64
}

type WalletTopUpResponse struct {
	Error error
}

type WalletUseUpRequest struct {
	Description     string
	TransactionCode string
	CoaDebit        string
	CoaCredit       string
	Amount          float64
}

type WalletUseUpResponse struct {
	Error error
}
