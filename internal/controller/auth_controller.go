package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type AuthController struct {
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthController(i *do.Injector) (*AuthController, error) {
	return &AuthController{
		AuthUseCase: do.MustInvoke[*usecase.AuthUseCase](i),
	}, nil
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var request web.AuthRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.AuthUseCase.Login(&request)

	return ctx.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
