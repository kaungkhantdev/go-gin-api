package services

import (
	"go-gin-api/internal/models"
	"go-gin-api/internal/repository"
)

type TestService struct {
	repo *repository.TestRepository;
}

func NewTestService(repo *repository.TestRepository) *TestService {
	return &TestService{repo: repo}
}

func (s *TestService) CreateTestService(data repository.TestCreateData) models.Test {
	return s.repo.CreateTest(data);
}

func (s *TestService) UpdateTestService(id int, data repository.TestPatchData) models.Test {
	return s.repo.UpdateTest(id, data);
}

func (s *TestService) GetAll() []models.Test {
	return s.repo.GetAll();
}

func (s *TestService) GetById(id int) models.Test {
	return s.repo.GetById(id)
}