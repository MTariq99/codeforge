package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateBlogTitle(c *gin.Context) {
	updateTitle := models.UpdateBlogTitleReq{}
	if err := c.BindJSON(&updateTitle); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := s.api.UpdateBlogTitleAPI(c, &updateTitle); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "title of the blog updated successfully",
		Status:  http.StatusOK,
	})

}
