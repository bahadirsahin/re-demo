package controller

import (
	"github.com/gofiber/fiber/v2"
)

// ping
func Ping(c *fiber.Ctx) error {
	return c.
		SendStatus(fiber.StatusNoContent)
}

// packs
func Packs(c *fiber.Ctx) error {
	return c.
		SendStatus(fiber.StatusNoContent)
}
