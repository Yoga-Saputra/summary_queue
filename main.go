package main

import (
	"summary/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	summary := app.Group("summary")
	router.Route(summary)

	app.Listen(":8090")
}
