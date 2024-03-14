package config

import (
	"github.com/fathoor/simkes-api/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp(cfg *Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cfg.Get("APP_NAME"),
		ErrorHandler: exception.Handler,
	})

	app.Use(recover.New())
	app.Use(cors.New())

	return app
}
