package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type PegawaiController struct {
	PegawaiUseCase *usecase.PegawaiUseCase
	Log            *zerolog.Logger
}

func NewPegawaiController(i *do.Injector) (*PegawaiController, error) {
	return &PegawaiController{
		PegawaiUseCase: do.MustInvoke[*usecase.PegawaiUseCase](i),
		Log:            do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *PegawaiController) Create(ctx *fiber.Ctx) error {
	var request model.PegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.PegawaiUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *PegawaiController) Get(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")
	size := ctx.QueryInt("size")

	if size < 10 {
		size = 10
	}

	if page < 1 {
		response := c.PegawaiUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.PegawaiUseCase.GetPage(page, size)

		return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *PegawaiController) GetByNIP(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	response := c.PegawaiUseCase.GetByNIP(nip)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *PegawaiController) Update(ctx *fiber.Ctx) error {
	var request model.PegawaiRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := ctx.Params("nip")

	response := c.PegawaiUseCase.Update(nip, &request)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *PegawaiController) Delete(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	c.PegawaiUseCase.Delete(nip)

	return ctx.SendStatus(fiber.StatusNoContent)
}
