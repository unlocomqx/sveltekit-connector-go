package main

import (
	"fmt"
	"os"
	"os/exec"
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

func executeRemoteFunction(filePath string, functionName string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(fileContent)
	contentStr = strings.Replace(contentStr, "package main", "", 1)

	wrapperCode := fmt.Sprintf(`package main

import (
	"encoding/json"
	"fmt"
)

%s

func main() {
	result := %s()
	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}
`, contentStr, functionName)

	wrapperPath := filepath.Join("tmp", "wrapper.go")
	err = os.WriteFile(wrapperPath, []byte(wrapperCode), 0644)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "run", wrapperPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("execution error: %s", string(output))
	}

	return []byte(strings.TrimSpace(string(output))), nil
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

		if fn == "" {
			return c.Status(400).SendString("Function name (fn) is required")
		}

		result, err := executeRemoteFunction(fullPath, fn)
		if err != nil {
			fmt.Println("Error executing function:", err)
			return c.Status(500).SendString(fmt.Sprintf("Error executing function: %v", err))
		}

		c.Set("Content-Type", "application/json")
		return c.Send(result)
	})
	app.Post("/rpc/*", func(c *fiber.Ctx) error {
		// Get raw body from POST request:
		return c.SendString(`{"status": "OK"}`) // []byte("user=john")
	})

	app.Listen(":9999")
}
