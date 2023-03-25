package repository

import (
	"boilerplate/module/management/repository/model"
	"database/sql"
	"fmt"

	"github.com/blockloop/scan/v2"
)

type AuthenticationRepository interface {
	Login(*model.LoginRequest) model.LoginResponse
	Register(*model.RegisterRequest) model.RegisterResponse
}

type authenticationRepository struct {
	Database *sql.DB
}

func NewAuthenticationRepository(database *sql.DB) AuthenticationRepository {
	return &authenticationRepository{
		Database: database,
	}
}

func (repository *authenticationRepository) Login(modelRequest *model.LoginRequest) (modelResponse model.LoginResponse) {
	query := fmt.Sprintf(`select u.id, u.username, u.user_type_code as user_type, u."password", coalesce(u.company_id, '') as company_id from users u 
	where u.username like '%%%v' or u.email like '%%%v' or u.phone_number like '%%%v'`, modelRequest.Username, modelRequest.Username, modelRequest.Username)
	row, err := repository.Database.QueryContext(modelRequest.Context, query)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err := scan.Row(&modelResponse.User, row); err != nil {
		modelResponse.Error = err
	}
	return
}

func (repository *authenticationRepository) Register(modelRequest *model.RegisterRequest) (modelResponse model.RegisterResponse) {
	query := fmt.Sprintf(`
	insert into users(username, user_type_code, email, password, phone_number, company_id)values
	('%v', '%v', '%v', '%v', '%v', '%v') 
	returning user_type_code, id
	`, modelRequest.Username, "user", modelRequest.Email, modelRequest.Password, modelRequest.PhoneNumber, modelRequest.CompanyID)
	row, err := repository.Database.QueryContext(modelRequest.Context, query)
	if err != nil {
		modelResponse.Error = err
		return
	}
	if err := scan.Row(&modelResponse.User, row); err != nil {
		modelResponse.Error = err
	}
	return
}
