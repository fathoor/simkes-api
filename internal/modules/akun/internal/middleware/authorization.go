package middleware

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/gofiber/fiber/v2"
)

func AuthorizeUserAkun(id string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(string)
		role := ctx.Locals("role").(int)

		if role == 1337 || role == 1 {
			return ctx.Next()
		} else if user == id {
			return ctx.Next()
		} else {
			panic(&exception.ForbiddenError{
				Message: "You are not allowed to access this akun",
			})
		}
	}
}

func AuthorizeUserAlamat(id string) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(string)
		role := ctx.Locals("role").(int)

		if role == 1337 || role == 1 {
			return ctx.Next()
		} else if user == id {
			return ctx.Next()
		} else {
			panic(&exception.ForbiddenError{
				Message: "You are not allowed to access this alamat",
			})
		}
	}
}
