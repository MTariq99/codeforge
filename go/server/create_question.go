package server

import (
	"codeforge/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateQuestion(c *gin.Context) {
	question := models.Question{}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	reqBody := c.Request.FormValue("body")
	reqBytes := []byte(reqBody)
	if err := json.Unmarshal(reqBytes, &question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := s.api.CreateQuestionAPI(c, file, &question); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "question created successfully",
	})

}
