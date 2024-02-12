package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	if err := s.api.DeleteUser(c, &userId); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "deleted successfully",
		Status:  http.StatusOK,
	})
}
