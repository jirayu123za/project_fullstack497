package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Init fiber server
	app := fiber.New()

	dsn := os.Getenv("DATABASE_DSN")
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	_, err = ConnectPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

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

func ConnectPostgres(dsn string) (*gorm.DB, error) {
	// Config gorm logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database successfully")

	// Migration

	return db, err
}
