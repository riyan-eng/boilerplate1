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
	ID        string `db:"id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	UserType  string `db:"user_type"`
	CompanyID string `db:"company_id"`
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
	ID       string `db:"id"`
	UserType string `db:"user_type_code"`
}
