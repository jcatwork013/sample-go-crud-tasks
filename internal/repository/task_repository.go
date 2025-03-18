package repository

import (
	"todo-api/internal/entity"
	"todo-api/pkg/database"
)

// TaskRepository handles database operations for tasks
type TaskRepository struct{}

// Create inserts a new task into the database
func (r *TaskRepository) Create(task *entity.Task) error {
	return database.DB.Create(task).Error
}

// GetAll retrieves all tasks
func (r *TaskRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

// GetByID retrieves a task by ID
func (r *TaskRepository) GetByID(id uint) (*entity.Task, error) {
	var task entity.Task
	err := database.DB.First(&task, id).Error
	return &task, err
}

// Update updates a task
func (r *TaskRepository) Update(task *entity.Task) error {
	return database.DB.Save(task).Error
}

// Delete removes a task from the database
func (r *TaskRepository) Delete(id uint) error {
	return database.DB.Delete(&entity.Task{}, id).Error
}
