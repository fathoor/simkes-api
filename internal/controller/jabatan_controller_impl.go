package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type jabatanControllerImpl struct {
	usecase.JabatanService
}

func (controller *jabatanControllerImpl) Create(c *fiber.Ctx) error {
	var request web.JabatanRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.JabatanService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *jabatanControllerImpl) Get(c *fiber.Ctx) error {
	response := controller.JabatanService.GetAll()

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jabatanControllerImpl) GetByNama(c *fiber.Ctx) error {
	jabatan := c.Params("jabatan")

	response := controller.JabatanService.GetByJabatan(jabatan)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jabatanControllerImpl) Update(c *fiber.Ctx) error {
	var request web.JabatanRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	jabatan := c.Params("jabatan")

	response := controller.JabatanService.Update(jabatan, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *jabatanControllerImpl) Delete(c *fiber.Ctx) error {
	jabatan := c.Params("jabatan")

	controller.JabatanService.Delete(jabatan)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewJabatanControllerProvider(service *usecase.JabatanService) JabatanController {
	return &jabatanControllerImpl{*service}
}
