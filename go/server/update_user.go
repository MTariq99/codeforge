package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UpdateUser(c *gin.Context) {
	updateUser := models.UpdateUserReq{}
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	if err := s.api.UpdateUserAPI(c, &updateUser); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "updated successfully",
	})
}
