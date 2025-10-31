package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const BASE_PATH = "../"

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

	app.All("/rpc/*", func(c *fiber.Ctx) error {
		path := c.Params("*")
		fn := c.Query("fn")
		fullPath := filepath.Join(BASE_PATH, path)

		var postData []byte
		if c.Method() == "POST" {
			postData = c.Body()
		}

		if !strings.HasSuffix(path, ".remote.go") {
			return c.Status(400).SendString("Invalid file type")
		}

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return c.Status(404).SendString("File not found")
		}

		if fn == "" {
			return c.Status(400).SendString("Function name (fn) is required")
		}

		result, err := executeRemoteFunction(fullPath, fn, postData)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error executing function: %v", err))
		}

		c.Set("Content-Type", "application/json")
		return c.Send(result)
	})

	app.Listen(":9999")
}
