package repository

import (
	"go-gin-api/internal/database"
	"go-gin-api/internal/models"
)

type TestRepository struct {}

type TestCreateData struct {
	Name string
	Email string
}

func (repo *TestRepository) CreateTest(data TestCreateData) models.Test {
	test := models.Test{
		Name: data.Name,
		Email: data.Email,
	}
	database.DB.Create(&test);
	return test;
}

type TestPatchData struct {
	Name *string
	Email *string
}
func (repo *TestRepository) UpdateTest (id int, data TestPatchData) (models.Test) {
	test := models.Test{ID: uint(id)}
	database.DB.First(&test)

	if data.Name != nil {
		test.Name = *data.Name
	}
	if data.Email != nil {
		test.Email = *data.Email
	}

	database.DB.Save(&test);

	return test;
}


func (repo *TestRepository) GetAll() []models.Test {
	tests := []models.Test{}
	database.DB.Find(&tests);

	return tests;
}

func (repo *TestRepository) GetById(id int) ( models.Test) {
	test := models.Test{ID: uint(id)}

	database.DB.First(&test)
	return test
}