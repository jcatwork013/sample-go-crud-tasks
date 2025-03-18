package main

import (
	"todo-api/internal/delivery"
	"todo-api/internal/repository"
	"todo-api/internal/usecase"
	"todo-api/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Automatically connect to DB, migrate schema, and seed data
	database.Connect()

	// Set up repository and use case layers
	taskRepo := repository.TaskRepository{}
	taskUseCase := usecase.TaskUseCase{Repo: taskRepo}
	taskHandler := delivery.TaskHandler{Usecase: taskUseCase}

	// Create the router
	r := gin.Default()

	// Define API endpoints
	r.POST("/tasks", taskHandler.CreateTask)         // Create a new task
	r.GET("/tasks", taskHandler.GetTasks)           // Get all tasks
	r.GET("/tasks/:id", taskHandler.GetTaskByID)    // Get a specific task by ID
	r.PUT("/tasks/:id", taskHandler.UpdateTask)     // Update an existing task
	r.PATCH("/tasks/:id/completed", taskHandler.MarkTaskCompleted) // Mark task as completed
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)  // Delete a task

	// Start the API server
	r.Run(":8080")
}
