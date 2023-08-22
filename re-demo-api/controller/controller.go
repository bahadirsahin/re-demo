package controller

import (
	"re-demo-api/service"
	"re-demo-api/structs"

	"github.com/gofiber/fiber/v2"
)

// ping
func Ping(c *fiber.Ctx) error {
	return c.
		SendStatus(fiber.StatusNoContent)
}

// packs
func Packs(c *fiber.Ctx) error {
	// validate input
	preq := structs.PackRequest{}
	err := c.BodyParser(&preq)

	// if input is invalid
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"error":   true,
					"message": "Invalid input.",
				},
			)
	}

	pres, err := service.GetPacksResponse(preq)

	// if there is error
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"error":   true,
					"message": err.Error(),
				},
			)
	}

	return c.
		Status(fiber.StatusOK).
		JSON(pres)
}
