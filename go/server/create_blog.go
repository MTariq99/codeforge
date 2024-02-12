package server

import (
	"codeforge/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBlog(c *gin.Context) {
	blog := models.BlogsReq{}
	// Parse form file for the image
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	reqBody := c.Request.FormValue("body")

	// Marshal the reqBody into a byte slice
	bodyBytes := []byte(reqBody)

	// Unmarshal the byte slice into the blog variable
	if err := json.Unmarshal(bodyBytes, &blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call API method to create blog with image
	if err := s.api.CreateBlog(c, file, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully"})
}
