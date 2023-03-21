package management

import (
	"boilerplate/config"
	"boilerplate/middleware"
	"boilerplate/module/management/controller"
	"boilerplate/module/management/repository"
	"boilerplate/module/management/service"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router, enforcer *casbin.Enforcer) {
	authRepository := repository.NewAuthenticationRepository(config.Database)
	authService := service.NewAuthenticationService(authRepository)
	authController := controller.NewAuthenticationController(authService)

	auth := router.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Post("/refresh_token", authController.Refresh)
	auth.Post("/register", authController.Register(enforcer))

	admin := router.Group("/admin")
	adminUser := admin.Group("/users", middleware.AuthorizeJwt())
	adminUser.Get("/", middleware.AuthorizeCasbin(enforcer), controller.GetUsers)
	adminUser.Get("/:id", middleware.AuthorizeCasbin(enforcer), controller.GetUser)
	adminUser.Post("/", middleware.AuthorizeCasbin(enforcer), controller.CreateUser(enforcer))
	adminUser.Put("/", middleware.AuthorizeCasbin(enforcer), controller.UpdateUser(enforcer))
	adminUser.Delete("/:id", middleware.AuthorizeCasbin(enforcer), controller.DeleteUser(enforcer))

	user := router.Group("/users", middleware.AuthorizeJwt())
	user.Get("/:id", middleware.AuthorizeCasbin(enforcer), controller.GetUser)
}
