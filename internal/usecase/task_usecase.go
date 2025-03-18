package usecase

import (
	"todo-api/internal/entity"
	"todo-api/internal/repository"
)

// TaskUseCase handles business logic for tasks
type TaskUseCase struct {
	Repo repository.TaskRepository
}

// CreateTask creates a new task
func (uc *TaskUseCase) CreateTask(title, desc string) (*entity.Task, error) {
	task := &entity.Task{Title: title, Description: desc, Completed: false}
	err := uc.Repo.Create(task)
	return task, err
}

// GetAllTasks retrieves all tasks
func (uc *TaskUseCase) GetAllTasks() ([]entity.Task, error) {
	return uc.Repo.GetAll()
}

// GetTask retrieves a single task
func (uc *TaskUseCase) GetTask(id uint) (*entity.Task, error) {
	return uc.Repo.GetByID(id)
}

// UpdateTask updates an existing task
func (uc *TaskUseCase) UpdateTask(task *entity.Task) error {
	return uc.Repo.Update(task)
}

// DeleteTask removes a task
func (uc *TaskUseCase) DeleteTask(id uint) error {
	return uc.Repo.Delete(id)
}
