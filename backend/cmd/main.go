package main

import (
	"backend_fullstack/internal/adapters"
	"backend_fullstack/internal/adapters/auth"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/models"
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

	db, err := ConnectPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Init GoogleOAut configured
	auth.InitializeGoogleOAuth()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server is running",
			"error":   err,
		})
	})

	googleOAuthRepo := adapters.NewOAuthRepository()
	googleOAuthService := services.NewOAuthService(googleOAuthRepo)
	googleOAuthHandler := adapters.NewHttpOAuthHandler(googleOAuthService)

	userRepo := adapters.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := adapters.NewHttpUserHandler(userService)

	adminRepo := adapters.NewGormAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	adminHandler := adapters.NewHttpAdminHandler(adminService)

	authGroup := app.Group("/auth")
	googleGroup := authGroup.Group("/google")
	googleGroup.Get("/login", googleOAuthHandler.GetGoogleLoginURL)
	googleGroup.Get("/callback", googleOAuthHandler.GetGoogleCallback)

	app.Post("/CreateUser", userHandler.CreateUser)
	app.Get("/QueryUserById", userHandler.GetUserByID)
	app.Get("/QueryUserByUserName", userHandler.GetUserByUserName)
	app.Get("/QueryUsers", userHandler.GetUsers)
	app.Put("/UpdateUser", userHandler.UpdateUser)
	app.Delete("/DeleteUser", userHandler.DeleteUser)

	app.Post("/CreateUserGroup", adminHandler.CreateUserGroup)
	app.Get("/QueryUserGroupById", adminHandler.GetUserGroupByID)
	app.Get("/QueryUserGroups", adminHandler.GetUserGroups)
	app.Put("/UpdateUserGroup", adminHandler.UpdateUserGroup)
	app.Delete("/DeleteUserGroup", adminHandler.DeleteUserGroup)

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
	err = db.AutoMigrate(
		&models.UserGroup{},
		&models.User{},
		&models.Course{},
		&models.Assignment{},
		&models.AssignmentFile{},
		&models.Enrollment{},
		&models.InstructorList{},
		&models.Submission{},
		&models.Upload{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	fmt.Println("Database migration completed!")

	return db, err
}
