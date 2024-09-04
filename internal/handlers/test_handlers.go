package handlers

import (
	"go-gin-api/internal/models"
	"go-gin-api/internal/repository"
	"go-gin-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	testService *services.TestService
}

func NewTestHandler(service *services.TestService) *TestHandler {
	return &TestHandler{testService: service}
}

func (h *TestHandler) Create(c *gin.Context) {
	input := models.Test{}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
	}

	data := repository.TestCreateData {
		Name: input.Name,
		Email: input.Email,
	}

	test := h.testService.CreateTestService(data);
	c.JSON(http.StatusCreated, test);
}

func (h *TestHandler) Update(c *gin.Context) {
    id := c.Param("id")

    input := models.Test{}
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

    var data repository.TestPatchData
    if input.Name != "" {
        data.Name = &input.Name
    }
    if input.Email != "" {
        data.Email = &input.Email
    }
    

    user := h.testService.UpdateTestService(intId, data)

	c.JSON(http.StatusOK, user)
}

func (h *TestHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	test := h.testService.GetById(intId)

	c.JSON(http.StatusOK, test)
}

func (h *TestHandler) GetAll(c *gin.Context) {
	tests := h.testService.GetAll()

	c.JSON(http.StatusOK, tests)
}