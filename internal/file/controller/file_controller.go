package controller

import "github.com/gofiber/fiber/v2"

type FileController interface {
	Upload(c *fiber.Ctx) error
	Download(c *fiber.Ctx) error
	View(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
