package handlers

import (
	"go-gin-api/internal/repository"
	"go-gin-api/internal/services"
	"github.com/gin-gonic/gin"
)

// UserRoutes registers the user-related routes.
func UserRoutes(router *gin.RouterGroup) {
	// Initialize UserRepository and UserService
	userRepo := &repository.UserRepository{}
	userService := services.NewUserService(userRepo)

	// Initialize the UserHandler
	userHandler := NewUserHandler(userService)

	// Register routes with userHandler
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetById)
	router.POST("/users", userHandler.CreateUser)
	router.PATCH("/users/:id", userHandler.UpdateUser)

}

func TestRoutes(router *gin.RouterGroup) {
	testRepo := &repository.TestRepository{}
	testService := services.NewTestService(testRepo)

	testHandler := NewTestHandler(testService);

	router.GET("/tests", testHandler.GetAll)
	router.GET("/tests/:id", testHandler.GetById)
	router.POST("/tests", testHandler.Create)
	router.PATCH("/tests/:id", testHandler.Update)

}