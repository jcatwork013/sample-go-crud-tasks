package entity

// AuthRequest represents the login request payload
type AuthRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
