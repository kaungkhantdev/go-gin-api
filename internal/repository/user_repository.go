package repository

import (
	"errors"
	"fmt"
	"go-gin-api/internal/models"
)

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

func (repo *UserRepository) GetById(id int) (models.User, error) {
	for _, user := range users { 
		if user.ID == id {
			return user, nil
		}
	}
	
	return models.User{}, errors.New("user not found")
}

type CreateUserData struct {
	Name string
	Email string
}

func (repo *UserRepository) CreateUser(data CreateUserData) (models.User) {

	newId := len(users) + 1

	newUser := models.User {
		ID: newId,
		Name: data.Name,
		Email: data.Email,
	}

	users = append(users, newUser)

	return newUser;

}

type PatchUserData struct {
	Name *string
	Email *string
}

func (repo *UserRepository) UpdateUser(id int, data PatchUserData ) ( models.User, error ) {
	for i, user := range users {
		if user.ID == id {

			if data.Name != nil {
				users[i].Name = *data.Name
			}

			if data.Email != nil {
				users[i].Email = *data.Email
			}
			
			return users[i], nil
		}
	}

	return models.User{}, fmt.Errorf("user with ID %d not found", id)
}