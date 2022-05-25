package Home

import (
	"trashgo/Database"
	"trashgo/Models"

	"github.com/gofiber/fiber/v2"
)

func Laporan(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	laporan := Models.Laporan{
		
		Data_laporan: data["data_laporan"],
		Id_komplek:   data["id_komplek"],
		Id_kendaraan: data["id_kendaraan"],
		Id_supir:     data["id_supir"],
	}

	Database.DB.Create(&laporan)

	return c.JSON(fiber.Map{
		"message": "Terima Kasih laporan anda akan di proses ",
	})
}