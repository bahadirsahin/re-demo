package api

import (
	"re-demo-api/controller"
	"re-demo-api/vars"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Init() {
	// create fiber app
	app := fiber.New()

	// set cors
	app.Use(cors.New())

	// set rate limit
	app.Use(limiter.New(
		limiter.Config{
			Max:        100,
			Expiration: 60 * time.Second,
		},
	))

	// define api base
	api := app.Group("api/v1")

	// setup routes
	api.Get("ping", controller.Ping)
	api.Post("packs", verifyToken, controller.Packs)

	// start api
	err := app.Listen(":80")

	if err != nil {
		panic(err)
	}
}

// verify api token
func verifyToken(c *fiber.Ctx) error {
	// get token
	token := c.Get("X-Token")

	// if not verified, return error
	if !strings.EqualFold(token, vars.Env.ApiToken) {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(
				&fiber.Map{
					"error":   true,
					"message": "Unauthorized.",
				},
			)
	}

	// if verified, continue
	return c.Next()
}
