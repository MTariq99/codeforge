package server

import (
	"codeforge/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUser(c *gin.Context) {
	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	currentTime := time.Now().UTC().Unix()
	user.Created_at = &currentTime
	if err := s.api.CreateUserAPI(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

}
