package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type shiftControllerImpl struct {
	usecase.ShiftService
}

func (controller *shiftControllerImpl) Create(c *fiber.Ctx) error {
	var request web.ShiftRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.ShiftService.Create(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Get(c *fiber.Ctx) error {
	response := controller.ShiftService.GetAll()

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) GetByNama(c *fiber.Ctx) error {
	shift := c.Params("shift")

	response := controller.ShiftService.GetByNama(shift)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Update(c *fiber.Ctx) error {
	var request web.ShiftRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	shift := c.Params("shift")

	response := controller.ShiftService.Update(shift, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *shiftControllerImpl) Delete(c *fiber.Ctx) error {
	shift := c.Params("shift")

	controller.ShiftService.Delete(shift)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewShiftControllerProvider(service *usecase.ShiftService) ShiftController {
	return &shiftControllerImpl{*service}
}
