package repository

import (
	"boilerplate/module/finance/repository/model"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Balance()
	History()
	TopUp(*model.WalletTopUpRequest) model.WalletTopUpResponse
	UseUp(*model.WalletUseUpRequest) model.WalletUseUpResponse
}

type walletRepositoryImpl struct {
	Database *gorm.DB
}

func NewWalletRepository(database *gorm.DB) WalletRepository {
	return &walletRepositoryImpl{
		Database: database,
	}
}

func (repository *walletRepositoryImpl) Balance() {

}

func (repository *walletRepositoryImpl) History() {

}

func (repository *walletRepositoryImpl) TopUp(walletModelRequest *model.WalletTopUpRequest) (walletModelResponse model.WalletTopUpResponse) {
	query1 := fmt.Sprintf(`
		insert into transactions(code, description, amount)
		values('%v', '%v', '%v')
	`, walletModelRequest.TransactionCode, walletModelRequest.Description, walletModelRequest.Amount)
	query2 := fmt.Sprintf(`
		insert into journals(transaction_code, coa_debit, coa_credit, amount)
		values('%v', '%v', '%v', '%v')
	`, walletModelRequest.TransactionCode, walletModelRequest.CoaDebit, walletModelRequest.CoaCredit, walletModelRequest.Amount)
	tx := repository.Database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		walletModelResponse.Error = err
		return
	}
	if err := tx.Raw(query1).Error; err != nil {
		tx.Rollback()
		walletModelResponse.Error = err
		return
	}
	if err := tx.Raw(query2).Error; err != nil {
		tx.Rollback()
		walletModelResponse.Error = err
		return
	}
	walletModelResponse.Error = tx.Commit().Error
	return
}

func (repository *walletRepositoryImpl) UseUp(walletModelRequest *model.WalletUseUpRequest) (walletModelResponse model.WalletUseUpResponse) {
	query1 := fmt.Sprintf(`
		insert into transactions(code, description, amount)
		values('%v', '%v', '%v')
	`, walletModelRequest.TransactionCode, walletModelRequest.Description, walletModelRequest.Amount)
	query2 := fmt.Sprintf(`
		insert into journals(transaction_code, coa_debit, coa_credit, amount)
		values('%v', '%v', '%v', '%v')
	`, walletModelRequest.TransactionCode, walletModelRequest.CoaDebit, walletModelRequest.CoaCredit, walletModelRequest.Amount)
	tx := repository.Database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		walletModelResponse.Error = err
		return
	}
	if err := tx.Raw(query1).Error; err != nil {
		tx.Rollback()
		walletModelResponse.Error = err
		return
	}
	if err := tx.Raw(query2).Error; err != nil {
		tx.Rollback()
		walletModelResponse.Error = err
		return
	}
	walletModelResponse.Error = tx.Commit().Error
	return
}
