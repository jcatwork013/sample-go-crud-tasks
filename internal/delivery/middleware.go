package delivery

import (
    "net/http"
    "strings"
    "todo-api/pkg/jwt"

    "github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Lấy header Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            c.Abort()
			return
        }

        // Kiểm tra định dạng "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
            c.Abort()
            return
        }

        // Lấy token từ header
        tokenString := parts[1]

        // Xác thực token
        claims, err := jwt.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Lưu thông tin user vào context
        c.Set("user_id", claims["user_id"])
        c.Set("role", claims["role"])

        c.Next()
    }
}
