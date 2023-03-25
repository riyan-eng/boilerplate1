package model

import "github.com/valyala/fasthttp"

type WalletTopUpRequest struct {
	Context         *fasthttp.RequestCtx
	Description     string
	TransactionCode string
	CoaDebit        string
	CoaCredit       string
	Amount          float64
	UserID          string
	CompanyID       string
	CreatedBy       string
}

type WalletTopUpResponse struct {
	Error error
}

type WalletUseUpRequest struct {
	Context         *fasthttp.RequestCtx
	Description     string
	TransactionCode string
	CoaDebit        string
	CoaCredit       string
	Amount          float64
	UserID          string
	CompanyID       string
	CreatedBy       string
}

type WalletUseUpResponse struct {
	Error error
}

type WalletBalanceRequest struct {
	Context *fasthttp.RequestCtx
	CoaKas  string
	UserID  string
}

type WalletBalanceResponse struct {
	Balance float64
	Error   error
}

type WalletHistoryRequest struct {
	Context *fasthttp.RequestCtx
	CoaKas  string
	UserID  string
	Limit   int
	Offset  int
}

type WalletHistoryResponse struct {
	Items []WalletHistoryItem
	Error error
}

type WalletHistoryItem struct {
	TransactionCode string  `db:"transaction_code" json:"transaction_code"`
	Decription      string  `db:"description" json:"description"`
	Amount          float64 `db:"amount" json:"amount"`
	CreatedAt       string  `db:"created_at" json:"created_at"`
	Total           int     `db:"total"`
}
