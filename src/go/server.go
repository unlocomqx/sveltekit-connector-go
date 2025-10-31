package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Starting server...")
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

	app.All("/rpc/*", func(c *fiber.Ctx) error {
		path := c.Params("*")
		fn := c.Query("fn")

		postData := make(map[string]any)
		if c.Method() == "POST" && len(c.Body()) > 0 {
			if err := c.BodyParser(&postData); err != nil {
				log.Printf("Error parsing request body: %v", err)
				return c.Status(400).SendString(fmt.Sprintf("Error parsing request body: %v", err))
			}
		}

		log.Printf("POST data: %v", postData)

		if !strings.HasSuffix(path, ".remote.go") {
			log.Printf("Invalid file type")
			return c.Status(400).SendString("Invalid file type")
		}

		if fn == "" {
			log.Printf("Function name (fn) is required")
			return c.Status(400).SendString("Function name (fn) is required")
		}

		// Call src/routes/todos.remote.go:queryTodos, see src/routes/registry.go
		result, err := executeRemoteFunction(path, fn, postData)
		if err != nil {
			log.Printf("Error executing function: %v", err)
			return c.Status(500).SendString(fmt.Sprintf("Error executing function: %v", err))
		}

		c.Set("Content-Type", "application/json")
		return c.Send(result)
	})

	fmt.Println("Listening on :9999")
	if err := app.Listen(":9999"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
