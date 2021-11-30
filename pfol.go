package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create new fiber instance
	app := fiber.New()

	app.Use(cors.New()) // CORS middleware

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))

}
