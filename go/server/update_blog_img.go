package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateBlogImg(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := s.api.UpdateBlogImgAPI(c, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog image updated successfully"})
}
