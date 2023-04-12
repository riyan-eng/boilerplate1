package controller

import (
	"boilerplate/module/other/controller/dto"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type ExcelizeController interface {
	Download(*fiber.Ctx) error
	Upload(*fiber.Ctx) error
}

type excelizeControllerImpl struct{}

func NewExcelizeController() ExcelizeController {
	return &excelizeControllerImpl{}
}

func (controller *excelizeControllerImpl) Download(c *fiber.Ctx) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	category := []string{"ANTIPIRAI NON", "ANTIINFLAMASI NON STEROID GANTI", "ANALGESIK NON NON NARKOTIK", "ANTIPIRETIK NON NARKOTIK", "ANTIPIRETIK", "ANTIPIRAI", "ANALGESIK NON NARKOTIK", "ANALGESIK NARKOTIK"}
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "No.")
	f.SetCellValue("Sheet1", "B1", "Nama Obat")
	f.SetCellValue("Sheet1", "C1", "Kategori Obat")
	f.SetCellValue("Sheet1", "E2", "Catatan")
	f.SetCellValue("Sheet1", "E3", "Masukkan kategori obat sesuai kategori yang tersedia")
	f.SetCellValue("Sheet1", "E5", "Kategori tersedia")
	for i, item := range category {
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(6+i), item)
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// set style
	// style, _ := f.NewStyle()
	style1, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center"},
		Font:      &excelize.Font{Bold: true},
	})
	style2, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	f.SetCellStyle("Sheet1", "A1", "A1", style1)
	f.SetCellStyle("Sheet1", "B1", "B1", style1)
	f.SetCellStyle("Sheet1", "C1", "C1", style1)
	f.SetCellStyle("Sheet1", "E2", "E2", style2)
	f.SetCellStyle("Sheet1", "E5", "E5", style2)
	f.SetColWidth("Sheet1", "A", "A", 5)
	f.SetColWidth("Sheet1", "B", "B", 15)
	f.SetColWidth("Sheet1", "C", "C", 20)
	f.SetColWidth("Sheet1", "E", "E", 45)

	c.Response().Header.Set("Content-Type", "application/octet-stream")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=template_obat.xlsx")
	c.Response().Header.Set("File-Name", "template_obat.xlsx")
	c.Response().Header.Set("Content-Transfer-Encoding", "binary")
	c.Response().Header.Set("Expires", "0")
	return f.Write(c)
}

func (controller *excelizeControllerImpl) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	buffer, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	f, err := excelize.OpenReader(buffer)
	if err != nil {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// fmt.Println("la", rows[0])
	dtoRequest := dto.ExcelizeUploadRequest{}
	var item dto.ExcelizeUploadItem
	for _, row := range rows[1:] {
		data := row[1:3]
		if data[0] != "" && data[1] != "" {
			item.Name = data[0]
			item.Category = data[1]
			dtoRequest.Item = append(dtoRequest.Item, item)
		}
	}
	fmt.Println(dtoRequest)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "ok",
	})
}
