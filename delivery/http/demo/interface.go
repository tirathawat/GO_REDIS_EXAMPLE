package demo_handler

import "github.com/gofiber/fiber/v2"

type IDemoHandler interface {
	GetDemo(c *fiber.Ctx) error
}
