package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Simple hello world test
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Add healthcheck middleware for /livez and /readyz
	app.Use(healthcheck.New(healthcheck.Config{}))

	if err := app.Listen(":4000"); err != nil {
		log.Fatal(err)
	}
}
