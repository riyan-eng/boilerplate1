package service

import (
	"boilerplate/module/management/repository"
	"boilerplate/module/management/repository/model"
	"boilerplate/module/management/service/entity"
	"boilerplate/util"
	"errors"
)

type AuthenticationService interface {
	Login(*entity.LoginRequest) *entity.LoginReponse
	Refresh(*entity.RefreshRequest) *entity.RefreshReponse
	Register(*entity.RegisterRequest) *entity.RegisterResponse
}

type authenticationServiceImpl struct {
	Authenticaton repository.AuthenticationRepository
}

func NewAuthenticationService(authentication repository.AuthenticationRepository) AuthenticationService {
	return &authenticationServiceImpl{
		Authenticaton: authentication,
	}
}

func (service *authenticationServiceImpl) Login(loginEntityRequest *entity.LoginRequest) (loginEntityReponse *entity.LoginReponse) {
	loginModelRequest := model.LoginRequest{Context: loginEntityRequest.Context, Username: loginEntityRequest.Username, Password: loginEntityRequest.Password}
	loginModelResponse := service.Authenticaton.Login(&loginModelRequest)
	if loginModelResponse.Error != nil {
		loginEntityReponse.Error = errors.New("internal server error")
		return
	}
	loginEntityReponse.AccessToken,
		loginEntityReponse.RefreshToken,
		loginEntityReponse.Error = util.GenerateJWT(loginEntityRequest.Issuer, loginModelResponse.User.ID, loginModelResponse.User.CompanyID, 15)
	return
}

func (service *authenticationServiceImpl) Refresh(refreshEntityRequest *entity.RefreshRequest) (refreshEntityResponse *entity.RefreshReponse) {
	claims, err := util.ParseToken(refreshEntityRequest.RefreshToken, "AllYourBaseRefresh")
	if err != nil {
		refreshEntityResponse.Error = errors.New("not authorized")
		return
	}
	if err := util.ValidateToken(claims, true, refreshEntityRequest.Context); err != nil {
		refreshEntityResponse.Error = errors.New("not authorized")
		return
	}
	refreshEntityResponse.AccessToken,
		refreshEntityResponse.RefreshToken,
		refreshEntityResponse.Error =
		util.GenerateJWT(refreshEntityRequest.Issuer, claims.UserID, claims.CompanyID, 15)
	return
}

func (service *authenticationServiceImpl) Register(registerEntityRequest *entity.RegisterRequest) (registerEntityResponse *entity.RegisterResponse) {
	registerEntityResponse = &entity.RegisterResponse{}
	registerModelRequest := model.RegisterRequest{Context: registerEntityRequest.Context, Username: registerEntityRequest.Username, Password: util.GenerateHash(registerEntityRequest.Password), CompanyID: registerEntityRequest.CompanyID, Email: registerEntityRequest.Email, PhoneNumber: registerEntityRequest.PhoneNumber}
	registerModelResponse := service.Authenticaton.Register(&registerModelRequest)
	if registerModelResponse.Error != nil {
		registerEntityResponse.Error = errors.New("internal server error")
		return
	}
	_, casbinErr := registerEntityRequest.Enforcer.AddGroupingPolicy(registerModelResponse.User.ID, registerModelResponse.User.UserType)
	if casbinErr != nil {
		registerEntityResponse.Error = errors.New("internal server error")
	}
	return
}
