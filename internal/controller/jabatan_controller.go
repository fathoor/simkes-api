package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type JabatanController struct {
	JabatanUseCase *usecase.JabatanUseCase
}

func NewJabatanController(i *do.Injector) (*JabatanController, error) {
	return &JabatanController{
		JabatanUseCase: do.MustInvoke[*usecase.JabatanUseCase](i),
	}, nil
}

func (c *JabatanController) Create(ctx *fiber.Ctx) error {
	var request web.JabatanRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.JabatanUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *JabatanController) Get(ctx *fiber.Ctx) error {
	response := c.JabatanUseCase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JabatanController) GetByNama(ctx *fiber.Ctx) error {
	jabatan := ctx.Params("jabatan")

	response := c.JabatanUseCase.GetByJabatan(jabatan)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JabatanController) Update(ctx *fiber.Ctx) error {
	var request web.JabatanRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	jabatan := ctx.Params("jabatan")

	response := c.JabatanUseCase.Update(jabatan, &request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *JabatanController) Delete(ctx *fiber.Ctx) error {
	jabatan := ctx.Params("jabatan")

	c.JabatanUseCase.Delete(jabatan)

	return ctx.SendStatus(fiber.StatusNoContent)
}
