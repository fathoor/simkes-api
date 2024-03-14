package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type ShiftController struct {
	ShiftUseCase *usecase.ShiftUseCase
	Log          *zerolog.Logger
}

func NewShiftController(i *do.Injector) (*ShiftController, error) {
	return &ShiftController{
		ShiftUseCase: do.MustInvoke[*usecase.ShiftUseCase](i),
		Log:          do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *ShiftController) Create(ctx *fiber.Ctx) error {
	var request model.ShiftRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.ShiftUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *ShiftController) Get(ctx *fiber.Ctx) error {
	response := c.ShiftUseCase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *ShiftController) GetByNama(ctx *fiber.Ctx) error {
	shift := ctx.Params("shift")

	response := c.ShiftUseCase.GetByNama(shift)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *ShiftController) Update(ctx *fiber.Ctx) error {
	var request model.ShiftRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	shift := ctx.Params("shift")

	response := c.ShiftUseCase.Update(shift, &request)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *ShiftController) Delete(ctx *fiber.Ctx) error {
	shift := ctx.Params("shift")

	c.ShiftUseCase.Delete(shift)

	return ctx.SendStatus(fiber.StatusNoContent)
}
