package services

import (
	"go-gin-api/internal/models"
	"go-gin-api/internal/repository"
)

// UserService provides user-related operations.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.repo.GetUsers();
}

