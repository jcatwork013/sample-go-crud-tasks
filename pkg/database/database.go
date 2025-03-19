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

// Connect initializes the database connection
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
}

// MigrateTables ensures the `tasks` and `users` tables exist
func MigrateTables() {
    err := DB.AutoMigrate(&entity.Task{}, &entity.User{})
    if err != nil {
        log.Fatalf("❌ Failed to migrate database: %v", err)
    }

    fmt.Println("✅ Database schema migrated successfully!")
}

// SeedData inserts default data for tasks and users
func SeedData() {
    seedTasks()
    seedUsers()
}

// DropTables drops the `tasks` and `users` tables
func DropTables() {
    err := DB.Migrator().DropTable(&entity.Task{}, &entity.User{})
    if err != nil {
        log.Fatalf("❌ Failed to drop tables: %v", err)
    }

    fmt.Println("✅ Tables dropped successfully!")
}

// seedTasks inserts default tasks if the `tasks` table is empty
func seedTasks() {
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
        fmt.Println("✅ Tasks seeded successfully!")
    } else {
        fmt.Println("✅ Tasks already exist, skipping seed.")
    }
}

// seedUsers inserts default users if the `users` table is empty
func seedUsers() {
    var count int64
    DB.Model(&entity.User{}).Count(&count)

    if count == 0 {
        fmt.Println("🌱 Seeding database with sample users...")

        users := []entity.User{
            {Username: "admin", Email: "admin@example.com", Password: "password"},
            {Username: "user1", Email: "user1@example.com", Password: "password"},
        }

        DB.Create(&users)
        fmt.Println("✅ Users seeded successfully!")
    } else {
        fmt.Println("✅ Users already exist, skipping seed.")
    }
}
