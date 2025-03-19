package delivery

import (
	"net/http"
	"strconv"
	"todo-api/internal/entity"
	"todo-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	Usecase usecase.UserUseCase
}

// CreateUser API
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req entity.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.Usecase.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers API
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Usecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID API
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Usecase.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
