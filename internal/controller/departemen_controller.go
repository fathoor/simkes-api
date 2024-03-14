package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type DepartemenController struct {
	DepartemenUseCase *usecase.DepartemenUseCase
	Log               *zerolog.Logger
}

func NewDepartemenController(i *do.Injector) (*DepartemenController, error) {
	return &DepartemenController{
		DepartemenUseCase: do.MustInvoke[*usecase.DepartemenUseCase](i),
		Log:               do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *DepartemenController) Create(ctx *fiber.Ctx) error {
	var request model.DepartemenRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.DepartemenUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *DepartemenController) Get(ctx *fiber.Ctx) error {
	response := c.DepartemenUseCase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *DepartemenController) GetByNama(ctx *fiber.Ctx) error {
	departemen := ctx.Params("departemen")

	response := c.DepartemenUseCase.GetByDepartemen(departemen)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *DepartemenController) Update(ctx *fiber.Ctx) error {
	var request model.DepartemenRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	departemen := ctx.Params("departemen")

	response := c.DepartemenUseCase.Update(departemen, &request)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *DepartemenController) Delete(ctx *fiber.Ctx) error {
	departemen := ctx.Params("departemen")

	c.DepartemenUseCase.Delete(departemen)

	return ctx.SendStatus(fiber.StatusNoContent)
}
