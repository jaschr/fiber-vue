package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/helmet/v2"
)

func main() {
	app := fiber.New()

	fc := favicon.Config{
		File: "./public/assets/favicon.ico",
	}

	app.Use(helmet.New())

	app.Use(favicon.New(fc))

	app.Static("/", "./public")

	app.Listen(":3000")

}
