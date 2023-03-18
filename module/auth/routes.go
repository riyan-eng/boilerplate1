package auth

import (
	"boilerplate/module/auth/controller"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router, enforcer *casbin.Enforcer) {
	auth := router.Group("/auth")
	auth.Post("/login", controller.Login)
	auth.Post("/refresh_token", controller.RefreshToken)
}
