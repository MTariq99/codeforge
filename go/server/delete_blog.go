package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteBlog(c *gin.Context) {
	deleteBlogReq := models.DeleteBlogReq{}
	if err := c.BindJSON(&deleteBlogReq); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := s.api.DeleteBlogAPI(c, &deleteBlogReq); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "blog deleted successfully"})
}
