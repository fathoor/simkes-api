package exception

import (
	"errors"
	"github.com/fathoor/simkes-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx, e error) error {
	var (
		badRequestError     *BadRequestError
		unauthorizedError   *UnauthorizedError
		forbiddenError      *ForbiddenError
		notFoundError       *NotFoundError
		internalServerError *InternalServerError
	)

	switch {
	case errors.As(e, &badRequestError):
		return c.Status(badRequestError.Code()).JSON(model.WebResponse{
			Code:   badRequestError.Code(),
			Status: badRequestError.Status(),
			Data:   badRequestError.Error(),
		})
	case errors.As(e, &unauthorizedError):
		return c.Status(unauthorizedError.Code()).JSON(model.WebResponse{
			Code:   unauthorizedError.Code(),
			Status: unauthorizedError.Status(),
			Data:   unauthorizedError.Error(),
		})
	case errors.As(e, &forbiddenError):
		return c.Status(forbiddenError.Code()).JSON(model.WebResponse{
			Code:   forbiddenError.Code(),
			Status: forbiddenError.Status(),
			Data:   forbiddenError.Error(),
		})
	case errors.As(e, &notFoundError):
		return c.Status(notFoundError.Code()).JSON(model.WebResponse{
			Code:   notFoundError.Code(),
			Status: notFoundError.Status(),
			Data:   notFoundError.Error(),
		})
	case errors.As(e, &internalServerError):
		return c.Status(internalServerError.Code()).JSON(model.WebResponse{
			Code:   internalServerError.Code(),
			Status: internalServerError.Status(),
			Data:   internalServerError.Error(),
		})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   e.Error(),
		})
	}
}
