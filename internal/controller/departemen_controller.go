package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type DepartemenController struct {
	DepartemenUseCase *usecase.DepartemenUseCase
}

func NewDepartemenController(i *do.Injector) (*DepartemenController, error) {
	return &DepartemenController{
		DepartemenUseCase: do.MustInvoke[*usecase.DepartemenUseCase](i),
	}, nil
}

func (c *DepartemenController) Create(ctx *fiber.Ctx) error {
	var request web.DepartemenRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.DepartemenUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *DepartemenController) Get(ctx *fiber.Ctx) error {
	response := c.DepartemenUseCase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *DepartemenController) GetByNama(ctx *fiber.Ctx) error {
	departemen := ctx.Params("departemen")

	response := c.DepartemenUseCase.GetByDepartemen(departemen)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *DepartemenController) Update(ctx *fiber.Ctx) error {
	var request web.DepartemenRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	departemen := ctx.Params("departemen")

	response := c.DepartemenUseCase.Update(departemen, &request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
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
