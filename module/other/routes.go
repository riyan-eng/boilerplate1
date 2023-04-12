package other

import (
	"boilerplate/module/other/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	excelizeController := controller.NewExcelizeController()
	excelize := router.Group("/excelize")
	excelize.Get("/download", excelizeController.Download)
	excelize.Post("/upload", excelizeController.Upload)
}
