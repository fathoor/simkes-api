package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type KehadiranController struct {
	KehadiranUseCase *usecase.KehadiranUseCase
}

func NewKehadiranController(i *do.Injector) (*KehadiranController, error) {
	return &KehadiranController{
		KehadiranUseCase: do.MustInvoke[*usecase.KehadiranUseCase](i),
	}, nil
}

func (c *KehadiranController) CheckIn(ctx *fiber.Ctx) error {
	var request web.KehadiranRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.KehadiranUseCase.CheckIn(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *KehadiranController) CheckOut(ctx *fiber.Ctx) error {
	var request web.KehadiranRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.KehadiranUseCase.CheckOut(&request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *KehadiranController) Get(ctx *fiber.Ctx) error {
	nip := ctx.Query("nip")

	if nip != "" {
		response := c.KehadiranUseCase.GetByNIP(nip)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.KehadiranUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *KehadiranController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.KehadiranUseCase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *KehadiranController) Update(ctx *fiber.Ctx) error {
	var request web.KehadiranUpdateRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	id := ctx.Params("id")

	response := c.KehadiranUseCase.Update(id, &request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *KehadiranController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.KehadiranUseCase.Delete(id)

	return ctx.SendStatus(fiber.StatusNoContent)
}
