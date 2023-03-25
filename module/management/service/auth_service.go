package service

import (
	"boilerplate/module/management/repository"
	"boilerplate/module/management/repository/model"
	"boilerplate/module/management/service/entity"
	"boilerplate/util"
	"errors"
)

type AuthenticationService interface {
	Login(*entity.LoginRequest) entity.LoginReponse
	Refresh(*entity.RefreshRequest) entity.RefreshReponse
	Register(*entity.RegisterRequest) entity.RegisterResponse
}

type authenticationServiceImpl struct {
	Authenticaton repository.AuthenticationRepository
}

func NewAuthenticationService(authentication repository.AuthenticationRepository) AuthenticationService {
	return &authenticationServiceImpl{
		Authenticaton: authentication,
	}
}

func (service *authenticationServiceImpl) Login(entityRequest *entity.LoginRequest) (entityReponse entity.LoginReponse) {
	loginModelRequest := model.LoginRequest{Context: entityRequest.Context, Username: entityRequest.Username, Password: entityRequest.Password}
	loginModelResponse := service.Authenticaton.Login(&loginModelRequest)
	if loginModelResponse.Error != nil {
		entityReponse.Error = errors.New("internal server error")
		return
	}
	if !util.VerifyHash(loginModelResponse.User.Password, entityRequest.Password) {
		entityReponse.Error = errors.New("invalid usename or password")
		return
	}
	entityReponse.AccessToken,
		entityReponse.RefreshToken,
		entityReponse.Error = util.GenerateJWT(entityRequest.Issuer, loginModelResponse.User.ID, loginModelResponse.User.CompanyID, 15)
	return
}

func (service *authenticationServiceImpl) Refresh(entityRequest *entity.RefreshRequest) (entityResponse entity.RefreshReponse) {
	claims, err := util.ParseToken(entityRequest.RefreshToken, "AllYourBaseRefresh")
	if err != nil {
		entityResponse.Error = errors.New("not authorized")
		return
	}
	if err := util.ValidateToken(claims, true, entityRequest.Context); err != nil {
		entityResponse.Error = errors.New("not authorized")
		return
	}
	entityResponse.AccessToken,
		entityResponse.RefreshToken,
		entityResponse.Error =
		util.GenerateJWT(entityRequest.Issuer, claims.UserID, claims.CompanyID, 15)
	return
}

func (service *authenticationServiceImpl) Register(entityRequest *entity.RegisterRequest) (entityResponse entity.RegisterResponse) {
	registerModelRequest := model.RegisterRequest{Context: entityRequest.Context, Username: entityRequest.Username, Password: util.GenerateHash(entityRequest.Password), CompanyID: entityRequest.CompanyID, Email: entityRequest.Email, PhoneNumber: entityRequest.PhoneNumber}
	registerModelResponse := service.Authenticaton.Register(&registerModelRequest)
	if registerModelResponse.Error != nil {
		entityResponse.Error = errors.New("user has been exist")
		return
	}
	_, casbinErr := entityRequest.Enforcer.AddGroupingPolicy(registerModelResponse.User.ID, registerModelResponse.User.UserType)
	if casbinErr != nil {
		entityResponse.Error = errors.New("internal server error")
	}
	return
}
