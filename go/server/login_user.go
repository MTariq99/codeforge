package server

import (
	"codeforge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) LoginUser(c *gin.Context) {
	userLogin := models.LoginUserReq{}
	if err := c.BindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	res, err := s.api.LoginUserAPI(c, &userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "user logged in successfully",
		Token:   res,
	})

}
