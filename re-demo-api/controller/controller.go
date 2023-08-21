package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.
		SendStatus(fiber.StatusNoContent)
}
