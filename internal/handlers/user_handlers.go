package handlers

import (
	"go-gin-api/internal/models"
	"go-gin-api/internal/repository"
	"go-gin-api/internal/services"
	"net/http"
	"strconv"

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

func (h *UserHandler) GetById(c *gin.Context) {
    id := c.Param("id")

    intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetById(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    input := models.User{}

    if err := c.ShouldBindBodyWithJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    data := repository.CreateUserData{
        Name: input.Name,
        Email: input.Email,
    }

    user := h.userService.CreateUser(data);

    c.JSON(http.StatusCreated, user)
    return
    
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")

    input := models.User{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

    var data repository.PatchUserData
    if input.Name != "" {
        data.Name = &input.Name
    }
    if input.Email != "" {
        data.Email = &input.Email
    }
    

    user, err := h.userService.UpdateUser(intId, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}