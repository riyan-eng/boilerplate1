package model

import "github.com/valyala/fasthttp"

type LoginRequest struct {
	Context  *fasthttp.RequestCtx
	Username string
	Password string
}

type LoginResponse struct {
	Error error
	User  loginResponseUser
}

type loginResponseUser struct {
	ID        string `gorm:"column:id"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	UserType  string `gorm:"column:user_type"`
	CompanyID string `gorm:"column:company_id"`
}

type RegisterRequest struct {
	Context     *fasthttp.RequestCtx
	Username    string
	Password    string
	UserType    string
	CompanyID   string
	Email       string
	PhoneNumber string
}

type RegisterResponse struct {
	User  registerResponseUser
	Error error
}

type registerResponseUser struct {
	ID       string `gorm:"column:id"`
	UserType string `gorm:"column:user_type_code"`
}
