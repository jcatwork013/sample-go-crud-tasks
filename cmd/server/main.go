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
    taskHandler, userHandler, authHandler := initializeHandlers()

    // Thiết lập router
    r := setupRouter(taskHandler, userHandler, authHandler)

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

// initializeHandlers khởi tạo các handler cho Task, User, và Auth
func initializeHandlers() (delivery.TaskHandler, delivery.UserHandler, delivery.AuthHandler) {
    // Khởi tạo repository
    taskRepo := repository.TaskRepository{}
    userRepo := repository.UserRepository{}

    // Khởi tạo usecase
    taskUseCase := usecase.TaskUseCase{Repo: taskRepo}
    userUseCase := usecase.UserUseCase{Repo: userRepo}
    authUseCase := usecase.AuthUseCase{UserRepo: userRepo}

    // Khởi tạo handler
    taskHandler := delivery.TaskHandler{Usecase: taskUseCase}
    userHandler := delivery.UserHandler{Usecase: userUseCase}
    authHandler := delivery.AuthHandler{Usecase: authUseCase}

    return taskHandler, userHandler, authHandler
}

// setupRouter thiết lập router và định nghĩa các endpoint
func setupRouter(taskHandler delivery.TaskHandler, userHandler delivery.UserHandler, authHandler delivery.AuthHandler) *gin.Engine {
    r := gin.Default()
    r.Use(gin.Logger(), gin.Recovery())

    // Các API không yêu cầu xác thực
    r.POST("/login", authHandler.Login)

    // Middleware xác thực JWT
    authMiddleware := delivery.AuthMiddleware()

    // Các API yêu cầu xác thực
    auth := r.Group("/")
    auth.Use(authMiddleware)
    {
        // Task routes
        auth.POST("/tasks", taskHandler.CreateTask)
        auth.GET("/tasks", taskHandler.GetTasks)
        auth.GET("/tasks/:id", taskHandler.GetTaskByID)
        auth.PUT("/tasks/:id", taskHandler.UpdateTask)
        auth.PATCH("/tasks/:id/completed", taskHandler.MarkTaskCompleted)
        auth.DELETE("/tasks/:id", taskHandler.DeleteTask)

        // User routes
        auth.POST("/users", userHandler.CreateUser)
        auth.GET("/users", userHandler.GetAllUsers)
        auth.GET("/users/:id", userHandler.GetUserByID)
    }

    return r
}
