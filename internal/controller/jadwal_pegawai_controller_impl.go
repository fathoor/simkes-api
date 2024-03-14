package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type jadwalPegawaiControllerImpl struct {
	usecase.JadwalPegawaiService
}

func (controller *jadwalPegawaiControllerImpl) Create(c *fiber.Ctx) error {
	var request web.JadwalPegawaiRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.JadwalPegawaiService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *jadwalPegawaiControllerImpl) Get(c *fiber.Ctx) error {
	nip := c.Query("nip")
	tahun := c.QueryInt("tahun")
	bulan := c.QueryInt("bulan")

	switch {
	case nip != "":
		response := controller.JadwalPegawaiService.GetByNIP(nip)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	case tahun != 0 && bulan != 0:
		response := controller.JadwalPegawaiService.GetByTahunBulan(int16(tahun), int16(bulan))

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	default:
		response := controller.JadwalPegawaiService.GetAll()

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *jadwalPegawaiControllerImpl) GetByPK(c *fiber.Ctx) error {
	nip := c.Params("nip")

	tahun, err := c.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := c.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := c.ParamsInt("hari")
	exception.PanicIfError(err)

	response := controller.JadwalPegawaiService.GetByPK(nip, int16(tahun), int16(bulan), int16(hari))

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPegawaiControllerImpl) Update(c *fiber.Ctx) error {
	var request web.JadwalPegawaiRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := c.Params("nip")

	tahun, err := c.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := c.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := c.ParamsInt("hari")
	exception.PanicIfError(err)

	response := controller.JadwalPegawaiService.Update(nip, int16(tahun), int16(bulan), int16(hari), &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jadwalPegawaiControllerImpl) Delete(c *fiber.Ctx) error {
	nip := c.Params("nip")

	tahun, err := c.ParamsInt("tahun")
	exception.PanicIfError(err)

	bulan, err := c.ParamsInt("bulan")
	exception.PanicIfError(err)

	hari, err := c.ParamsInt("hari")
	exception.PanicIfError(err)

	controller.JadwalPegawaiService.Delete(nip, int16(tahun), int16(bulan), int16(hari))

	return c.SendStatus(fiber.StatusNoContent)
}

func NewJadwalPegawaiControllerProvider(service *usecase.JadwalPegawaiService) JadwalPegawaiController {
	return &jadwalPegawaiControllerImpl{*service}
}
