package usecase

import (
    "errors"
    "todo-api/internal/repository"
    "todo-api/pkg/jwt"
)

// AuthUseCase handles authentication logic
type AuthUseCase struct {
    UserRepo repository.UserRepository
}

// Login authenticates a user and returns a JWT token
func (uc *AuthUseCase) Login(email, password string) (string, error) {
    // Sử dụng phương thức GetByEmail từ UserRepository
    user, err := uc.UserRepo.GetByEmail(email)
    if err != nil || user.Password != password {
        return "", errors.New("invalid email or password")
    }

    // Generate JWT token
    token, err := jwt.GenerateToken(user.ID, "user") // Role là "user"
    if err != nil {
        return "", err
    }

    return token, nil
}
