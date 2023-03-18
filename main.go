package main

import (
	"boilerplate/config"
	"boilerplate/middleware"
	"boilerplate/module/auth"
	"boilerplate/module/management"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}
	config.ConnDatabase()
	config.ConnRedis()
}

func main() {
	app := fiber.New()
	app.Use(recover.New())

	enforcer := config.Casbin()

	app.Get("/", middleware.AuthorizeJwt(), middleware.AuthorizeCasbin(enforcer), func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")        // Public route
	management.Setup(api, enforcer) // Admin route
	auth.Setup(api, enforcer)

	app.Listen(":3000")
}
