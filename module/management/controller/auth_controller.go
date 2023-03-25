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

// @Tags				Authentication
// @Summary			Create new user based on paramters
// @Description	Create new user
// @Accept			json
// @Param				account	body	dto.LoginRequest	true	"Add account"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /api/auth/login [post]
func (controller *authenticationControllerImpl) Login(c *fiber.Ctx) error {
	dtoRequest := new(dto.LoginRequest)
	if err := c.BodyParser(&dtoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err.Error(), Message: "bad"},
		)
	}
	if err := util.ValidateRequest(dtoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: "err", Message: "bad"},
		)
	}
	entityRequest := entity.LoginRequest{
		Context: c.Context(), Username: dtoRequest.Username, Password: dtoRequest.Password, Issuer: string(c.Request().Host()),
	}
	entityResponse := controller.Authentication.Login(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: entityResponse.Error.Error(), Message: "bad"},
		)
	}
	dtoResponse := dto.LoginResponse{
		AccessToken: entityResponse.AccessToken, RefreshToken: entityResponse.RefreshToken,
	}
	return c.Status(fiber.StatusOK).JSON(
		util.Response{Data: dtoResponse, Message: "ok"},
	)
}

// @Tags				Authentication
// @Summary			Create new user based on paramters
// @Description	Create new user
// @Accept			json
// @Param				account	body	dto.RefreshRequest	true	"Add account"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router /api/auth/refresh_token [post]
func (controller *authenticationControllerImpl) Refresh(c *fiber.Ctx) error {
	dtoRequest := new(dto.RefreshRequest)
	if err := c.BodyParser(&dtoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err.Error(), Message: "bad"},
		)
	}
	if err := util.ValidateRequest(dtoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: err, Message: "bad"},
		)
	}
	entityRequest := entity.RefreshRequest{
		Context: c.Context(), RefreshToken: dtoRequest.RefreshToken, Issuer: string(c.Request().Host()),
	}
	entityResponse := controller.Authentication.Refresh(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			util.Response{Data: entityResponse.Error.Error(), Message: "bad"},
		)
	}
	dtoResponse := dto.RefreshResponse{
		AccessToken: entityResponse.AccessToken, RefreshToken: entityResponse.RefreshToken,
	}
	return c.Status(fiber.StatusOK).JSON(
		util.Response{Data: dtoResponse, Message: "ok"},
	)
}

func (controller *authenticationControllerImpl) Register(enforcer *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dtoRequest := new(dto.RegisterRequest)
		if err := c.BodyParser(&dtoRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: err.Error(), Message: "bad"},
			)
		}
		if err := util.ValidateRequest(dtoRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: err, Message: "bad"},
			)
		}
		entityRequest := entity.RegisterRequest{
			Context: c.Context(), Enforcer: enforcer, Username: dtoRequest.Username, Password: dtoRequest.Password, CompanyID: dtoRequest.CompanyID, Email: dtoRequest.Email, PhoneNumber: dtoRequest.PhoneNumber,
		}
		entityResponse := controller.Authentication.Register(&entityRequest)
		if entityResponse.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				util.Response{Data: entityResponse.Error.Error(), Message: "bad"},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			util.Response{Data: "successfully insert data", Message: "ok"},
		)
	}
}
