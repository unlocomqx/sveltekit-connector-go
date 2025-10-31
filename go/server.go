package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./frontend/build")
	app.Get("/", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString("OK2") // []byte("user=john")
	})
	app.Post("/", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.Send(c.BodyRaw()) // []byte("user=john")
	})

	app.Get("/rpc", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString("OK") // []byte("user=john")
	})
	app.Post("/rpc", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString("OK") // []byte("user=john")
	})

	app.Listen(":9999")
}
