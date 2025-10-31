package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// app.Static("/", "./frontend/build")
	app.Get("/", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString("OK") // []byte("user=john")
	})
	app.Post("/", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.Send(c.BodyRaw()) // []byte("user=john")
	})

	app.Get("/rpc/*", func(c *fiber.Ctx) error {
		todos := []Todo{
			{ID: 1, Title: "Buy groceries", Completed: false},
			{ID: 2, Title: "Walk the dog", Completed: true},
			{ID: 3, Title: "Finish project", Completed: false},
		}
		jsonData, err := json.Marshal(todos)
		if err != nil {
			return c.Status(500).SendString("Error encoding JSON")
		}
		c.Set("Content-Type", "application/json")
		return c.Send(jsonData)
	})
	app.Post("/rpc/*", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString(`{"status": "OK"}`) // []byte("user=john")
	})

	app.Listen(":9999")
}
