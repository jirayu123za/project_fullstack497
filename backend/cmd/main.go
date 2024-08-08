package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Init fiber server
	app := fiber.New()

	// Init routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server is running",
			"error":   err,
		})
	})

	port := os.Getenv("PORT")
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	fmt.Println("Server is running on port: ", port)
}
