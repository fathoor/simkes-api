package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/samber/do"
)

type AuthController struct {
	AuthUseCase *usecase.AuthUseCase
	Log         *zerolog.Logger
}

func NewAuthController(i *do.Injector) (*AuthController, error) {
	return &AuthController{
		AuthUseCase: do.MustInvoke[*usecase.AuthUseCase](i),
		Log:         do.MustInvoke[*zerolog.Logger](i),
	}, nil
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var request model.AuthRequest

	if parse := ctx.BodyParser(&request); parse != nil {
		c.Log.Error().Err(parse).Msg("Invalid request body")
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := c.AuthUseCase.Login(&request)

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
