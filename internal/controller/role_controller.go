package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type RoleController struct {
	RoleUseCase *usecase.RoleUseCase
}

func NewRoleController(i *do.Injector) (*RoleController, error) {
	return &RoleController{
		RoleUseCase: do.MustInvoke[*usecase.RoleUseCase](i),
	}, nil
}

func (c *RoleController) Create(ctx *fiber.Ctx) error {
	var request web.RoleRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.RoleUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *RoleController) Get(ctx *fiber.Ctx) error {
	response := c.RoleUseCase.GetAll()

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *RoleController) GetByNama(ctx *fiber.Ctx) error {
	role := ctx.Params("role")

	response := c.RoleUseCase.GetByRole(role)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *RoleController) Update(ctx *fiber.Ctx) error {
	var request web.RoleRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	role := ctx.Params("role")

	response := c.RoleUseCase.Update(role, &request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *RoleController) Delete(ctx *fiber.Ctx) error {
	role := ctx.Params("role")

	c.RoleUseCase.Delete(role)

	return ctx.SendStatus(fiber.StatusNoContent)
}
