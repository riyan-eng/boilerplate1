package entity

import (
	"boilerplate/module/finance/repository/model"

	"github.com/valyala/fasthttp"
)

type WalletTopUpRequest struct {
	Context     *fasthttp.RequestCtx
	Description string
	Amount      float64
	UserID      string
	CompanyID   string
	CreatedBy   string
}

type WalletTopUpResponse struct {
	Error error
}

type WalletUseUpRequest struct {
	Context     *fasthttp.RequestCtx
	Description string
	Amount      float64
	UserID      string
	CompanyID   string
	CreatedBy   string
}

type WalletUseUpResponse struct {
	Error error
}

type WalletBalanceRequest struct {
	Context *fasthttp.RequestCtx
	UserID  string
}

type WalletBalanceResponse struct {
	Balance float64
	Error   error
}

type WalletHistoryRequest struct {
	Context *fasthttp.RequestCtx
	UserID  string
	Page    int
	Limit   int
}

type WalletHistoryResponse struct {
	Items []model.WalletHistoryItem
	Total int
	Page  int
	Limit int
	Error error
}
