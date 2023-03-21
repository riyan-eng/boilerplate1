package entity

import (
	"github.com/casbin/casbin/v2"
	"github.com/valyala/fasthttp"
)

type LoginRequest struct {
	Context  *fasthttp.RequestCtx
	Username string
	Password string
	Issuer   string
}

type LoginReponse struct {
	AccessToken  string
	RefreshToken string
	Error        error
}

type RefreshRequest struct {
	Context      *fasthttp.RequestCtx
	RefreshToken string
	Issuer       string
}

type RefreshReponse struct {
	AccessToken  string
	RefreshToken string
	Error        error
}

type RegisterRequest struct {
	Context     *fasthttp.RequestCtx
	Enforcer    *casbin.Enforcer
	Username    string
	Password    string
	CompanyID   string
	Email       string
	PhoneNumber string
}

type RegisterResponse struct {
	Error error
}
