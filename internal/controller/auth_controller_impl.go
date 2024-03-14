package controller

import (
	"github.com/fathoor/simkes-api/internal/exception"
	web "github.com/fathoor/simkes-api/internal/model"
	"github.com/fathoor/simkes-api/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type authControllerImpl struct {
	usecase.AuthService
}

func (controller *authControllerImpl) Login(c *fiber.Ctx) error {
	var request web.AuthRequest

	if parse := c.BodyParser(&request); parse != nil {
		panic(&exception.BadRequestError{
			Message: "Invalid request body",
		})
	}

	response := controller.AuthService.Login(&request)

	return c.Status(fiber.StatusOK).JSON(web.Response{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func NewAuthControllerProvider(service *usecase.AuthService) AuthController {
	return &authControllerImpl{*service}
}
