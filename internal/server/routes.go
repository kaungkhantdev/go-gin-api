package server

import (
	"go-gin-api/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()


	r.GET("/", s.HelloWorldHandler)

	apiGroup := r.Group("/api")
	handlers.UserRoutes(apiGroup)
	handlers.TestRoutes(apiGroup)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
