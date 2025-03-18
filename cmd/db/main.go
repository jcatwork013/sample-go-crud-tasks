package main

import (
	"fmt"
	"log"
	"os"
	"todo-api/internal/entity"
	"todo-api/pkg/database"

	"gorm.io/gorm"
)

// SeedDatabase inserts sample data into the database
func SeedDatabase(db *gorm.DB) {
	var count int64
	db.Model(&entity.Task{}).Count(&count)

	if count == 0 {
		fmt.Println("🌱 Seeding database with sample tasks...")

		tasks := []entity.Task{
			{Title: "Learn Golang", Description: "Study Go and build a REST API", Completed: false},
			{Title: "Read Clean Code", Description: "Read the book by Robert C. Martin", Completed: false},
			{Title: "Exercise", Description: "Go for a run and stay healthy", Completed: false},
		}

		// Insert tasks
		db.Create(&tasks)
		fmt.Println("✅ Database seeded successfully!")
	} else {
		fmt.Println("✅ Database already has tasks, skipping seed.")
	}
}

func main() {
	// Connect to the database
	database.Connect()
	db := database.DB

	// Read the command argument
	if len(os.Args) < 2 {
		fmt.Println("❌ Please provide a database command: db:setup | db:migrate | db:seed")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "db:setup":
		fmt.Println("🔧 Running `db:setup`...")
		db.AutoMigrate(&entity.Task{})
		SeedDatabase(db)
		fmt.Println("✅ `db:setup` completed!")

	case "db:migrate":
		fmt.Println("🚀 Running `db:migrate`...")
		db.AutoMigrate(&entity.Task{})
		fmt.Println("✅ `db:migrate` completed!")

	case "db:seed":
		fmt.Println("🌱 Running `db:seed`...")
		SeedDatabase(db)
		fmt.Println("✅ `db:seed` completed!")

	default:
		fmt.Println("❌ Invalid command. Use: db:setup | db:migrate | db:seed")
		os.Exit(1)
	}
}
