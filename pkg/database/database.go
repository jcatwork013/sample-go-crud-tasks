package database

import (
	"fmt"
	"log"
	"todo-api/config"
	"todo-api/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global database instance
var DB *gorm.DB

// Connect initializes the database connection and ensures tables are created
func Connect() {
	// Get database connection string from environment variables
	dsn := config.GetDatabaseURI()

	// Open database connection
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Database connected successfully!")

	// Ensure the `tasks` table exists
	err = DB.AutoMigrate(&entity.Task{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	fmt.Println("✅ Database schema migrated successfully!")

	// Seed database with initial data
	SeedDatabase()
}

// SeedDatabase inserts default tasks if the table is empty
func SeedDatabase() {
	var count int64
	DB.Model(&entity.Task{}).Count(&count)

	if count == 0 {
		fmt.Println("🌱 Seeding database with sample tasks...")

		tasks := []entity.Task{
			{Title: "Learn Golang", Description: "Build a REST API", Completed: false},
			{Title: "Read Clean Code", Description: "Read the book by Robert C. Martin", Completed: false},
			{Title: "Exercise", Description: "Go for a run and stay healthy", Completed: false},
		}

		DB.Create(&tasks)
		fmt.Println("✅ Database seeded successfully!")
	} else {
		fmt.Println("✅ Database already has tasks, skipping seed.")
	}
}
