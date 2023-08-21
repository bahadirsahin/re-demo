package api

import (
	"re-demo-api/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init() {
	// create fiber app
	app := fiber.New()

	// set cors
	app.Use(cors.New())

	// define api base
	api := app.Group("api/v1")

	// setup routes
	api.Get("ping", controller.Ping)

	// start api
	err := app.Listen(":80")

	if err != nil {
		panic(err)
	}
}
