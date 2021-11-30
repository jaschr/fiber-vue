package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Create new fiber instance
	app := fiber.New()

	// Favicon config -- file location of the 'favicon.ico'
	fc := favicon.Config{
		File: "./favicon.ico",
	}

	app.Use(cors.New())      // CORS middleware
	app.Use(favicon.New(fc)) // Favicon middleware
	app.Use(helmet.New())    // Helmet middleware
	app.Use(logger.New())    // Logger middleware

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))

}
