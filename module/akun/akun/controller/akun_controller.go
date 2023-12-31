package controller

import "github.com/gofiber/fiber/v2"

type AkunController interface {
	Route(app *fiber.App)
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	PegawaiGet(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	PegawaiUpdate(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
