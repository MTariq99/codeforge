package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateBlogContent(c *gin.Context) {
	updateContent := models.UpdateBlogContentReq{}
	if err := c.BindJSON(&updateContent); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := s.api.UpdateBlogContentAPI(c, &updateContent); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "content of the blog updated successfully",
		Status:  http.StatusOK,
	})

}
