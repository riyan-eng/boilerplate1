package main

import (
	"boilerplate/config"
	"boilerplate/migration"
	"boilerplate/module/finance"
	"boilerplate/module/management"
	"boilerplate/module/other"
	"runtime"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "boilerplate/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}
	config.ConnDatabase()
	migration.Create()
	config.ConnRedis()
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()
	// middleware.Fiber(app)
	app.Use(logger.New())
	app.Use(recover.New())

	enforcer := config.Casbin()

	api := app.Group("/api")                      // route
	management.Setup(api, enforcer)               // management route
	finance.Setup(api, enforcer)                  // finance route
	other.Setup(api)                              // finance route
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}
