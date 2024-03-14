package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type kehadiranControllerImpl struct {
	usecase.KehadiranService
}

func (controller *kehadiranControllerImpl) CheckIn(c *fiber.Ctx) error {
	var request web.KehadiranRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.KehadiranService.CheckIn(&request)

	return c.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (controller *kehadiranControllerImpl) CheckOut(c *fiber.Ctx) error {
	var request web.KehadiranRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.KehadiranService.CheckOut(&request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *kehadiranControllerImpl) Get(c *fiber.Ctx) error {
	nip := c.Query("nip")

	if nip != "" {
		response := controller.KehadiranService.GetByNIP(nip)

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := controller.KehadiranService.GetAll()

		return c.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (controller *kehadiranControllerImpl) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	response := controller.KehadiranService.GetByID(id)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *kehadiranControllerImpl) Update(c *fiber.Ctx) error {
	var request web.KehadiranUpdateRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	id := c.Params("id")

	response := controller.KehadiranService.Update(id, &request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *kehadiranControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.KehadiranService.Delete(id)

	return c.SendStatus(fiber.StatusNoContent)
}

func NewKehadiranControllerProvider(service *usecase.KehadiranService) KehadiranController {
	return &kehadiranControllerImpl{*service}
}
