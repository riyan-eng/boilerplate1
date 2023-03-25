package repository

import (
	"boilerplate/module/finance/repository/model"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan/v2"
)

type WalletRepository interface {
	Balance(*model.WalletBalanceRequest) model.WalletBalanceResponse
	History(*model.WalletHistoryRequest) model.WalletHistoryResponse
	Transfer()
	TopUp(*model.WalletTopUpRequest) model.WalletTopUpResponse
	UseUp(*model.WalletUseUpRequest) model.WalletUseUpResponse
}

type walletRepositoryImpl struct {
	Database *sql.DB
}

func NewWalletRepository(database *sql.DB) WalletRepository {
	return &walletRepositoryImpl{
		Database: database,
	}
}

func (repository *walletRepositoryImpl) Balance(modelRequest *model.WalletBalanceRequest) (modelResponse model.WalletBalanceResponse) {
	query := fmt.Sprintf(`
		select coalesce(sum(cash.amount), 0) as balance from (
			select j.transaction_code, amount from journals j 
			where j.user_id = '%v' and j.coa_debit = '%v'
			union 
			select j.transaction_code, -amount from journals j 
			where j.user_id = '%v' and j.coa_credit = '%v'
		) as cash
	`, modelRequest.UserID, modelRequest.CoaKas, modelRequest.UserID, modelRequest.CoaKas)
	row, err := repository.Database.QueryContext(modelRequest.Context, query)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err := scan.Row(&modelResponse.Balance, row); err != nil {
		modelResponse.Error = err
		return
	}
	return
}

func (repository *walletRepositoryImpl) Transfer() {

}

func (repository *walletRepositoryImpl) History(modelRequest *model.WalletHistoryRequest) (modelResponse model.WalletHistoryResponse) {
	query := fmt.Sprintf(`
		select *, count(*) over() as total from (
			select j.transaction_code, t.description, j.amount, j.created_at from journals j 
			left join transactions t on t.code = j.transaction_code
			where j.user_id = '%v' and j.coa_debit = '%v'
			union
			select j.transaction_code, t.description, -j.amount, j.created_at from journals j 
			left join transactions t on t.code = j.transaction_code
			where j.user_id = '%v' and j.coa_credit = '%v'
		) as trx order by trx.created_at desc limit '%v' offset '%v'
	`, modelRequest.UserID, modelRequest.CoaKas, modelRequest.UserID, modelRequest.CoaKas, modelRequest.Limit, modelRequest.Offset)
	rows, err := repository.Database.QueryContext(modelRequest.Context, query)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err := scan.Rows(&modelResponse.Items, rows); err != nil {
		modelResponse.Error = err
		return
	}
	return
}

func (repository *walletRepositoryImpl) TopUp(modelRequest *model.WalletTopUpRequest) (modelResponse model.WalletTopUpResponse) {
	queryTransaction := fmt.Sprintf(`
		insert into transactions(code, description, amount, user_id, company_id, created_by)
		values('%v', '%v', '%v', '%v', '%v', '%v')
	`, modelRequest.TransactionCode, modelRequest.Description, modelRequest.Amount, modelRequest.UserID, modelRequest.CompanyID, modelRequest.CreatedBy)
	queryJournal := fmt.Sprintf(`
		insert into journals(transaction_code, coa_debit, coa_credit, amount, user_id, company_id, created_by)
		values('%v', '%v', '%v', '%v', '%v', '%v', '%v')
	`, modelRequest.TransactionCode, modelRequest.CoaDebit, modelRequest.CoaCredit, modelRequest.Amount, modelRequest.UserID, modelRequest.CompanyID, modelRequest.CreatedBy)
	tx, err := repository.Database.BeginTx(modelRequest.Context, nil)
	if err != nil {
		modelResponse.Error = err
		return
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(modelRequest.Context, queryTransaction)
	if err != nil {
		modelResponse.Error = err
		return
	}
	_, err = tx.ExecContext(modelRequest.Context, queryJournal)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err = tx.Commit(); err != nil {
		modelResponse.Error = err
		return
	}
	return
}

func (repository *walletRepositoryImpl) UseUp(modelRequest *model.WalletUseUpRequest) (modelResponse model.WalletUseUpResponse) {
	queryTransaction := fmt.Sprintf(`
		insert into transactions(code, description, amount, user_id, company_id, created_by)
		values('%v', '%v', '%v', '%v', '%v', '%v')
	`, modelRequest.TransactionCode, modelRequest.Description, modelRequest.Amount, modelRequest.UserID, modelRequest.CompanyID, modelRequest.CreatedBy)
	queryJournal := fmt.Sprintf(`
		insert into journals(transaction_code, coa_debit, coa_credit, amount, user_id, company_id, created_by)
		values('%v', '%v', '%v', '%v', '%v', '%v', '%v')
	`, modelRequest.TransactionCode, modelRequest.CoaDebit, modelRequest.CoaCredit, modelRequest.Amount, modelRequest.UserID, modelRequest.CompanyID, modelRequest.CreatedBy)
	tx, err := repository.Database.BeginTx(modelRequest.Context, nil)
	if err != nil {
		modelResponse.Error = err
		return
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(modelRequest.Context, queryTransaction)
	if err != nil {
		modelResponse.Error = err
		return
	}
	_, err = tx.ExecContext(modelRequest.Context, queryJournal)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err = tx.Commit(); err != nil {
		modelResponse.Error = err
		return
	}
	return
}
