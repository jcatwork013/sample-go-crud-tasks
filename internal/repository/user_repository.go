package repository

import (
    "todo-api/internal/entity"
    "todo-api/pkg/database"
)

// UserRepository handles database operations for users
type UserRepository struct{}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *entity.User) error {
    return database.DB.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uint) (*entity.User, error) {
    var user entity.User
    err := database.DB.First(&user, id).Error
    return &user, err
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]entity.User, error) {
    var users []entity.User
    err := database.DB.Find(&users).Error
    return users, err
}

// GetByEmail retrieves a user by their email
func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
    var user entity.User
    err := database.DB.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
