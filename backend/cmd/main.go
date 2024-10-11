package main

import (
	"backend_fullstack/internal/adapters"
	"backend_fullstack/internal/adapters/auth"
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/core/services"
	"backend_fullstack/internal/database"
	"backend_fullstack/internal/storage"
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

	// Initialize MinIO storage
	// minioClient, err := storage.MinioConnection()
	minioClient, err := storage.MinioConnection()
	if err != nil {
		log.Fatalf("Failed to connect to MinIO: %v", err)
	}

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

	minioRepo := adapters.NewMinIOFileRepository(minioClient)
	minioService := services.NewMinIOServiceService(minioRepo)
	minioHandler := adapters.NewHttpMinIOHandler(minioService)

	instructorRepo := adapters.NewGormInstructorRepository(db)
	instructorService := services.NewInstructorService(instructorRepo, minioRepo)
	instructorHandler := adapters.NewHttpInstructorHandler(instructorService, userService)

	studentRepo := adapters.NewGormStudentRepository(db)
	studentService := services.NewStudentService(studentRepo, minioRepo)
	studentHandler := adapters.NewHttpStudentHandler(studentService)

	authGroup := app.Group("/auth")
	googleGroup := authGroup.Group("/google")
	googleGroup.Get("/login", googleOAuthHandler.GetGoogleLoginURL)
	googleGroup.Get("/callback", googleOAuthHandler.GetGoogleCallback)

	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

	// use middleware to verify token
	apiGroup := app.Group("/api", authHandler.VerifyToken)
	apiGroup.Get("/QueryUsers", userHandler.GetUsers)
	apiGroup.Post("/CreateCourse", instructorHandler.CreateCourse)
	apiGroup.Get("/QueryCourseByUserID", instructorHandler.GetCourseByUserID)
	apiGroup.Get("/QueryNameByUserID", instructorHandler.GetNameByUserID)
	apiGroup.Get("/QueryUserGroupByUserID", instructorHandler.GetUserGroupByUserID)
	apiGroup.Get("/QueryAssignmentByUserID", instructorHandler.GetAssignmentByUserID)
	apiGroup.Post("/CreateAssignment", instructorHandler.CreateAssignment)
	apiGroup.Get("/QueryAssignmentsByCourseID", instructorHandler.GetAssignmentsByCourseID)
	apiGroup.Get("/QueryAssignmentsByUserIDSorted", instructorHandler.GetAssignmentByUserIDSorted)
	apiGroup.Get("/QueryAssignmentsByCourseIDAndAssignmentID", instructorHandler.GetAssignmentByCourseIDAndAssignmentID)
	apiGroup.Put("/UpdateAssignmentByCourseIDAndAssignmentID", instructorHandler.UpdateAssignmentByCourseIDAndAssignmentID)
	apiGroup.Get("/QueryUsersEnrollment", instructorHandler.GetUsersEnrollment)
	apiGroup.Post("/CreateEnrollment", instructorHandler.CreateEnrollment)
	apiGroup.Delete("/DeleteUserEnrollment", instructorHandler.DeleteUserEnrollment)
	apiGroup.Delete("/DeleteCourse", instructorHandler.DeleteCourse)
	apiGroup.Delete("/DeleteAssignmentByCourseIDAndAssignmentID", instructorHandler.DeleteAssignmentByCourseIDAndAssignmentID)
	apiGroup.Get("/QuerySubmissionsByCourseIDAndAssignmentID", instructorHandler.GetSubmissionsByCourseIDAndAssignmentID)
	apiGroup.Post("/UploadFiles", minioHandler.CreateFileToMinIO)
	apiGroup.Post("UploadAssignmentFiles", instructorHandler.UploadAssignmentFile)

	apiGroup.Get("QueryCourseByUserIDStd", studentHandler.GetCourseByUserIDStd)
	apiGroup.Get("QueryAssignmentByUserIDStd", studentHandler.GetAssignmentByUserIDStd)
	apiGroup.Get("QueryAssignmentByUserIDSortedStd", studentHandler.GetAssignmentByUserIDSortedStd)
	apiGroup.Get("QueryUpcomingAssignmentsStd", studentHandler.GetUpcomingAssignments)
	apiGroup.Post("CreateSubmission", studentHandler.UploadAssignmentFile)
	apiGroup.Get("QuerySubmissionsStatus", studentHandler.GetSubmissionsStatus)

	app.Post("/CreateUser", userHandler.CreateUser)
	app.Get("/QueryUserById", userHandler.GetUserByID)
	app.Get("/QueryUserByUserName", userHandler.GetUserByUserName)
	app.Get("/QueryUserIDByEmail", userHandler.GetUserIDByEmail)
	//app.Get("/QueryUsers", userHandler.GetUsers)
	app.Put("/UpdateUser", userHandler.UpdateUser)
	app.Delete("/DeleteUser", userHandler.DeleteUser)

	app.Post("/CreateUserGroup", adminHandler.CreateUserGroup)
	app.Get("/QueryUserGroupById", adminHandler.GetUserGroupByID)
	app.Get("/QueryUserGroups", adminHandler.GetUserGroups)
	app.Put("/UpdateUserGroup", adminHandler.UpdateUserGroup)
	app.Delete("/DeleteUserGroup", adminHandler.DeleteUserGroup)

	//app.Post("/CreateCourse", instructorHandler.CreateCourse)
	app.Get("/QueryCourseById", instructorHandler.GetCourseByID)
	app.Get("/QueryCourses", instructorHandler.GetCourses)
	app.Put("/UpdateCourse", instructorHandler.UpdateCourse)
	//app.Delete("/DeleteCourse", instructorHandler.DeleteCourse)

	//app.Post("/CreateAssignment", instructorHandler.CreateAssignment)
	app.Get("/QueryAssignmentByAssignmentID", instructorHandler.GetAssignmentByAssignmentID)
	app.Get("/QueryAssignments", instructorHandler.GetAssignments)
	//app.Get("/QueryAssignmentsByCourseID", instructorHandler.GetAssignmentsByCourseID)
	app.Put("/UpdateAssignment", instructorHandler.UpdateAssignment)
	app.Delete("/DeleteAssignment", instructorHandler.DeleteAssignment)

	app.Post("/CreateInstructorList", instructorHandler.CreateInstructorList)
	app.Get("/QueryInstructorList", instructorHandler.GetInstructorsList)
	app.Get("/QueryInstructorListByCourseId", instructorHandler.GetInstructorsListByCourseID)
	app.Get("/QueryInstructorListByListId", instructorHandler.GetInstructorsListByListID)
	app.Delete("/DeleteInstructorList", instructorHandler.DeleteInstructorList)

	app.Get("/QueryPersonDataByUserID", instructorHandler.GetPersonDataByUserID)
	app.Get("/QueryCourseByUserID", instructorHandler.GetCourseByUserID)

	//app.Post("/CreateEnrollment", instructorHandler.CreateEnrollment)
	app.Get("/QueryEnrollments", instructorHandler.GetEnrollments)
	app.Get("/QueryEnrollmentsByCourseID", instructorHandler.GetEnrollmentsByCourseID)
	app.Delete("/DeleteEnrollment", instructorHandler.DeleteEnrollment)

	port := os.Getenv("PORT")
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Println("Server is running on port: ", port)
}

//// First: Add api Query assignment by course_id and assignment_id
//// Then use details to INS_Assignment
//// Second: Add api update assignment
//// Third: Add api delete assignment
//// Fourth: Add api Query Submissions by course_id and assignment_id
//// Fifth: Add api Upload Files to Upload on assignment
