package controller

import (
	"boilerplate/module/management/controller/dto"
	"boilerplate/module/management/service"
	"boilerplate/module/management/service/entity"
	"boilerplate/util"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

type AuthenticationController interface {
	Login(*fiber.Ctx) error
	Refresh(*fiber.Ctx) error
	Register(*casbin.Enforcer) fiber.Handler
}
type authenticationControllerImpl struct {
	Authentication service.AuthenticationService
}

func NewAuthenticationController(authentocation service.AuthenticationService) AuthenticationController {
	return &authenticationControllerImpl{
		Authentication: authentocation,
	}
}

func (controller *authenticationControllerImpl) Login(c *fiber.Ctx) error {
	loginControllerRequest := new(dto.LoginRequest)
	if err := c.BodyParser(&loginControllerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err.Error(), Message: "bad"},
		)
	}
	if err := util.ValidateRequest(loginControllerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: "err", Message: "bad"},
		)
	}
	loginServiceRequest := entity.LoginRequest{
		Context: c.Context(), Username: loginControllerRequest.Username, Password: loginControllerRequest.Password, Issuer: string(c.Request().Host()),
	}
	loginServiceResponse := controller.Authentication.Login(&loginServiceRequest)
	if loginServiceResponse.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: loginServiceResponse.Error.Error(), Message: "bad"},
		)
	}
	loginControllerResponse := dto.LoginResponse{
		AccessToken: loginServiceResponse.AccessToken, RefreshToken: loginServiceResponse.RefreshToken,
	}
	return c.Status(fiber.StatusOK).JSON(
		util.Response{Data: loginControllerResponse, Message: "ok"},
	)
}

func (controller *authenticationControllerImpl) Refresh(c *fiber.Ctx) error {
	refreshControllerRequest := new(dto.RefreshRequest)
	if err := c.BodyParser(&refreshControllerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err.Error(), Message: "bad"},
		)
	}
	if err := util.ValidateRequest(refreshControllerRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err, Message: "bad"},
		)
	}
	refreshServiceRequest := entity.RefreshRequest{
		Context: c.Context(), RefreshToken: refreshControllerRequest.RefreshToken, Issuer: string(c.Request().Host()),
	}
	refreshServiceResponse := controller.Authentication.Refresh(&refreshServiceRequest)
	if refreshServiceResponse.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: refreshServiceResponse.Error.Error(), Message: "bad"},
		)
	}
	refreshControllerResponse := dto.RefreshResponse{
		AccessToken: refreshServiceResponse.AccessToken, RefreshToken: refreshServiceResponse.RefreshToken,
	}
	return c.Status(fiber.StatusOK).JSON(
		util.Response{Data: refreshControllerResponse, Message: "ok"},
	)
}

func (controller *authenticationControllerImpl) Register(enforcer *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		registerControllerRequest := new(dto.RegisterRequest)
		if err := c.BodyParser(&registerControllerRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: err.Error(), Message: "bad"},
			)
		}
		if err := util.ValidateRequest(registerControllerRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: err, Message: "bad"},
			)
		}
		registerServiceRequest := entity.RegisterRequest{
			Context: c.Context(), Enforcer: enforcer, Username: registerControllerRequest.Username, Password: registerControllerRequest.Password, CompanyID: registerControllerRequest.CompanyID, Email: registerControllerRequest.Email, PhoneNumber: registerControllerRequest.PhoneNumber,
		}
		registerServiceResponse := controller.Authentication.Register(&registerServiceRequest)
		if registerServiceResponse.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: registerServiceResponse.Error.Error(), Message: "bad"},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			util.Response{Data: "successfully insert data", Message: "ok"},
		)
	}
}

// func Login(c *fiber.Ctx) error {
// 	userID := c.Query("user_id", "1")
// 	companyID := c.Query("company_id", "comp_123")

// 	// Create JWT token with userID.
// 	accessToken, refreshToken, err := util.GenerateJWT(string(c.Request().Host()), userID, companyID, 1)

// 	// accessToken, refreshToken, exp, err := util.GenerateTokenPair(string(c.Request().Host()), userID, "comp1", c.Context())
// 	if err != nil {
// 		c.Status(fiber.StatusInternalServerError)
// 		return c.JSON(fiber.Map{
// 			"data":    "internal server error",
// 			"message": "bad",
// 		})
// 	}

// 	// Create cookie
// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    accessToken,
// 		Expires:  time.Now().Add(time.Minute * 1),
// 		HTTPOnly: true,
// 	}

// 	// save cookie
// 	c.Cookie(&cookie)
// 	return c.JSON(fiber.Map{
// 		"data": fiber.Map{
// 			"accessToken":  accessToken,
// 			"refreshToken": refreshToken,
// 			"exp":          "exp",
// 		},
// 		"message": "ok",
// 	})
// }

// func RefreshToken(c *fiber.Ctx) error {
// 	refreshRequest := new(dto.RefreshRequest)
// 	if err := c.BodyParser(refreshRequest); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"data":    "bad request",
// 			"message": "bad",
// 		})
// 	}

// 	claims, err := util.ParseToken(refreshRequest.RefreshToken, "AllYourBaseRefresh")
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "not authorized!",
// 		})
// 	}

// 	if err := util.ValidateToken(claims, true, c.Context()); err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "not authorized!",
// 		})
// 	}

// 	accessToken, refreshToken, err := util.GenerateJWT(string(c.Request().Host()), claims.UserID, claims.CompanyID, 3)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"data":    "internal server error",
// 			"message": "bad",
// 		})
// 	}

// 	// Create cookie
// 	cookie := fiber.Cookie{
// 		Name:     "jwt",
// 		Value:    accessToken,
// 		Expires:  time.Now().Add(time.Minute * 1),
// 		HTTPOnly: true,
// 	}

// 	// save cookie
// 	c.Cookie(&cookie)
// 	return c.JSON(fiber.Map{
// 		"data": fiber.Map{
// 			"accessToken":  accessToken,
// 			"refreshToken": refreshToken,
// 			"exp":          "exp",
// 		},
// 		"message": "ok",
// 	})
// }
