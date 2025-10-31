package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const BASE_PATH = "../"

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
		path := c.Params("*")
		fn := c.Query("fn")
		fullPath := filepath.Join(BASE_PATH, path)
		fmt.Println("RPC path:", path)
		fmt.Println("Function:", fn)
		fmt.Println("Full path:", fullPath)

		if !strings.HasSuffix(path, ".remote.go") {
			fmt.Println("Invalid file type, must end with .remote.go")
			return c.Status(400).SendString("Invalid file type")
		}

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			fmt.Println("File does not exist:", fullPath)
			return c.Status(404).SendString("File not found")
		}
		fmt.Println("File exists:", fullPath)

		todos := []Todo{
			{ID: 1, Title: "Buy groceries", Completed: false},
			{ID: 2, Title: "Walk the dog", Completed: true},
			{ID: 3, Title: "Finish project", Completed: false},
		}

		response := map[string]interface{}{
			"path":  path,
			"fn":    fn,
			"todos": todos,
		}

		jsonData, err := json.Marshal(response)
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
