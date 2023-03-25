package service

import (
	"boilerplate/module/finance/repository"
	"boilerplate/module/finance/repository/model"
	"boilerplate/module/finance/service/entity"
	"boilerplate/util"
	"errors"
)

type WalletService interface {
	TopUp(*entity.WalletTopUpRequest) entity.WalletTopUpResponse
	UseUp(*entity.WalletUseUpRequest) entity.WalletUseUpResponse
	Balance(*entity.WalletBalanceRequest) entity.WalletBalanceResponse
	History(*entity.WalletHistoryRequest) entity.WalletHistoryResponse
}

type walletServiceImpl struct {
	Wallet repository.WalletRepository
}

func NewWalletService(wallet repository.WalletRepository) WalletService {
	return &walletServiceImpl{
		Wallet: wallet,
	}
}

func (service *walletServiceImpl) TopUp(entityRequest *entity.WalletTopUpRequest) (entityResponse entity.WalletTopUpResponse) {
	modelRequest := model.WalletTopUpRequest{
		Context:         entityRequest.Context,
		TransactionCode: util.GenerateTransCode(),
		Description:     entityRequest.Description,
		CoaDebit:        util.CoaKas,
		CoaCredit:       util.CoaUtang,
		Amount:          entityRequest.Amount,
		UserID:          entityRequest.UserID,
		CompanyID:       entityRequest.CompanyID,
		CreatedBy:       entityRequest.CreatedBy,
	}
	modelResponse := service.Wallet.TopUp(&modelRequest)
	if modelResponse.Error != nil {
		entityResponse.Error = errors.New("internal server error")
		return
	}
	return
}

func (service *walletServiceImpl) UseUp(entityRequest *entity.WalletUseUpRequest) (entityResponse entity.WalletUseUpResponse) {
	modelRequest := model.WalletUseUpRequest{
		Context:         entityRequest.Context,
		TransactionCode: util.GenerateTransCode(),
		Description:     entityRequest.Description,
		CoaDebit:        util.CoaUtang,
		CoaCredit:       util.CoaKas,
		Amount:          entityRequest.Amount,
		UserID:          entityRequest.UserID,
		CompanyID:       entityRequest.CompanyID,
		CreatedBy:       entityRequest.CreatedBy,
	}
	modelResponse := service.Wallet.UseUp(&modelRequest)
	if modelResponse.Error != nil {
		entityResponse.Error = errors.New("internal server error")
		return
	}
	return
}

func (service *walletServiceImpl) Balance(entityRequest *entity.WalletBalanceRequest) (entityResponse entity.WalletBalanceResponse) {
	modelRequest := model.WalletBalanceRequest{
		Context: entityRequest.Context,
		CoaKas:  util.CoaKas,
		UserID:  entityRequest.UserID,
	}
	modelResponse := service.Wallet.Balance(&modelRequest)
	if modelResponse.Error != nil {
		entityResponse.Error = errors.New("internal server error")
		return
	}
	entityResponse.Balance = modelResponse.Balance
	return
}

func (service *walletServiceImpl) History(entityRequest *entity.WalletHistoryRequest) (entityResponse entity.WalletHistoryResponse) {
	page := entityRequest.Page
	if page < 1 {
		page = 1
	}
	limit := entityRequest.Limit
	if limit < 1 {
		limit = 1
	}
	offset := limit * (page - 1)
	modelRequest := model.WalletHistoryRequest{
		Context: entityRequest.Context,
		CoaKas:  util.CoaKas,
		UserID:  entityRequest.UserID,
		Limit:   limit,
		Offset:  offset,
	}
	modelResponse := service.Wallet.History(&modelRequest)
	if modelResponse.Error != nil {
		entityResponse.Error = errors.New("internal server error")
		return
	}
	var total int
	if len(modelResponse.Items) != 0 {
		total = modelResponse.Items[0].Total
	}
	entityResponse = entity.WalletHistoryResponse{
		Items: modelResponse.Items,
		Total: total,
		Page:  page,
		Limit: limit,
	}
	return
}
