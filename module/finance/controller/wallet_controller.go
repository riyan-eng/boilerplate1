package controller

import (
	"boilerplate/module/finance/controller/dto"
	"boilerplate/module/finance/service"
	"boilerplate/module/finance/service/entity"
	"boilerplate/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type WalletController interface {
	TopUp(*fiber.Ctx) error
	UseUp(*fiber.Ctx) error
	Balance(*fiber.Ctx) error
	History(*fiber.Ctx) error
}

type walletControllerImpl struct {
	Wallet service.WalletService
}

func NewWalletController(wallet service.WalletService) WalletController {
	return &walletControllerImpl{
		Wallet: wallet,
	}
}

func (controller *walletControllerImpl) TopUp(c *fiber.Ctx) error {
	currentUser := c.Locals("userID").(string)
	companyID := c.Locals("companyID").(string)
	dtoRequest := new(dto.WalletTopUpRequest)
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
	entityRequest := entity.WalletTopUpRequest{
		Context:     c.Context(),
		Description: dtoRequest.Description,
		Amount:      dtoRequest.Amount,
		UserID:      dtoRequest.UserID,
		CompanyID:   companyID,
		CreatedBy:   currentUser,
	}
	entityResponse := controller.Wallet.TopUp(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(util.Response{
			Data:    entityResponse.Error.Error(),
			Message: "bad",
		})
	}
	return c.Status(fiber.StatusOK).JSON(util.Response{
		Data:    "success insert data",
		Message: "ok",
	})
}

func (controller *walletControllerImpl) UseUp(c *fiber.Ctx) error {
	currentUser := c.Locals("userID").(string)
	companyID := c.Locals("companyID").(string)
	dtoRequest := new(dto.WalletUseUpRequest)
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
	entityRequest := entity.WalletUseUpRequest{
		Context:     c.Context(),
		Description: dtoRequest.Description,
		Amount:      dtoRequest.Amount,
		UserID:      currentUser,
		CompanyID:   companyID,
		CreatedBy:   currentUser,
	}
	entityResponse := controller.Wallet.UseUp(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(util.Response{
			Data:    entityResponse.Error.Error(),
			Message: "bad",
		})
	}
	return c.Status(fiber.StatusOK).JSON(util.Response{
		Data:    "success insert data",
		Message: "ok",
	})
}

func (controller *walletControllerImpl) Balance(c *fiber.Ctx) error {
	currentUser := c.Locals("userID").(string)
	entityRequest := entity.WalletBalanceRequest{Context: c.Context(), UserID: currentUser}
	entityResponse := controller.Wallet.Balance(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(util.Response{
			Data:    entityResponse.Error.Error(),
			Message: "bad",
		})
	}
	dtoResponse := dto.WalletBalanceResponse{Balance: entityResponse.Balance}
	return c.Status(fiber.StatusOK).JSON(util.Response{
		Data:    dtoResponse,
		Message: "ok",
	})
}

func (controller *walletControllerImpl) History(c *fiber.Ctx) error {
	currentUser := c.Locals("userID").(string)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	entityRequest := entity.WalletHistoryRequest{Context: c.Context(), UserID: currentUser, Page: page, Limit: limit}
	entityResponse := controller.Wallet.History(&entityRequest)
	if entityResponse.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(util.Response{
			Data:    entityResponse.Error.Error(),
			Message: "bad",
		})
	}
	dtoResponse := dto.WalletHistoryResponse{
		Items: entityResponse.Items,
		Page:  entityRequest.Page,
		Limit: entityResponse.Limit,
		Total: entityResponse.Total,
	}
	return c.Status(fiber.StatusOK).JSON(util.Response{
		Data:    dtoResponse,
		Message: "ok",
	})
}
