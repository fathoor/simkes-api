package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
)

type AkunController struct {
	AkunUseCase *usecase.AkunUseCase
}

func NewAkunController(i *do.Injector) (*AkunController, error) {
	return &AkunController{
		AkunUseCase: do.MustInvoke[*usecase.AkunUseCase](i),
	}, nil
}

func (c *AkunController) Create(ctx *fiber.Ctx) error {
	var request web.AkunRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.AkunUseCase.Create(&request)

	return ctx.Status(fiber.StatusCreated).JSON(web.Response{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   response,
	})
}

func (c *AkunController) Get(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page")
	size := ctx.QueryInt("size")

	if size < 10 {
		size = 10
	}

	if page < 1 {
		response := c.AkunUseCase.GetAll()

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.AkunUseCase.GetPage(page, size)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *AkunController) GetByNIP(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	response := c.AkunUseCase.GetByNIP(nip)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (c *AkunController) Update(ctx *fiber.Ctx) error {
	var request web.AkunRequest

	claims := ctx.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	nip := ctx.Params("nip")

	if role == "Admin" {
		response := c.AkunUseCase.UpdateAdmin(nip, &request)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	} else {
		response := c.AkunUseCase.Update(nip, &request)

		return ctx.Status(fiber.StatusOK).JSON(web.Response{
			Code:   fiber.StatusOK,
			Status: "OK",
			Data:   response,
		})
	}
}

func (c *AkunController) Delete(ctx *fiber.Ctx) error {
	nip := ctx.Params("nip")

	c.AkunUseCase.Delete(nip)

	return ctx.SendStatus(fiber.StatusNoContent)
}
