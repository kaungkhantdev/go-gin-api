package handlers

import (
	"go-gin-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct encapsulates dependencies, including the service layer.
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new Handler with the given UserService.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetAllUsers is an HTTP handler function that retrieves all users.
func (h *UserHandler) GetUsers(c *gin.Context ) {
    users, err := h.userService.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}