package main

import (
	"backend_fullstack/internal/adapters"
	"backend_fullstack/internal/adapters/auth"
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/database"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load .env file
	config.LoadEnv()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in the environment")
	}

	// Init fiber server
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] FROM IP: ${ip}  STATUS: ${status} LATENCY: ${latency} METHOD: ${method} PATH: ${path}\n",
	}))

	db := database.ConnectPostgres(true)

	// Init GoogleOAut configured
	auth.InitializeGoogleOAuth()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server is running",
			"error":   nil,
		})
	})

	googleOAuthRepo := adapters.NewOAuthRepository()
	googleOAuthService := services.NewOAuthService(googleOAuthRepo)
	googleOAuthHandler := adapters.NewHttpOAuthHandler(googleOAuthService)

	userRepo := adapters.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := adapters.NewHttpUserHandler(userService)

	authRepo := adapters.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo, userRepo, jwtSecret)
	authHandler := adapters.NewHttpAuthHandler(authService)

	adminRepo := adapters.NewGormAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	adminHandler := adapters.NewHttpAdminHandler(adminService)

	authGroup := app.Group("/auth")
	googleGroup := authGroup.Group("/google")
	googleGroup.Get("/login", googleOAuthHandler.GetGoogleLoginURL)
	googleGroup.Get("/callback", googleOAuthHandler.GetGoogleCallback)

	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

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
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Println("Server is running on port: ", port)
}
