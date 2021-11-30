package main

import (
	"github.com/alitemel89/gofiber-restapi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/alitemel89/gofiber-restapi/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
