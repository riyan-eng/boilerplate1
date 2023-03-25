package finance

import (
	"boilerplate/config"
	"boilerplate/middleware"
	"boilerplate/module/finance/controller"
	"boilerplate/module/finance/repository"
	"boilerplate/module/finance/service"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router, enforcer *casbin.Enforcer) {
	walletRepository := repository.NewWalletRepository(config.PostgreSQLDB)
	walletService := service.NewWalletService(walletRepository)
	walletController := controller.NewWalletController(walletService)

	wallet := router.Group("/wallet", middleware.AuthorizeJwt())
	wallet.Post("/top_up", middleware.AuthorizeCasbin(enforcer), walletController.TopUp)
	wallet.Post("/use_up", middleware.AuthorizeCasbin(enforcer), walletController.UseUp)
	wallet.Get("/balance", middleware.AuthorizeCasbin(enforcer), walletController.Balance)
	wallet.Get("/history", middleware.AuthorizeCasbin(enforcer), walletController.History)
}
