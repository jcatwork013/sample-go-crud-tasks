package main

import (
    "log"
    "os"
    "todo-api/internal/delivery"
    "todo-api/internal/repository"
    "todo-api/internal/usecase"
    "todo-api/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    // Kiểm tra nếu có lệnh thủ công được truyền vào
    if len(os.Args) > 1 {
        command := os.Args[1]
        handleCommand(command)
        return
    }

    // Kết nối cơ sở dữ liệu
    database.Connect()

    // Khởi tạo các phụ thuộc
    taskHandler, userHandler := initializeHandlers()

    // Thiết lập router
    r := setupRouter(taskHandler, userHandler)

    // Chạy server
    log.Println("🚀 Starting server on :8080...")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("❌ Failed to start server: %v", err)
    }
}

// handleCommand xử lý các lệnh thủ công như setup, seed, drop
func handleCommand(command string) {
    database.Connect()

    switch command {
    case "setup":
        log.Println("🔧 Running `setup`...")
        database.MigrateTables()
        log.Println("✅ Setup completed! (Tables created, no data seeded)")

    case "seed":
        log.Println("🌱 Running `seed`...")
        database.SeedData()
        log.Println("✅ Seeding completed!")

    case "drop":
        log.Println("⚠️ Running `drop`...")
        database.DropTables()
        log.Println("✅ Tables dropped successfully!")

    default:
        log.Fatalf("❌ Invalid command. Use: setup | seed | drop")
    }
}

// initializeHandlers khởi tạo các handler cho Task và User
func initializeHandlers() (delivery.TaskHandler, delivery.UserHandler) {
    taskRepo := repository.TaskRepository{}
    taskUseCase := usecase.TaskUseCase{Repo: taskRepo}
    taskHandler := delivery.TaskHandler{Usecase: taskUseCase}

    userRepo := repository.UserRepository{}
    userUseCase := usecase.UserUseCase{Repo: userRepo}
    userHandler := delivery.UserHandler{Usecase: userUseCase}

    return taskHandler, userHandler
}

// setupRouter thiết lập router và định nghĩa các endpoint
func setupRouter(taskHandler delivery.TaskHandler, userHandler delivery.UserHandler) *gin.Engine {
    r := gin.Default()
    r.Use(gin.Logger(), gin.Recovery())

    // Định nghĩa các endpoint cho module Task
    r.POST("/tasks", taskHandler.CreateTask)
    r.GET("/tasks", taskHandler.GetTasks)
    r.GET("/tasks/:id", taskHandler.GetTaskByID)
    r.PUT("/tasks/:id", taskHandler.UpdateTask)
    r.PATCH("/tasks/:id/completed", taskHandler.MarkTaskCompleted)
    r.DELETE("/tasks/:id", taskHandler.DeleteTask)

    // Định nghĩa các endpoint cho module User
    r.POST("/users", userHandler.CreateUser)
    r.GET("/users", userHandler.GetAllUsers)
    r.GET("/users/:id", userHandler.GetUserByID)

    return r
}
