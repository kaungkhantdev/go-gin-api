package repository

import "go-gin-api/internal/models"

// UserRepository handles data operations related to users.
type UserRepository struct {}

var (
	users = []models.User{
		{ ID: 1, Name: "John Doe", Email: "john@gmail.com" },
		{ ID: 2, Name: "Rio", Email: "rio@gmail.com" },
	}
)

func (repo *UserRepository) GetUsers() ([]models.User, error) {
	return users, nil;
}
