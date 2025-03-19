package usecase

import (
	"todo-api/internal/entity"
	"todo-api/internal/repository"
)

// UserUseCase handles business logic for users
type UserUseCase struct {
	Repo repository.UserRepository
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(username, email, password string) (*entity.User, error) {
	user := &entity.User{Username: username, Email: email, Password: password}
	err := uc.Repo.Create(user)
	return user, err
}

// GetAllUsers retrieves all users
func (uc *UserUseCase) GetAllUsers() ([]entity.User, error) {
	return uc.Repo.GetAll()
}

// GetUser retrieves a single user by ID
func (uc *UserUseCase) GetUser(id uint) (*entity.User, error) {
	return uc.Repo.GetByID(id)
}
