package repository

import (
	"boilerplate/module/management/repository/model"
	"fmt"

	"gorm.io/gorm"
)

type AuthenticationRepository interface {
	Login(*model.LoginRequest) *model.LoginResponse
	Register(*model.RegisterRequest) *model.RegisterResponse
}

type authenticationRepository struct {
	Database *gorm.DB
}

func NewAuthenticationRepository(database *gorm.DB) AuthenticationRepository {
	return &authenticationRepository{
		Database: database,
	}
}

func (repository *authenticationRepository) Login(loginModelRequest *model.LoginRequest) (loginModelResponse *model.LoginResponse) {
	query := `select * from management.user`
	loginModelResponse.Error = repository.Database.WithContext(loginModelRequest.Context).Raw(query).Scan(&loginModelResponse.User).Error
	return
}

func (repository *authenticationRepository) Register(registerModelRequest *model.RegisterRequest) (registerModelResponse *model.RegisterResponse) {
	registerModelResponse = &model.RegisterResponse{}
	query := fmt.Sprintf(`
	insert into users(username, user_type_code, email, password, phone_number, company_id)values
	('%v', '%v', '%v', '%v', '%v', '%v') 
	returning user_type_code, id
	`, registerModelRequest.Username, "user", registerModelRequest.Email, registerModelRequest.Password, registerModelRequest.PhoneNumber, registerModelRequest.CompanyID)
	if err := repository.Database.WithContext(registerModelRequest.Context).Raw(query).Scan(&registerModelResponse.User).Error; err != nil {
		registerModelResponse.Error = err
	}
	return
}
