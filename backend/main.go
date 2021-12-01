package main

import (
	"log"

	"github.com/gahtoe/pfol/database"
	"github.com/gahtoe/pfol/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Set up routes
func setUpRoutes(app *fiber.App) {
	app.Post("/api/users", routes.CreateUser) // Create a user
	app.Get("/api/users", routes.GetUsers)    // Get all users
}

func main() {
	// Connect database
	database.ConnectDb()

	// Create new fiber instance
	app := fiber.New()

	// API
	setUpRoutes(app)

	// Middleware
	app.Use(cors.New()) // CORS middleware

	// Listen on port :3000
	log.Fatal(app.Listen(":3000"))

}
