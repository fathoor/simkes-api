package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

type CutiController struct {
	CutiUseCase *usecase.CutiUseCase
}

func NewCutiController(i *do.Injector) (*CutiController, error) {
	return &CutiController{
		CutiUseCase: do.MustInvoke[*usecase.CutiUseCase](i),
	}, nil
}

func (c *CutiController) Create(ctx *fiber.Ctx) error {
	var request web.CutiCreateRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.CutiUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *CutiController) Get(ctx *fiber.Ctx) error {
	nip := ctx.Query("nip")

	if nip != "" {
		response := c.CutiUseCase.GetByNIP(nip)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.CutiUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *CutiController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	response := c.CutiUseCase.GetByID(id)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *CutiController) Update(ctx *fiber.Ctx) error {
	var request web.CutiUpdateRequest

	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	id := ctx.Params("id")

	if role == "Admin" {
		response := c.CutiUseCase.UpdateStatus(id, &request)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.CutiUseCase.Update(id, &request)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *CutiController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.CutiUseCase.Delete(id)

	return ctx.SendStatus(fiber.StatusNoContent)
}
