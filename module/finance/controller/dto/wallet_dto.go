package dto

import "boilerplate/module/finance/repository/model"

type WalletTopUpRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	UserID      string  `json:"user_id"`
}

type WalletUseUpRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

type WalletBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type WalletHistoryResponse struct {
	Items []model.WalletHistoryItem `json:"items"`
	Page  int                       `json:"page"`
	Limit int                       `json:"limit"`
	Total int                       `json:"total"`
}
