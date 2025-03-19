package repository

import (
    "todo-api/internal/entity"
    "todo-api/pkg/database"
)

// AuthRepository handles authentication-related database operations
type AuthRepository struct{}

// GetByEmail retrieves a user by their email
func (r *AuthRepository) GetByEmail(email string) (*entity.User, error) {
    var user entity.User
    err := database.DB.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
