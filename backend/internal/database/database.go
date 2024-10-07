package database

import (
	"backend_fullstack/internal/config"
	"backend_fullstack/internal/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgres(migrate bool) *gorm.DB {
	config.LoadEnv()
	dsn := os.Getenv("DATABASE_DSN")

	// Config log
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database successfully")

	// Migration
	if migrate {
		/*
			db.Migrator().DropTable(
				&models.UserGroup{},
				&models.User{},
				&models.Course{},
				&models.Assignment{},
				&models.AssignmentFile{},
				&models.Enrollment{},
				&models.InstructorList{},
				&models.Submission{},
				&models.Upload{})
		*/
		/*
			db.Migrator().DropTable(
				&models.AssignmentFile{},
				&models.Upload{},
			)
		*/
		err := db.AutoMigrate(
			&models.UserGroup{},
			&models.User{},
			&models.Course{},
			&models.Assignment{},
			&models.AssignmentFile{},
			&models.Enrollment{},
			&models.InstructorList{},
			&models.Submission{},
			&models.Upload{})

		if err != nil {
			log.Fatal("Failed to migrate database: ", err)
		}
		fmt.Println("Database migration completed!")
	}
	return db
}
