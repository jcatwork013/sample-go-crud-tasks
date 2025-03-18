package delivery

import (
	"net/http"
	"strconv"
	"todo-api/internal/entity"
	"todo-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	Usecase usecase.TaskUseCase
}

// CreateTask API
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req entity.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task, err := h.Usecase.CreateTask(req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetAllTasks API
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.Usecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID API
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.Usecase.GetTask(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask API
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req entity.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	task, err := h.Usecase.GetTask(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Title = req.Title
	task.Description = req.Description
	task.Completed = req.Completed

	err = h.Usecase.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// MarkTaskCompleted API
func (h *TaskHandler) MarkTaskCompleted(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	task, err := h.Usecase.GetTask(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Completed = true
	err = h.Usecase.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark task as completed"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask API
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.Usecase.DeleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
