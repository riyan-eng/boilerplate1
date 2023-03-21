package main

import (
	"boilerplate/config"
	"boilerplate/module/management"
	"runtime"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
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
	// middleware.Fiber(app)
	app.Use(logger.New())
	app.Use(recover.New())

	enforcer := config.Casbin()

	// app.Get("/", middleware.AuthorizeJwt(), middleware.AuthorizeCasbin(enforcer), func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	api := app.Group("/api")        // route
	management.Setup(api, enforcer) // management route

	app.Listen(":3000")
}
